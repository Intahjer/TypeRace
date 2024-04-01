package main

import (
	"TypeRace/game"
	"bufio"
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"net"
	"os"
	"strings"

	"TypeRace/comms"
)

func main() {
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
					updatePlayer(command[1], game.ReadPlayer(command[2]))
				case comms.SC_START:
					game.StartGame(command[1])
				case comms.SC_SPRITES:
					if comms.Id == command[1] {
						game.Sprites = append(game.Sprites, readSprite(command[2]))
					}
				}
			}
		}
	}
}

func updatePlayer(id string, player game.PlayerInfo) {
	if comms.Id != id {
		comms.UpdatePlayerConnection(id)
		game.Players.Store(id, player)
	} else {
		if player.IsDead {
			myPlayer := game.KillPlayer(game.GetMyPlayer())
			game.Players.Store(comms.Id, myPlayer)
		}
	}
}

func readSprite(str string) image.Image {
	bt, _ := base64.StdEncoding.DecodeString(str)
	image, _ := png.Decode(bytes.NewReader(bt))
	return image
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
	} else if game.GetMyPlayer().IsPlaying {
		game.GameRun()
	} else if game.PlayersPlaying() {
		game.DisplayWaitingForOthers()
	} else {
		game.DisplayWaitingForHost()
		game.DisplayStats()
	}
}

func connect() {
	conn, _ := net.Dial("tcp", comms.ADDR)
	game.MakeMyPlayer()
	go readConnection(conn)
	go commsTick(conn)
}
