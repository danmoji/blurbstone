package lib

import (
	"fmt"
	"net"
	"time"
)

type Game struct {
	Id                  int
	CreatorId           int
	Player1             *Player
	Player2             *Player
	TurnNo              int
	StartingPlayer      int
	CurrentTurnPlayerID int
	CurrentTurnTime     time.Time
	Board               *Board
	Spectators          *map[net.Conn]*Player
}

type Player struct {
	Id          int
	Conn        net.Conn
	Username    string
	DeckID      int
	OcupiesGame bool
	CurrGameId  int
	InGame      bool
	IsStarting  bool
	Hand        []Card
	Deck        Deck
	IsTurn      bool
}

type BoardMinion struct {
	BoardId  int
	Id       int
	Health   int
	Attack   int
	Statuses []string
}

type Board struct {
	P1Minions []BoardMinion
	P2Minions []BoardMinion
}

func (g *Game) StartTurn() {
	// this function is initiated after toss of coin and after milling
	// write down current time
}

func (g *Game) DrawCard() {

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
