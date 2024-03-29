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
var clients = make(map[net.Conn]string)

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
	for _, data := range game.Players {
		if data.Playing {
			return true
		}
	}
	return false
}

func writeToClients(messages []string) {
	for conn := range clients {
		_, err := comms.Write(conn, messages...)
		if err != nil {
			delete(clients, conn)
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
		_, exists := clients[conn]
		if !exists {
			clients[conn] = command[1]
		}

		updatePlayer(command[2], conn)
		return
	}
}

func sendStatusAll() {
	for id, player := range game.Players {
		status := []string{comms.SC_PLAYER, id, player.WritePlayer()}
		writeToClients(status)
	}
}

func commsTick() {
	for {
		sendStatusAll()
		game.RemovePlayers()
		comms.Tick()
	}
}

func updatePlayer(update string, conn net.Conn) {
	playerId := clients[conn]
	comms.UpdatePlayer(playerId)
	playerUpdate := game.ReadPlayer(update)
	game.Players[playerId] = playerUpdate
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
		game.DisplayBest()
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
		go commsTick()
	}
	game.MakeMyPlayer()
}
