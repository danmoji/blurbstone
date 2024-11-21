package lib

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type Game struct {
	Id                  int
	CreatorId           int
	P1                  *Player
	P2                  *Player
	TurnNo              int
	StartingPlayer      int
	CurrentTurnPlayerID int
	CurrentTurnTime     time.Time
	Spectators          *map[net.Conn]*Player
}

type Player struct {
	Id         int
	Conn       net.Conn
	Username   string
	CurrGameId int
	InGame     bool
	IsStarting bool
	Hero       Hero
	IsTurn     bool
	Board      []BoardMinion
}

type BoardMinion struct {
	BoardId     int
	Name        string
	Id          int
	Hp          int
	MaxHp       int
	Attack      int
	Description string
	Statuses    []string
}

func (h Hero) DrawCard(p *Player, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	if len(h.Deck) == 0 {
		fmt.Printf("Player %s has no more cards in their deck.\n", p.Username)
		DamageHero(p, 2, mu)
		return
	}

	card := h.Deck[0]
	h.Deck = h.Deck[1:]
	h.Hand = append(h.Hand, card)

	fmt.Fprintln(p.Conn, "You drew: ", card.Name, "Mana cost: ", card.ManaCost)
	fmt.Printf("Player %s drew card: %s", p.Username, card.Name)
}

func HealHero(p *Player, healAmount int, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	h := p.Hero
	if h.Health == h.MaxHealth {
		fmt.Printf("%s is already at full health.\n", h.Name)
		return
	}

	originalHealth := h.Health
	h.Health += healAmount
	if h.Health > h.MaxHealth {
		h.Health = h.MaxHealth
	}

	fmt.Printf("%s heals for %d. Health: %d -> %d\n", h.Name, healAmount, originalHealth, h.Health)
}

func DamageHero(p *Player, damage int, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	h := p.Hero
	if damage >= h.Health+h.Armor {
		h.Health = 0
		h.Armor = 0
		fmt.Printf("%s has been defeated!", h.Name)
		return
	}

	if h.Armor > 0 {
		if damage <= h.Armor {
			h.Armor -= damage
			fmt.Printf("%s's armor absorbs the damage.", h.Name)
			return
		} else {
			damage -= h.Armor
			h.Armor = 0
		}
	}

	h.Health -= damage
	fmt.Printf("%s takes damage. Remaining Health: %d, Armor: %d\n", h.Name, h.Health, h.Armor)
}

func DamageMinion(you *Player, opponent *Player, mu *sync.Mutex, tgt string, damage int) {
	mu.Lock()
	defer mu.Unlock()

	// First character: '0' for your minions, '1' for opponent's minions
	targetType := tgt[0]
	minionIndexStr := tgt[1:]
	i, err := strconv.Atoi(minionIndexStr)
	if err != nil {
		fmt.Println("Invalid target index:", tgt)
		return
	}

	if targetType == '0' {
		if i < len(you.Board) && i > 0 {
			minion := you.Board[i]
			minion.Hp -= damage
			if minion.Hp <= 0 {
				fmt.Printf("Your minion %s dies!\n", minion.Name)
				you.Board = RemoveMinion(you.Board, i)
			}
		} else {
			fmt.Fprintf(you.Conn, "Invalid target: your minion %s does not exist.", tgt)
			fmt.Printf("Invalid target: minion %s does not exist.\n", tgt)
			return
		}
	} else if targetType == '1' {
		if i < len(opponent.Board) && i > 0 {
			minion := opponent.Board[i]
			minion.Hp -= damage
			if minion.Hp <= 0 {
				fmt.Printf("Opponent's minion %d dies!\n", minion.Id)
				opponent.Board = RemoveMinion(opponent.Board, i)
			}
		} else {
			fmt.Fprintf(you.Conn, "Invalid target: opponent's minion %s does not exist.", tgt)
			fmt.Printf("Invalid target: minion %s does not exist.\n", tgt)
			return
		}
	} else {
		fmt.Fprintf(you.Conn, "Invalid target type %d: minion %s does not exist.", targetType, tgt)
		fmt.Println("Invalid target: minion does not exist.")
		return
	}
}

func HealMinion(you *Player, opponent *Player, mu *sync.Mutex, tgt string, healAmount int) {
	mu.Lock()
	defer mu.Unlock()

	// First character: '0' for your minions, '1' for opponent's minions
	targetType := tgt[0]
	minionIndexStr := tgt[1:]
	i, err := strconv.Atoi(minionIndexStr)
	if err != nil {
		fmt.Println("Invalid target index:", tgt)
		return
	}

	if targetType == '0' {
		if i < len(you.Board) && i > 0 {
			minion := you.Board[i]
			originalHealth := minion.Hp
			minion.Hp += healAmount
			if minion.Hp > minion.MaxHp {
				minion.Hp = minion.MaxHp
			}
			fmt.Printf("Your minion %s healed from %d to %d health.\n", minion.Name, originalHealth, minion.Hp)
		} else {
			fmt.Fprintf(you.Conn, "Invalid target: your minion %s does not exist.", tgt)
			fmt.Printf("Invalid target: minion %s does not exist.\n", tgt)
			return
		}
	} else if targetType == '1' {
		if i < len(opponent.Board) && i > 0 {
			minion := opponent.Board[i]
			originalHealth := minion.Hp
			minion.Hp += healAmount
			if minion.Hp > minion.MaxHp {
				minion.Hp = minion.MaxHp
			}
			fmt.Printf("Opponent's minion %s healed from %d to %d health.\n", minion.Name, originalHealth, minion.Hp)
		} else {
			fmt.Fprintf(you.Conn, "Invalid target: opponent's minion %s does not exist.", tgt)
			fmt.Printf("Invalid target: minion %s does not exist.\n", tgt)
			return
		}
	} else {
		fmt.Fprintf(you.Conn, "Invalid target type %d: minion %s does not exist.", targetType, tgt)
		fmt.Println("Invalid target: minion does not exist.")
		return
	}
}

func RemoveMinion(minions []BoardMinion, i int) []BoardMinion {
	return append(minions[:i], minions[i+1:]...)
}

func (g *Game) StartTurn() {
	// this function is initiated after toss of coin and after milling
	// write down current time
}

func startRecurrentTimer() {
	interval := 1*time.Minute + 30*time.Second

	fmt.Println("Recurrent timer started for", interval)

	for {
		timer := time.NewTimer(interval)

		go func() {
			time.Sleep(interval - 10*time.Second)
			for i := 10; i > 0; i-- {
				fmt.Println(i)
				time.Sleep(1 * time.Second)
			}
		}()

		// Wait for the timer to expire
		<-timer.C

		fmt.Println("Time is up!")
	}
}

// func (g *Game) StartTurn() {

// 	duration := 1* time.Minute + 30* time.Second
// 	mu.Lock()
// 	g.CurrentTurnStartTime = time.Now()
// 	fmt.Printf("Turn %d: %s's turn (ID: %d)\n", g.TurnNumber, currentPlayer.Username, currentPlayer.Id)
// 	mu.Unlock()

// 	go func() {
// 		ticker := time.NewTicker(1 * time.Second)
// 		defer ticker.Stop()

// 		for {
// 			select {
// 			case <-ticker.C:
// 				mu.Lock()
// 				elapsed := time.Since(g.CurrentTurnStartTime)
// 				if elapsed >= g.TurnDuration {
// 					fmt.Printf("%s's turn has ended due to time expiration.\n", currentPlayer.Username)

// 					// switch current player and start the timer again
// 					g.StartTurn()
// 					g.TurnNumber++
// 					mu.Unlock()
// 					return
// 				}
// 				mu.Unlock()
// 			}
// 		}
// 	}()
// }
