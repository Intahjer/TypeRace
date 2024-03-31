package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"net"
	"os"
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
var lastKilledByHost string

func main() {
	game.NAME = "Host"
	initImages()
	game.GameLoop(mainLoop)
}

func initImages() {
	locs := []string{
		"./sprites/Character0.png",
		"./sprites/Character1.png",
		"./sprites/Character2.png",
		"./sprites/Character3.png",
		"./sprites/Character4.png",
		"./sprites/Enemy0.png",
		"./sprites/Dead0.png",
	}
	for _, loc := range locs {
		game.Sprites = append(game.Sprites, getImageFromFilePath(loc))
	}
}

func getImageFromFilePath(filePath string) image.Image {
	f, _ := os.Open(filePath)
	image, _, _ := image.Decode(f)
	f.Close()
	return image
}

func writeSprites(conn net.Conn) {
	for _, sprite := range game.Sprites {
		buffer := new(bytes.Buffer)
		png.Encode(buffer, sprite)
		strSprites := base64.StdEncoding.EncodeToString(buffer.Bytes())
		comms.Write(conn, comms.SC_SPRITES, clients[conn], strSprites)
	}
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
			writeSprites(conn)
		}
		updatePlayer(command[2], conn)
		return
	}
}

func sendStatusAll() {
	game.LoopPlayers(func(id string, player game.PlayerInfo) {
		status := []string{comms.SC_PLAYER, id, player.WritePlayer()}
		writeToClients(status)
		if game.LastKilled != lastKilledByHost {
			status = []string{comms.SC_DEAD, game.LastKilled}
			writeToClients(status)
			lastKilledByHost = game.LastKilled
		}
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
		game.DisplayMissileMode()
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
