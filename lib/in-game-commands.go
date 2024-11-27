package lib

import (
	"fmt"
	"strconv"
	"sync"
)

func CmdShowTimer(p *Player, g *Game, mu *sync.Mutex) {
	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}

	mu.Lock()
	fmt.Fprintln(p.Conn, "Turn number: ", g.TurnNo, "Time: ", g.CurrentTurnTime)
	mu.Unlock()

	fmt.Println("Showing timer")
}

func CmdShowBoard(p *Player, g *Game, mu *sync.Mutex) {
	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}
	fmt.Println("Showing status of board")
}

func CmdPeekHand(p *Player, g *Game, mu *sync.Mutex) {
	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}
	fmt.Println("Peeking hand")
}

func CmdInspectTarget(p *Player, g *Game, mu *sync.Mutex) {
	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}
	fmt.Println("Inspecting opponent")
}

func CmdHeroPower(p *Player, g *Game, mu *sync.Mutex, tgt string) {
	h := p.Hero
	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}

	if !p.IsTurn {
		fmt.Fprintln(p.Conn, "Not your turn!")
		return
	}

	if p.Hero.ManaCrystals < 2 {
		fmt.Fprintln(p.Conn, "You don't have enough mana.")
		return
	}

	// TODO if herpower requires target error out
	if h.Class == "mage" || h.Class == "priest" {
		// TODO you need to have target
	}

	if h.Class != "mage" && h.Class != "priest" {
		// TODO you can't have target
	}

	switch h.Class {
	case "warrior":
		ArmorUp(p, g, mu)
	case "rouge":
		DaggerMastery(p, g, mu)
	case "mage":
		Fireblast(p, g, mu, tgt)
	case "priest":
		LesserHeal(p, g, mu, tgt)
	case "warlock":
		LifeTap(p, g, mu)
	case "paladin":
		Reinforce(p, g, mu)
	case "druid":
		ShapeShift(p, g, mu)
	case "hunter":
		SteadyShot(p, g, mu)
	case "shaman":
		TotemicCall(p, g, mu)
	default:
		fmt.Println("Wrong class.")
	}

	fmt.Println("Using hero power")
}

func CmdWeaponAttack(p *Player, g *Game, mu *sync.Mutex) {
	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}

	if !p.IsTurn {
		fmt.Fprintln(p.Conn, "Not your turn!")
		return
	}

	// TODO error if you dont select target

	// TODO error if you select invalid target

	fmt.Println("Attacking with weapon")
}

func CmdMinionAttack(p *Player, g *Game, mu *sync.Mutex, args []string) {
	var opponent *Player

	if g.P1.Id == p.Id {
		opponent = g.P2
	} else {
		opponent = g.P1
	}

	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}

	if !p.IsTurn {
		fmt.Fprintln(p.Conn, "Not your turn!")
		return
	}

	minion_marking := args[0]
	target_marking := args[1]

	minion_number, err := strconv.Atoi(minion_marking)
	if err != nil {
		fmt.Fprintln(p.Conn, "Invalid argument minion_number.", "Given:", args[0], "Expected number 11-17")
		fmt.Printf("Could not convert minion_number: %s to index", args[0])
		return
	}

	target_number, err := strconv.Atoi(target_marking)
	if err != nil {
		fmt.Fprintln(p.Conn, "Invalid argument minion_number.", "Given:", args[1], "Expected number 21-27")
		fmt.Printf("Could not convert targe_number: %s to index", args[1])
		return
	}

	if minion_number < 0 || minion_number >= len(p.Board) {
		fmt.Fprintf(p.Conn, "Minion does not exist on the board.")
		fmt.Printf("Minion %d does not exist on the board.", minion_number)
		return
	}

	if target_number < 0 || target_number >= len(opponent.Board) {
		fmt.Fprintf(p.Conn, "Opponents minion does not exist on the board.")
		fmt.Printf("Minion %d does not exist on the board.", target_number)
		return
	}

	// TODO validate args if args object is looking good

	// length must be 2 ?
	// need to be in format of 2 digit number
	// accepted values for minion number - 11 12 13 14 15 16 17
	// accepted values for target - 2 21 22 23 24 25 26 27
	// TODO make board of map that has indexes already predefined ?

	// TODO check targeting options ... if not attacking wrong target with wrong minion

	// TODO check statuses of the board

	// TODO attack accordingly with damage minon / hero functions

	fmt.Println("Attacking minion")
}

func CmdPlayCard(p *Player, g *Game, mu *sync.Mutex) {
	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}

	if !p.IsTurn {
		fmt.Fprintln(p.Conn, "Not your turn!")
		return
	}

	fmt.Println("Playing card")
}

func CmdMultistageCommand(p *Player, g *Game, mu *sync.Mutex) {
	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}

	if !p.IsTurn {
		fmt.Fprintln(p.Conn, "Not your turn!")
		return
	}
	// TODO check if p is in game
	// TODO check if p is eligible for multistage command (last card played needs to setup/buffer this command to player)
	// this design is prevention for catching arbitrary requests as multistage commands
	// Type of multistage commands can be Battlecry heal target, BattleCry deal damage, Choose One ...
}
