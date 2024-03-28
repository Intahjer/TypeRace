package main

import (
	"TypeRace/game"
	"bufio"
	"net"
	"strings"
	"time"

	"TypeRace/comms"

	"github.com/AllenDang/giu"
)

var disconnect = false
var strToType string
var updateInterval = 1 * time.Second

func main() {
	game.NAME = "Client"
	game.GameLoop(clientLoop)
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
					strToType = command[1]
					game.RunGame = true
				}
			}
		}
	}
}

func sendStatus(conn net.Conn) {
	for {
		thisPlayer := game.ThisPlayer()
		comms.Write(conn, comms.CC_JOIN, thisPlayer.Write())
		time.Sleep(updateInterval)
	}

}

func clientLoop() {
	if game.StartScreen {
		game.DisplayStartScreen(clientConnect)
	} else if disconnect {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(game.CENTER_X + "Disconnected!")))
	} else if game.RunGame {
		game.GameRun(strToType)
	} else {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(game.CENTER_X + "Waiting for host...")))
	}
}

func clientConnect() {
	conn, _ := net.Dial("tcp", comms.ADDR)
	go readConnection(conn)
	go sendStatus(conn)
}
