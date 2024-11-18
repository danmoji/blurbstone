package lib

import "fmt"

func CmdPeekHand() {
	fmt.Println("Peeking hand")
}

func CmdInspectOpponent() {
	fmt.Println("Inspecting opponent")
}

func CmdShowBoard() {
	fmt.Println("Showing status of board")
}

func CmdInspectMinion() {
	fmt.Println("Inspecting minion")
}

func CmdMinionAttack() {
	fmt.Println("Attacking minion")
}

func CmdPlayCard() {
	fmt.Println("Playing card")
}

func CmdHeroPower() {
	// TODO if herpower requires target error out
	// TODO if doesnt say that hero power doesnt error out
	fmt.Println("Using hero power")
}

func CmdWeaponAttack() {
	// TODO error if you dont select target
	// TODO error if you select invalid target
	fmt.Println("Attacking with weapon")
}

func CmdShowTimer() {

}
