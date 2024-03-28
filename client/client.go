package main

import (
	"TypeRace/game"
	"bufio"
	"net"
	"os"
	"strings"
	"time"

	"github.com/AllenDang/giu"
)

var addrSet = false
var disconnect = false
var strToType string

func main() {
	game.WINDOW.Run(loop)
}

func readConnection(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)
		for {
			ok := scanner.Scan()
			text := scanner.Text()
			if !ok || strings.Contains(text, game.SC_DISCONNECT) {
				disconnect = true
			} else {
				command := strings.Split(text, game.SPLIT)
				if strings.Contains(command[0], game.SC_START) {
					strToType = command[1]
					game.RunGame = true
				}
			}
		}
	}
}

func sendStatus(conn net.Conn) {
	for {
		time.Sleep(1 * time.Second)
		thisPlayer := game.ThisPlayer()
		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		conn.Write([]byte(game.CC_JOIN + game.SPLIT + thisPlayer.Write() + game.EOF))
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
				conn, _ := net.Dial("tcp", game.ADDR)
				go readConnection(conn)
				go sendStatus(conn)
			}})
	} else if disconnect {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(game.CENTER_X + "Disconnected!")))
		game.GUI.RegisterKeyboardShortcuts(giu.WindowShortcut{
			Key: giu.KeyEnter,
			Callback: func() {
				os.Exit(0)
			}})
	} else if game.RunGame {
		game.GameRun(strToType)
	} else {
		game.GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(game.CENTER_X + "Waiting for host...")))
	}
	giu.PopStyleColor()
}
