package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 8000, "The port to listen on; default is 8000.")
var clientsSettings = make(map[string]Pair)
var colorOptions = []string{"red", "green", "blue"}
var maxPlayers = len(colorOptions)

type Pair struct {
	name  string
	color string
}

func main() {
	flag.Parse()

	fmt.Println("Starting server...")

	src := *addr + ":" + strconv.Itoa(*port)
	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s.\n", src)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	scanner := bufio.NewScanner(conn)

	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		handleMessage(scanner.Text(), conn)
	}

	fmt.Println("Client \"" + clientsSettings[remoteAddr].name + "\" at " + remoteAddr + " disconnected.")
}

func handleMessage(message string, conn net.Conn) {
	fmt.Println("> " + message)
	split := strings.SplitN(message, " ", 2)
	currentClient := conn.RemoteAddr().String()
	if len(split) > 1 {
		command := strings.ToLower(split[0])
		switch {
		case command == "join":
			if len(colorOptions) == 0 {
				resp := "Cannot join! Max player count of " + strconv.Itoa(maxPlayers) + " reached!\n"
				fmt.Print("< " + resp)
				conn.Write([]byte(resp))
				return
			}
			clientSetting, hasClient := clientsSettings[currentClient]
			if hasClient {
				name := split[1]
				colorOptions = append(colorOptions, clientSetting.color)
				color := colorOptions[0]
				colorOptions = colorOptions[1:]
				clientsSettings[currentClient] = Pair{name, color}
				resp := "Renamed to \"" + name + "\" with " + color + "\n"
				fmt.Print("< " + resp)
				conn.Write([]byte(resp))
				return
			} else {
				name := split[1]
				color := colorOptions[0]
				colorOptions = colorOptions[1:]
				clientsSettings[currentClient] = Pair{name, color}
				resp := "Joined as \"" + name + "\" with " + color + "\n"
				fmt.Print("< " + resp)
				conn.Write([]byte(resp))
				return
			}
		}
	}
	conn.Write([]byte("Unrecognized command.\n"))
}
