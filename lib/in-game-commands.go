package lib

import (
	"fmt"
	"sync"
)

func CmdPeekHand(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
	fmt.Println("Peeking hand")
}

func CmdInspectOpponent(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
	fmt.Println("Inspecting opponent")
}

func CmdShowBoard(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
	fmt.Println("Showing status of board")
}

func CmdInspectMinion(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
	fmt.Println("Inspecting minion")
}

func CmdMinionAttack(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
	fmt.Println("Attacking minion")
}

func CmdPlayCard(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
	fmt.Println("Playing card")
}

func CmdHeroPower(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
	// TODO if herpower requires target error out
	// TODO if doesnt say that hero power doesnt error out
	fmt.Println("Using hero power")
}

func CmdWeaponAttack(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
	// TODO error if you dont select target
	// TODO error if you select invalid target
	fmt.Println("Attacking with weapon")
}

func CmdShowTimer(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
}

func CmdMultistageCommand(p *Player, g *Game, mu *sync.Mutex) {
	// TODO check if p is in game
	// TODO check if p is eligible for multistage command (last card played needs to setup/buffer this command to player)
	// this design is prevention for catching arbitrary requests as multistage commands
	// Type of multistage commands can be Battlecry heal target, BattleCry deal damage, Choose One ...
}
