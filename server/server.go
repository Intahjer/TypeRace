package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"TypeRace/game"

	"github.com/AllenDang/giu"
)

var playerSpace = 5
var addrSet = false
var LOCAL_MODE = ""
var test = "Jerin is a guy that made this. This string is to test the wpm accurancy which as of now should be sixty or so since these are easy words."
var clients = []net.Conn{}
var updating = make(map[string]bool)

func main() {
	game.WINDOW.Run(loop)
}

func connectionLoop(listener net.Listener) {
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	clients = append(clients, conn)
	fmt.Println("Client connected from " + remoteAddr)
	scanner := bufio.NewScanner(conn)
	for {
		ok := scanner.Scan()
		if !ok {
			break
		}
		handleMessage(scanner.Text(), conn)
		fmt.Println("Client disconnected from " + remoteAddr)
	}
}

func clientsStillPlaying() bool {
	for _, isStillPlaying := range updating {
		if isStillPlaying {
			return true
		}
	}
	return false
}

func writeToClients(message string) {
	for _, conn := range clients {
		idx, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			clients = append(clients[:idx], clients[idx+1:]...)
		}
	}
}

func handleMessage(message string, conn net.Conn) {
	currentClient := conn.RemoteAddr().String()
	command := strings.Split(message, game.SPLIT)
	switch {
	case command[0] == game.CC_JOIN:
		if playerSpace <= 0 {
			conn.Write([]byte(game.SC_DISCONNECT + game.SPLIT))
			return
		}
		updatedPlayer := game.ReadPlayer(command[1])
		player := game.Players[currentClient]
		if updatedPlayer.KeysPressed > player.KeysPressed {
			updating[currentClient] = true
		} else {
			updating[currentClient] = false
		}
		game.Players[currentClient] = updatedPlayer
		return
	}
}

func getWinner() string {
	winner := game.ADDR
	for addr, player := range game.Players {
		if game.GetWPM(player, game.TIMER) > game.GetWPM(game.Players[winner], game.TIMER) {
			winner = addr
		}
	}
	return game.Players[winner].Name
}

func loop() {
	giu.PushColorWindowBg(game.DGRAY)
	if !addrSet {
		game.GUI.Layout(giu.Row(giu.Label("Address : "), giu.InputText(&game.ADDR)),
			giu.Row(giu.Label("Name : "), giu.InputText(&game.NAME)))
		game.GUI.RegisterKeyboardShortcuts(giu.WindowShortcut{
			Key: giu.KeyEnter,
			Callback: func() {
				addrSet = true
				if game.ADDR != LOCAL_MODE {
					listener, _ := net.Listen("tcp", game.ADDR)
					go connectionLoop(listener)
				}
				game.Players[game.ADDR] = game.MakePlayer(game.NAME, 0, 0)
				playerSpace--
			}})
	} else if game.RunGame {
		game.GameRun(test)
	} else if clientsStillPlaying() {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(game.CENTER_X + "Waiting for other players to finish...")))
	} else {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(game.CENTER_X + "Press enter to play")))
		game.GUI.RegisterKeyboardShortcuts(giu.WindowShortcut{
			Key: giu.KeyEnter,
			Callback: func() {
				if game.ADDR != LOCAL_MODE {
					writeToClients(game.SC_START + game.SPLIT + test)
				}
				game.RunGame = true
			}})
		if game.ThisPlayer().KeysPressed != 0 {
			game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Style().SetFontSize(40).To(giu.Label(game.CENTER_X + "WINNER : " + getWinner()))))
		}
	}
	giu.PopStyleColor()
}
