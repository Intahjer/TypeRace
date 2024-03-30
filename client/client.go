package main

import (
	"TypeRace/game"
	"bufio"
	"net"
	"os"
	"strings"

	"TypeRace/comms"
	c "TypeRace/constants"

	"github.com/AllenDang/giu"
)

var gameString string

func main() {
	game.NAME = "Client"
	game.GameLoop(mainLoop)
}

func readConnection(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)
		for {
			ok := scanner.Scan()
			text := scanner.Text()
			if !ok || strings.Contains(text, comms.SC_DISCONNECT) {
				os.Exit(0)
			} else {
				command := strings.Split(text, comms.SPLIT)
				switch command[0] {
				case comms.SC_PLAYER:
					if comms.Id != command[1] {
						comms.UpdatePlayer(command[1])
						game.Players.Store(command[1], game.ReadPlayer(command[2]))
					}
				case comms.SC_START:
					game.StartGame(command[1])
				}
			}
		}
	}
}

func commsTick(conn net.Conn) {
	for {
		sendStatus(conn)
		game.RemovePlayers()
		comms.Tick()
	}
}

func sendStatus(conn net.Conn) {
	myPlayer := game.GetMyPlayer()
	comms.Write(conn, comms.CC_UPDATE, comms.Id, myPlayer.WritePlayer())
}

func mainLoop() {
	if game.StartScreen {
		game.DisplayStartScreen(connect)
	} else if game.RunGame {
		game.GameRun()
	} else if game.ClientsPlaying() {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Waiting for other players to finish...")))
	} else {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Waiting for host...")))
		game.DisplayWinner()
		game.DisplayPlayers()
		game.DisplayBest()
	}
}

func connect() {
	conn, _ := net.Dial("tcp", comms.ADDR)
	game.MakeMyPlayer()
	go readConnection(conn)
	go commsTick(conn)
}
