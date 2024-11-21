package lib

import (
	"fmt"
	"math/rand"
	"sync"
)

func ArmorUp(p *Player, g *Game, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Armor Up!")
	p.Hero.Armor += 2
	p.Hero.ManaCrystals -= 2
}

func DaggerMastery(p *Player, g *Game, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Dagger Mastery!")
	p.Hero.ManaCrystals -= 2

	dagger := Weapon{
		Id:          1,
		Name:        "Dagger",
		ManaCost:    2,
		Durability:  2,
		Attack:      1,
		Description: "A simple dagger",
		Class:       "Rogue",
	}

	p.Hero.EquipedWeapon = dagger
	p.Hero.AttackCharges = 1

	fmt.Fprintf(p.Conn, "%s equips a Dagger! (Attack: %d, Durability: %d)\n", p.Hero.Name, dagger.Attack, dagger.Durability)
	fmt.Printf("%s equips a Dagger! (Attack: %d, Durability: %d)\n", p.Hero.Name, dagger.Attack, dagger.Durability)
}

func Fireblast(p *Player, g *Game, mu *sync.Mutex, tgt string) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Fireblast")
	p.Hero.ManaCrystals -= 2

	var you *Player
	var opponent *Player

	if p.Id == g.P1.Id {
		you = g.P1
		opponent = g.P2
	} else {
		you = g.P2
		opponent = g.P1
	}

	if tgt == "0" {
		DamageHero(you, 1, mu)
	} else if tgt == "1" {
		DamageHero(opponent, 1, mu)
	} else {
		DamageMinion(you, opponent, mu, tgt, 1)
	}

	// TODO check enemy hero hp if is 0 end the game
}

func LesserHeal(p *Player, g *Game, mu *sync.Mutex, tgt string) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Lesser Heal")
	p.Hero.ManaCrystals -= 2

	var you *Player
	var opponent *Player

	if p.Id == g.P1.Id {
		you = g.P1
		opponent = g.P2
	} else {
		you = g.P2
		opponent = g.P1
	}

	if tgt == "0" {
		DamageHero(you, 1, mu)
	} else if tgt == "1" {
		DamageHero(opponent, 1, mu)
	} else {
		DamageMinion(you, opponent, mu, tgt, 1)
	}
}

func LifeTap(p *Player, g *Game, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	h := p.Hero
	fmt.Println("Life Tap")
	h.ManaCrystals -= 2
	DamageHero(p, 2, mu)
	h.DrawCard(p, mu)
}

func Reinforce(p *Player, g *Game, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Reinforce")
	p.Hero.ManaCrystals -= 2

	if len(p.Board) >= 7 {
		fmt.Println("Your board is full! Cannot summon a Silverhand Recruit.")
		return
	}

	silverhandRecruit := BoardMinion{
		BoardId:     len(p.Board) + 1,
		Name:        "Silverhand Recruit",
		Id:          9999,
		Hp:          1,
		MaxHp:       1,
		Attack:      1,
		Description: "",
		Statuses:    []string{},
	}

	p.Board = append(p.Board, silverhandRecruit)
	fmt.Printf("%s summoned a Silverhand Recruit!\n", p.Hero.Name)
}

func ShapeShift(p *Player, g *Game, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Shape Shift")
	p.Hero.ManaCrystals -= 2
	p.Hero.Armor += 1
	p.Hero.AttackCharges = 1
	p.Hero.Attack = 1
}

func SteadyShot(p *Player, g *Game, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Steady Shot")
	p.Hero.ManaCrystals -= 2

	var opponent *Player

	if p.Id == g.P1.Id {
		opponent = g.P2
	} else {
		opponent = g.P1
	}

	DamageHero(opponent, 2, mu)

	// TODO check enemy hero hp if is 0 end the game
}

func TotemicCall(p *Player, g *Game, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	if len(p.Board) >= 7 {
		fmt.Println("Your board is full! Cannot summon a totem.")
		return
	}

	totems := []BoardMinion{
		{
			BoardId:     0,
			Name:        "Healing Totem",
			Id:          9998,
			Hp:          2,
			MaxHp:       2,
			Attack:      0,
			Description: "At the end of your turn restore 1 Health to all friendly minions.",
			Statuses:    []string{"Healing Aura"},
		},
		{
			BoardId:  0,
			Name:     "Searing Totem",
			Id:       9997,
			Hp:       1,
			MaxHp:    1,
			Attack:   1,
			Statuses: []string{},
		},
		{
			BoardId:     0,
			Name:        "Stoneclaw Totem",
			Id:          9996,
			Hp:          2,
			MaxHp:       2,
			Attack:      0,
			Description: "Taunt",
			Statuses:    []string{"Taunt"},
		},
		{
			BoardId:     0,
			Name:        "Wrath of Air Totem",
			Id:          9995,
			Hp:          2,
			MaxHp:       2,
			Attack:      0,
			Description: "Spell Damage +1",
			Statuses:    []string{"Spell Damage +1"},
		},
	}

	existingTotems := make(map[string]bool)
	for _, minion := range p.Board {
		existingTotems[minion.Name] = true
	}

	availableTotems := []BoardMinion{}
	for _, totem := range totems {
		if !existingTotems[totem.Name] {
			availableTotems = append(availableTotems, totem)
		}
	}

	if len(availableTotems) == 0 {
		fmt.Println("All totems are already on the board. Totemic Call cannot be used.")
		return
	}

	p.Hero.ManaCrystals -= 2
	fmt.Println("Totemic Call")

	selectedTotem := totems[rand.Intn(len(totems))]
	selectedTotem.BoardId = len(p.Board) + 1
	p.Board = append(p.Board, selectedTotem)

	fmt.Fprintf(p.Conn, "%s summoned a %s!\n", p.Hero.Name, selectedTotem.Name)
	fmt.Printf("%s summoned a %s!\n", p.Hero.Name, selectedTotem.Name)
}
