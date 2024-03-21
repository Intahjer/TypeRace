package main

import (
	"github.com/AllenDang/giu"
)

var letters = []letter{}

type letter struct {
	text rune
	key  giu.Key
}

var str = ""

func loop() {
	giu.SingleWindow().RegisterKeyboardShortcuts(getRKS()...).Layout(
		giu.Label(str),
	)
}

func main() {
	letters = append(letters, letter{'a', giu.KeyA})
	letters = append(letters, letter{'b', giu.KeyB})
	letters = append(letters, letter{'c', giu.KeyC})
	letters = append(letters, letter{'d', giu.KeyD})
	letters = append(letters, letter{'e', giu.KeyE})
	letters = append(letters, letter{'f', giu.KeyF})
	letters = append(letters, letter{'g', giu.KeyG})
	letters = append(letters, letter{'h', giu.KeyH})
	letters = append(letters, letter{'i', giu.KeyI})
	letters = append(letters, letter{'j', giu.KeyJ})
	letters = append(letters, letter{'k', giu.KeyK})
	letters = append(letters, letter{'l', giu.KeyL})
	letters = append(letters, letter{'m', giu.KeyM})
	letters = append(letters, letter{'n', giu.KeyN})
	letters = append(letters, letter{'o', giu.KeyO})
	letters = append(letters, letter{'p', giu.KeyP})
	letters = append(letters, letter{'q', giu.KeyQ})
	letters = append(letters, letter{'r', giu.KeyR})
	letters = append(letters, letter{'s', giu.KeyS})
	letters = append(letters, letter{'t', giu.KeyT})
	letters = append(letters, letter{'u', giu.KeyU})
	letters = append(letters, letter{'v', giu.KeyV})
	letters = append(letters, letter{'w', giu.KeyW})
	letters = append(letters, letter{'x', giu.KeyX})
	letters = append(letters, letter{'y', giu.KeyY})
	letters = append(letters, letter{'z', giu.KeyZ})
	wnd := giu.NewMasterWindow("keyboard shortcuts", 640, 480, 0)
	wnd.Run(loop)
}

func keyPress(char rune) {
	str = str + string(char)
}

func getRKS() []giu.WindowShortcut {
	rks := []giu.WindowShortcut{}
	for _, v := range letters {
		rks = append(rks, giu.WindowShortcut{
			Key:      v.key,
			Callback: func() { keyPress(v.text) }})
		rks = append(rks, giu.WindowShortcut{
			Key:      v.key,
			Modifier: giu.ModShift,
			Callback: func() { keyPress(v.text - 32) }})
	}
	return rks
}
