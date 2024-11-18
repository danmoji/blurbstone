package lib

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"
)

func CmdGetHelp(p *Player) {
	fmt.Fprintln(p.Conn, "General commands:")
	fmt.Fprintln(p.Conn, "")
	fmt.Fprintln(p.Conn, "get-help")
	fmt.Fprintln(p.Conn, "exit-server")
	fmt.Fprintln(p.Conn, "all-chat")
	fmt.Fprintln(p.Conn, "create-game")
	fmt.Fprintln(p.Conn, "destroy-game")
	fmt.Fprintln(p.Conn, "join-game GAME_ID")
	fmt.Fprintln(p.Conn, "leave-game")
	fmt.Fprintln(p.Conn, "forfeit-game")

	fmt.Fprintln(p.Conn, "In-game commands:")
	fmt.Fprintln(p.Conn, "")
	fmt.Fprintln(p.Conn, "peek-hand")
	fmt.Fprintln(p.Conn, "inspect-opponent")
	fmt.Fprintln(p.Conn, "inspect-minion")
	fmt.Fprintln(p.Conn, "show-board")
	fmt.Fprintln(p.Conn, "play-card")
	fmt.Fprintln(p.Conn, "hero-power")
	fmt.Fprintln(p.Conn, "minon-attack")
	fmt.Fprintln(p.Conn, "weapon-attack")
}

func CmdCreateGame(p *Player, games *map[int]*Game, mu *sync.Mutex) {
	if p.OcupiesGame {
		fmt.Fprintln(p.Conn, "You already have created a game.")
		fmt.Println(errors.New("creator has already created a game"))
		return
	}
	g := &Game{
		Id:        p.Id,
		CreatorId: p.Id,
	}

	mu.Lock()
	p.OcupiesGame = true
	(*games)[g.Id] = g
	mu.Unlock()

	fmt.Fprintf(p.Conn, "game created with id %d", g.Id)
	fmt.Println(p.Conn, "Join a created game with join command.")

	//TODO join game

}

func CmdDestroyGame(p *Player, games *map[int]*Game, mu *sync.Mutex) {
	if !p.OcupiesGame {
		fmt.Fprintln(p.Conn, "No active game found to destroy.")
		fmt.Fprintln(p.Conn, "You can create a game with the create-game command.")
		fmt.Println(errors.New("Player doesn't have any games created"))
		return
	}

	_, exists := (*games)[p.Id]
	if !exists {
		fmt.Fprintln(p.Conn, "No active game found to destroy.")
		fmt.Println(errors.New("no active game found for the player"))
		return
	}

	mu.Lock()
	delete(*games, p.Id)
	p.OcupiesGame = false
	mu.Unlock()
	fmt.Println(errors.New("Player does not have created any games"))
}

func CmdJoinGame(p *Player, games *map[int]*Game, mu *sync.Mutex, gameId string) {
	gId, err := strconv.Atoi(gameId)
	if err != nil {
		fmt.Fprintln(p.Conn, "Invalid game ID. Please provide a numeric value.")
		fmt.Println("Error converting string gameId to int:", err)
		return
	}

	game, exists := (*games)[gId]
	if !exists {
		fmt.Fprintln(p.Conn, "Game not found. Please check the game ID and try again.")
		fmt.Println("Game with ID", gId, "does not exist.")
		return
	}

	if game.Player1 != nil && game.Player2 != nil {
		fmt.Fprintln(p.Conn, "The game is already full.")
		fmt.Println("Game with ID", gId, "is full.")
		return
	}

	if game.Player1 == nil {
		mu.Lock()
		game.Player1 = p
		mu.Unlock()
		fmt.Fprintln(p.Conn, "You have joined the game as Player 1.")
		fmt.Printf("Player %d joined game %d as Player 1.\n", p.Id, gId)
	} else if game.Player2 == nil {
		mu.Lock()
		game.Player2 = p
		mu.Unlock()
		fmt.Fprintln(p.Conn, "You have joined the game as Player 2.")
		fmt.Printf("Player %d joined game %d as Player 2.\n", p.Id, gId)
	}

	mu.Lock()
	p.InGame = true
	p.CurrGameId = game.Id
	mu.Unlock()

	if game.Player1 != nil && game.Player2 != nil {
		fmt.Fprintln(game.Player1.Conn, "Player 2 has joined. The game can now start.")
		fmt.Fprintln(game.Player2.Conn, "Player 1 is ready. The game starts.")
		fmt.Println("Game", gId, "is ready to start.")
	}

	// TODO start the game
}

func CmdForfeitGame(p *Player, games *map[int]*Game, players *map[net.Conn]*Player, mu *sync.Mutex) {
	p1 := (*games)[p.CurrGameId].Player1
	p2 := (*games)[p.CurrGameId].Player2

	if !p.InGame {
		fmt.Fprintln(p.Conn, "You are not in a game")
		fmt.Println(errors.New("player is not currently in a game"))
		return
	}

	if p.CurrGameId == 0 {
		fmt.Fprintln(p.Conn, "Cannot find game Id.")
		fmt.Println(errors.New("players CurrGameId is 0"))
		return
	}

	if p.Id == p1.Id {
		fmt.Fprintln(p.Conn, "You have conceded the game.")
		fmt.Fprintln(p.Conn, "Returning to lobby ...")
		fmt.Fprintln((*games)[p.CurrGameId].Player2.Conn, "Player 1 conceded this game.")
		fmt.Fprintln((*games)[p.CurrGameId].Player2.Conn, "You win.")
		fmt.Fprintln((*games)[p.CurrGameId].Player2.Conn, "Returning to lobby ...")
	} else if p.Id == p2.Id {
		fmt.Fprintln(p.Conn, "You have conceded the game.")
		fmt.Fprintln(p.Conn, "Returning to lobby ...")
		fmt.Fprintln((*games)[p.CurrGameId].Player2.Conn, "Player 2 conceded this game.")
		fmt.Fprintln((*games)[p.CurrGameId].Player2.Conn, "You win.")
		fmt.Fprintln((*games)[p.CurrGameId].Player2.Conn, "Returning to lobby ...")
	} else {
		fmt.Println(errors.New("invalid player id found"))
		fmt.Fprintln(p.Conn, "error, found invalid player key")
	}

	mu.Lock()
	p1.InGame = false
	p2.InGame = false
	p1.CurrGameId = 0
	p2.CurrGameId = 0
	p1.Hand = nil
	p2.Hand = nil
	p1.OcupiesGame = false
	p2.OcupiesGame = false
	p1.IsStarting = false
	p2.IsStarting = false
	p1.IsTurn = false
	p1.IsTurn = false
	delete(*games, p.CurrGameId)
	mu.Unlock()
}

func CmdSpectateGame(player *Player, games *map[int]*Game, mu *sync.Mutex, gameId string) {
	// TODO later add game spectating functionality
}

func CmdLeaveSpectate(Player *Player, games *map[int]*Game, mu *sync.Mutex) {
	// TODO later add leave spectate mode functionality
}