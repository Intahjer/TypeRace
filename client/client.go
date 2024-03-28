package main

import (
	"TypeRace/game"
	"bufio"
	"net"
	"strings"
	"time"

	"TypeRace/comms"
	c "TypeRace/constants"

	"github.com/AllenDang/giu"
)

var disconnect = false
var gameString string
var updateInterval = 1 * time.Second

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
				disconnect = true
			} else {
				command := strings.Split(text, comms.SPLIT)
				if strings.Contains(command[0], comms.SC_START) {
					gameString = command[1]
					game.RunGame = true
				}
			}
		}
	}
}

func sendStatus(conn net.Conn) {
	for {
		myPlayer := game.GetMyPlayer()
		comms.Write(conn, comms.CC_UPDATE, myPlayer.WritePlayer())
		time.Sleep(updateInterval)
	}

}

func mainLoop() {
	if game.StartScreen {
		game.DisplayStartScreen(connect)
	} else if disconnect {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Disconnected!")))
	} else if game.RunGame {
		game.GameRun(gameString)
	} else {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Waiting for host...")))
	}
}

func connect() {
	conn, _ := net.Dial("tcp", comms.ADDR)
	game.MakeMyPlayer()
	go readConnection(conn)
	go sendStatus(conn)
}
