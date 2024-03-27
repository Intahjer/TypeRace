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
	fmt.Println("Client connected from " + remoteAddr)
	scanner := bufio.NewScanner(conn)
	for {
		ok := scanner.Scan()
		fmt.Println("A")
		if !ok {
			fmt.Println("B")
			break
		}
		fmt.Println("C")
		handleMessage(scanner.Text(), conn)
	}
}

func handleMessage(message string, conn net.Conn) {
	fmt.Println(message)
	currentClient := conn.RemoteAddr().String()
	command := strings.Split(message, game.SPLIT)
	switch {
	case command[0] == game.CC_JOIN:
		if playerSpace <= 0 {
			conn.Write([]byte(game.SC_DISCONNECT + game.SPLIT))
			return
		}
		game.Players[currentClient] = game.ReadPlayer(command[1])
		return
	}
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
	} else {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(game.CENTER_X + "Press enter to play")))
		game.GUI.RegisterKeyboardShortcuts(giu.WindowShortcut{
			Key: giu.KeyEnter,
			Callback: func() {
				game.RunGame = true
			}})
	}
	giu.PopStyleColor()
}
