package main

import (
	"image/color"

	"github.com/AllenDang/giu"
)

var keys = []key{}

type key struct {
	text    rune
	key     giu.Key
	shifted bool
}

var textStr = "Test string type me bitch `~1!2@3#4$5%6^7&8*9(0)-_=+[{\\|}];:'\",<.>/? qQwWeErRtTyYuUiIoOPpaAssSdDfFgGhHjJkKlLzZxXcCvVbBnNmM"
var str = ""
var colGreen = color.RGBA{25, 255, 25, 255}
var colRed = color.RGBA{255, 25, 25, 255}
var colWhite = color.RGBA{255, 255, 255, 255}

func loop() {
	giu.SingleWindow().RegisterKeyboardShortcuts(getRKS()...).Layout(giu.Row(getLayouts()...))
}

func getLayouts() []giu.Widget {
	layouts := []giu.Widget{}
	for _, char := range str {
		layouts = append(layouts, giu.Style().SetColor(giu.StyleColorText, colWhite).To(giu.Label(string(char)).Wrapped(true)))
	}
	return layouts
}

func main() {
	keys = append(keys, key{'a', giu.KeyA, false})
	keys = append(keys, key{'b', giu.KeyB, false})
	keys = append(keys, key{'c', giu.KeyC, false})
	keys = append(keys, key{'d', giu.KeyD, false})
	keys = append(keys, key{'e', giu.KeyE, false})
	keys = append(keys, key{'f', giu.KeyF, false})
	keys = append(keys, key{'g', giu.KeyG, false})
	keys = append(keys, key{'h', giu.KeyH, false})
	keys = append(keys, key{'i', giu.KeyI, false})
	keys = append(keys, key{'j', giu.KeyJ, false})
	keys = append(keys, key{'k', giu.KeyK, false})
	keys = append(keys, key{'l', giu.KeyL, false})
	keys = append(keys, key{'m', giu.KeyM, false})
	keys = append(keys, key{'n', giu.KeyN, false})
	keys = append(keys, key{'o', giu.KeyO, false})
	keys = append(keys, key{'p', giu.KeyP, false})
	keys = append(keys, key{'q', giu.KeyQ, false})
	keys = append(keys, key{'r', giu.KeyR, false})
	keys = append(keys, key{'s', giu.KeyS, false})
	keys = append(keys, key{'t', giu.KeyT, false})
	keys = append(keys, key{'u', giu.KeyU, false})
	keys = append(keys, key{'v', giu.KeyV, false})
	keys = append(keys, key{'w', giu.KeyW, false})
	keys = append(keys, key{'x', giu.KeyX, false})
	keys = append(keys, key{'y', giu.KeyY, false})
	keys = append(keys, key{'z', giu.KeyZ, false})
	keys = append(keys, key{' ', giu.KeySpace, false})
	keys = append(keys, key{'`', giu.KeyGraveAccent, false})
	keys = append(keys, key{'1', giu.Key1, false})
	keys = append(keys, key{'2', giu.Key2, false})
	keys = append(keys, key{'3', giu.Key3, false})
	keys = append(keys, key{'4', giu.Key4, false})
	keys = append(keys, key{'5', giu.Key5, false})
	keys = append(keys, key{'6', giu.Key6, false})
	keys = append(keys, key{'7', giu.Key7, false})
	keys = append(keys, key{'8', giu.Key8, false})
	keys = append(keys, key{'9', giu.Key9, false})
	keys = append(keys, key{'0', giu.Key0, false})
	keys = append(keys, key{'-', giu.KeyMinus, false})
	keys = append(keys, key{'=', giu.KeyEqual, false})
	keys = append(keys, key{'[', giu.KeyLeftBracket, false})
	keys = append(keys, key{']', giu.KeyRightBracket, false})
	keys = append(keys, key{'\\', giu.KeyBackslash, false})
	keys = append(keys, key{';', giu.KeySemicolon, false})
	keys = append(keys, key{'\'', giu.KeyApostrophe, false})
	keys = append(keys, key{',', giu.KeyComma, false})
	keys = append(keys, key{'.', giu.KeyPeriod, false})
	keys = append(keys, key{'/', giu.KeySlash, false})
	keys = append(keys, key{'~', giu.KeyGraveAccent, true})
	keys = append(keys, key{'!', giu.Key1, true})
	keys = append(keys, key{'@', giu.Key2, true})
	keys = append(keys, key{'#', giu.Key3, true})
	keys = append(keys, key{'$', giu.Key4, true})
	keys = append(keys, key{'%', giu.Key5, true})
	keys = append(keys, key{'^', giu.Key6, true})
	keys = append(keys, key{'&', giu.Key7, true})
	keys = append(keys, key{'*', giu.Key8, true})
	keys = append(keys, key{'(', giu.Key9, true})
	keys = append(keys, key{')', giu.Key0, true})
	keys = append(keys, key{'_', giu.KeyMinus, true})
	keys = append(keys, key{'+', giu.KeyEqual, true})
	keys = append(keys, key{'{', giu.KeyLeftBracket, true})
	keys = append(keys, key{'}', giu.KeyRightBracket, true})
	keys = append(keys, key{'|', giu.KeyBackslash, true})
	keys = append(keys, key{':', giu.KeySemicolon, true})
	keys = append(keys, key{'"', giu.KeyApostrophe, true})
	keys = append(keys, key{'<', giu.KeyComma, true})
	keys = append(keys, key{'>', giu.KeyPeriod, true})
	keys = append(keys, key{'?', giu.KeySlash, true})
	wnd := giu.NewMasterWindow("keyboard shortcuts", 640, 480, 0)
	wnd.Run(loop)
}

func keyPress(char rune) {
	str = str + string(char)
}

func getRKS() []giu.WindowShortcut {
	rks := []giu.WindowShortcut{}
	for _, v := range keys {
		if v.shifted {
			rks = append(rks, giu.WindowShortcut{
				Key:      v.key,
				Modifier: giu.ModShift,
				Callback: func() { keyPress(v.text) }})
		} else {
			rks = append(rks, giu.WindowShortcut{
				Key:      v.key,
				Callback: func() { keyPress(v.text) }})
		}
		if v.text <= 122 && v.text >= 97 {
			rks = append(rks, giu.WindowShortcut{
				Key:      v.key,
				Modifier: giu.ModShift,
				Callback: func() { keyPress(v.text - 32) }})
		}
	}
	return rks
}
