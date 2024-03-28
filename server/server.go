package main

import (
	"bufio"
	"net"
	"strings"

	"TypeRace/comms"
	c "TypeRace/constants"
	"TypeRace/game"

	"github.com/AllenDang/giu"
)

var playerSpace = 5
var LOCAL_MODE = ""
var test = "Jerin is a guy that made this. This string is to test the wpm accurancy which as of now should be sixty or so since these are easy words."
var clientStatus = make(map[net.Conn]bool)

func main() {
	game.NAME = "Host"
	game.GameLoop(mainLoop)
}

func connectionLoop(listener net.Listener) {
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	clientStatus[conn] = false
	scanner := bufio.NewScanner(conn)
	for {
		ok := scanner.Scan()
		if !ok {
			break
		}
		handleMessage(scanner.Text(), conn)
	}
}

func clientsPlaying() bool {
	for _, isTyping := range clientStatus {
		if isTyping {
			return true
		}
	}
	return false
}

func writeToClients(messages []string) {
	for conn := range clientStatus {
		_, err := comms.Write(conn, messages...)
		if err != nil {
			delete(clientStatus, conn)
		}
	}
}

func handleMessage(message string, conn net.Conn) {
	command := strings.Split(message, comms.SPLIT)
	switch command[0] {
	case comms.CC_UPDATE:
		if playerSpace <= 0 {
			comms.Write(conn, comms.SC_DISCONNECT)
			return
		}
		updatePlayer(command[1], conn)
		return
	}
}

func updatePlayer(update string, conn net.Conn) {
	playerUpdate := game.ReadPlayer(update)
	playerStored := game.Players[conn.RemoteAddr().String()]
	if playerUpdate.KeysCorrect > playerStored.KeysCorrect {
		clientStatus[conn] = true
	} else {
		clientStatus[conn] = false
	}
	game.Players[conn.RemoteAddr().String()] = playerUpdate
}

func mainLoop() {
	if game.StartScreen {
		game.DisplayStartScreen(connect)
	} else if game.RunGame {
		game.GameRun(test)
	} else if clientsPlaying() {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Waiting for other players to finish...")))
	} else {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Press enter to play")))
		game.EnterInput(start)
		game.DisplayWinner()
		game.DisplayPlayers()
	}
}

func start() {
	if comms.ADDR != LOCAL_MODE {
		writeToClients([]string{comms.SC_START, test})
	}
	game.RunGame = true
}

func connect() {
	if comms.ADDR != LOCAL_MODE {
		listener, _ := net.Listen("tcp", comms.ADDR)
		go connectionLoop(listener)
	}
	game.MakeMyPlayer()
}
