package main

import (
	"blurbstone/lib"
	"bufio"
	"fmt"
	"math/rand/v2"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	clients = make(map[net.Conn]*lib.Player)
	games   = make(map[int]*lib.Game)
	mu      sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Get the client's name
	conn.Write([]byte("Enter your username: "))
	nameInput, _ := bufio.NewReader(conn).ReadString('\n')
	username := strings.TrimSpace(nameInput)

	p := lib.Player{Conn: conn, Username: username}
	p.Id = rand.Int()

	mu.Lock()
	clients[conn] = &p
	mu.Unlock()

	fmt.Printf("%s connected.\n", username)
	broadcastLobby(fmt.Sprintf("%s has joined!", username))

	// Read messages from the client
	scanner := bufio.NewScanner(conn)
S:
	for scanner.Scan() {
		var input, command string
		var parts, args []string

		message := scanner.Text()
		input = strings.TrimSpace(message)
		parts = strings.Fields(input)

		if message != "" {
			command = parts[0]
			// TODO test if this is ok if parts is empty
			args = parts[1:]
		}

		switch command {
		case "all-chat":
			if len(args) > 0 {
				broadcastLobby(fmt.Sprintf("[%s] %s: %s", time.Now().Format("15:04:05"), username, args))
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: all-chat TEXT")
			}
		case "exit-server":
			if len(args) == 0 {
				broadcastLobby(fmt.Sprintf("Player %s has left.", username))
				break S
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: exit-server")
			}
		case "get-help":
			if len(args) == 0 {
				lib.CmdGetHelp(&p)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: get-help")
			}
		case "create-game":
			if len(args) == 0 {
				lib.CmdCreateGame(&p, &games, &mu)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: create-game")
			}
		case "destroy-game":
			if len(args) == 0 {
				lib.CmdDestroyGame(&p, &games, &mu)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: destroy-game")
			}
		case "join-game":
			if len(args) == 1 {
				lib.CmdJoinGame(&p, &games, &mu, args[0])
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: join-game GAME_ID")
			}
		case "forfeit-game":
			if len(args) == 0 {
				lib.CmdForfeitGame(&p, &games, &clients, &mu)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: forfeit-game")
			}
		case "peek-hand":
			if len(args) == 0 {
				lib.CmdPeekHand(&p, games[p.CurrGameId], &mu)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: peek-hand")
			}
		case "show-time":
			if len(args) == 0 {
				lib.CmdShowTimer(&p, games[p.CurrGameId], &mu)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: show-time")
			}
		case "show-board":
			if len(args) == 0 {
				lib.CmdShowBoard(&p, games[p.CurrGameId], &mu)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: show-board")
			}
		case "inspect-target":
			if len(args) == 1 {
				lib.CmdInspectTarget(&p, games[p.CurrGameId], &mu)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: inspect-target TARGET_NUMBER")
			}
		case "hero-power":
			if len(args) == 0 || len(args) == 1 {
				var tgt string
				if len(args) == 1 {
					tgt = args[0]
				} else {
					tgt = ""
				}
				lib.CmdHeroPower(&p, games[p.CurrGameId], &mu, tgt)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: hero-power or 'hero-power TARGET_NUMBER'")
			}
		case "weapon-attack":
			if len(args) == 1 {
				lib.CmdWeaponAttack(&p, games[p.CurrGameId], &mu)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: weapon-attack or 'weapon-attack TARGET_NUMBER'")
			}
		case "minion-attack":
			if len(args) == 2 {
				lib.CmdMinionAttack(&p, games[p.CurrGameId], &mu, args)
			} else {
				fmt.Fprintln(p.Conn, "Invalid number of arguments.")
				fmt.Fprintln(p.Conn, "Command usage: minion-attack or 'minion-attack MINION_NUMBER TARGET_NUMBER'")
			}
		case "play-card":
			if len(args) == 0 {
				lib.CmdPlayCard(&p, games[p.CurrGameId], &mu)
			} else {
				// TODO probably you want atleast 1 argument, sometimes 2 depending on a card played
			}
		default:
			fmt.Fprintln(p.Conn, "Unknown command.")
			fmt.Fprintln(p.Conn, "Write get-help")
		}
	}

	// Disconnect the client
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	fmt.Printf("%s disconnected.\n", username)
}

func broadcastLobby(message string) {
	mu.Lock()
	defer mu.Unlock()
	for conn, client := range clients {
		if !client.InGame {
			fmt.Fprintln(conn, message)
		}
	}
}
