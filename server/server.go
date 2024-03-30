package main

import (
	"bufio"
	"net"
	"strings"

	"TypeRace/comms"
	c "TypeRace/constants"
	"TypeRace/game"
	"TypeRace/stringgen"

	"github.com/AllenDang/giu"
)

var LOCAL_MODE = ""
var clients = make(map[net.Conn]string)
var difficulty = stringgen.Easy

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

func writeToClients(messages []string) {
	for conn := range clients {
		comms.Write(conn, messages...)
	}
}

func handleMessage(message string, conn net.Conn) {
	command := strings.Split(message, comms.SPLIT)
	switch command[0] {
	case comms.CC_UPDATE:
		_, exists := clients[conn]
		if !exists {
			clients[conn] = command[1]
		}

		updatePlayer(command[2], conn)
		return
	}
}

func sendStatusAll() {
	game.LoopPlayers(func(id string, player game.PlayerInfo) {
		status := []string{comms.SC_PLAYER, id, player.WritePlayer()}
		writeToClients(status)
	})
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
	game.Players.Store(playerId, playerUpdate)
}

func mainLoop() {
	if game.StartScreen {
		game.DisplayStartScreen(connect)
	} else if game.RunGame {
		game.GameRun()
	} else if game.ClientsPlaying() {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Waiting for other players to finish...")))
	} else {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Press enter to play")))
		game.EnterInput(start)
		difficulty = game.DisplayDifficultyOption(difficulty)
		game.DisplayWinner()
		game.DisplayPlayers()
		game.DisplayBest()
	}
}

func start() {
	str := stringgen.GetString(difficulty)
	if comms.ADDR != LOCAL_MODE {
		writeToClients([]string{comms.SC_START, str})
	}
	game.StartGame(str)
}

func connect() {
	if comms.ADDR != LOCAL_MODE {
		listener, _ := net.Listen("tcp", comms.ADDR)
		go connectionLoop(listener)
		go commsTick()
	}
	game.MakeMyPlayer()
}
