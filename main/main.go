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
	broadcast(fmt.Sprintf("%s has joined!", username))

	// Read messages from the client
	scanner := bufio.NewScanner(conn)
S:
	for scanner.Scan() {
		message := scanner.Text()
		input := strings.TrimSpace(message)
		parts := strings.Fields(input)
		command := parts[0]
		args := parts[1:]

		switch command {
		case "all-chat":
			broadcast(fmt.Sprintf("[%s] %s: %s", time.Now().Format("15:04:05"), username, args))
		case "exit-server":
			broadcast(fmt.Sprintf("Player %s has left.", username))
			break S
		case "get-help":
			lib.CmdGetHelp(&p)
		case "create-game":
			lib.CmdCreateGame(&p, &games, &mu)
		case "destroy-game":
			lib.CmdDestroyGame(&p, &games, &mu)
		case "join-game":
			if len(args) == 1 {
				lib.CmdJoinGame(&p, &games, &mu, args[1])
			} else {
				fmt.Fprintln(p.Conn, "Usage: join-game GAME_ID")
			}
		case "forfeit-game":
			lib.CmdForfeitGame(&p, &games, &clients, &mu)
		default:
			fmt.Fprintln(p.Conn, "Unknown command.")
			fmt.Println(p.Conn, "Write get-help")
		}
	}

	// Disconnect the client
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
	fmt.Printf("%s disconnected.\n", username)
}

func broadcast(message string) {
	mu.Lock()
	defer mu.Unlock()
	for conn := range clients {
		_, _ = fmt.Fprintln(conn, message)
	}
}
