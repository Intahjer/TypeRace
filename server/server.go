package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"TypeRace/comms"
	"TypeRace/game"

	"github.com/AllenDang/giu"
)

var playerSpace = 5
var LOCAL_MODE = ""
var test = "Jerin is a guy that made this. This string is to test the wpm accurancy which as of now should be sixty or so since these are easy words."
var clientStatus = make(map[net.Conn]bool)

func main() {
	game.NAME = "Host"
	game.GameLoop(serverLoop)
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
	clientStatus[conn] = false
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
	for _, isTyping := range clientStatus {
		if isTyping {
			return true
		}
	}
	return false
}

func writeToClients(messages []string) {
	for conn, _ := range clientStatus {
		_, err := comms.Write(conn, messages...)
		if err != nil {
			delete(clientStatus, conn)
		}
	}
}

func handleMessage(message string, conn net.Conn) {
	currentClient := conn.RemoteAddr().String()
	command := strings.Split(message, comms.SPLIT)
	switch {
	case command[0] == comms.CC_JOIN:
		if playerSpace <= 0 {
			conn.Write([]byte(comms.SC_DISCONNECT + comms.SPLIT))
			return
		}
		updatedPlayer := game.ReadPlayer(command[1])
		player := game.Players[currentClient]
		if updatedPlayer.KeysPressed > player.KeysPressed {
			clientStatus[conn] = true
		} else {
			clientStatus[conn] = false
		}
		game.Players[currentClient] = updatedPlayer
		return
	}
}

func serverLoop() {
	if game.StartScreen {
		game.DisplayStartScreen(serverConnect)
	} else if game.RunGame {
		game.GameRun(test)
	} else if clientsStillPlaying() {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(game.CENTER_X + "Waiting for other players to finish...")))
	} else {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(game.CENTER_X + "Press enter to play")))
		game.EnterInput(serverStart)
		game.DisplayWinner()
	}
}

func serverStart() {
	if comms.ADDR != LOCAL_MODE {
		writeToClients([]string{comms.SC_START, test})
	}
	game.RunGame = true
}

func serverConnect() {
	if comms.ADDR != LOCAL_MODE {
		listener, _ := net.Listen("tcp", comms.ADDR)
		go connectionLoop(listener)
	}
	game.Players[comms.ADDR] = game.MakePlayer(game.NAME, 0, 0)
}
