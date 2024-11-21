package lib

import (
	"fmt"
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

func CmdMinionAttack(p *Player, g *Game, mu *sync.Mutex) {
	if !p.InGame {
		fmt.Fprintln(p.Conn, "You must be in game in order to use this command.")
		return
	}

	if !p.IsTurn {
		fmt.Fprintln(p.Conn, "Not your turn!")
		return
	}

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
