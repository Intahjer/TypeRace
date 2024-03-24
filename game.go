package main

import (
	"image"
	"image/color"

	"github.com/AllenDang/giu"
)

var keys = make(map[rune]Key)

type Key struct {
	key     giu.Key
	size    int
	shifted bool
}

type KeyWidget struct {
	locx int
	locy int
	text rune
}

var textStr = "Test string type me bitch `~1!2@3#4$5%6^7&8*9(0)-_=+[{\\|}];:'\",<.>/? qQwWeErRtTyYuUiIoOPpaAssSdDfFgGhHjJkKlLzZxXcCvVbBnNmM"
var str = ""
var colGreen = color.RGBA{25, 255, 25, 255}
var colRed = color.RGBA{255, 25, 25, 255}
var colWhite = color.RGBA{255, 255, 255, 255}
var width = 640
var height = 480

func loop() {
	giu.SingleWindow().RegisterKeyboardShortcuts(getRKS()...).Layout(giu.Row(getLayouts()...))
}

func getLayouts() []giu.Widget {
	widgetLocX := 0
	widgetLocY := 0
	layouts := []giu.Widget{}
	// for _, char := range str {
	// 	layouts = append(layouts, giu.Style().SetColor(giu.StyleColorText, colWhite).To(giu.Label(string(char)).Wrapped(true)))
	// }
	for _, key := range str {
		if widgetLocX/(width-20) != 0 {
			widgetLocY++
			widgetLocX = 0
		}
		keyWidget := KeyWidget{widgetLocX, widgetLocY, key}
		layouts = append(layouts, &keyWidget)
		widgetLocX += keys[key].size
	}
	return layouts
}

func main() {
	keys['a'] = Key{giu.KeyA, 8, false}
	keys['b'] = Key{giu.KeyB, 9, false}
	keys['c'] = Key{giu.KeyC, 7, false}
	keys['d'] = Key{giu.KeyD, 9, false}
	keys['e'] = Key{giu.KeyE, 9, false}
	keys['f'] = Key{giu.KeyF, 6, false}
	keys['g'] = Key{giu.KeyG, 8, false}
	keys['h'] = Key{giu.KeyH, 9, false}
	keys['i'] = Key{giu.KeyI, 5, false}
	keys['j'] = Key{giu.KeyJ, 5, false}
	keys['k'] = Key{giu.KeyK, 8, false}
	keys['l'] = Key{giu.KeyL, 5, false}
	keys['m'] = Key{giu.KeyM, 13, false}
	keys['n'] = Key{giu.KeyN, 9, false}
	keys['o'] = Key{giu.KeyO, 9, false}
	keys['p'] = Key{giu.KeyP, 9, false}
	keys['q'] = Key{giu.KeyQ, 9, false}
	keys['r'] = Key{giu.KeyR, 6, false}
	keys['s'] = Key{giu.KeyS, 7, false}
	keys['t'] = Key{giu.KeyT, 6, false}
	keys['u'] = Key{giu.KeyU, 10, false}
	keys['v'] = Key{giu.KeyV, 9, false}
	keys['w'] = Key{giu.KeyW, 13, false}
	keys['x'] = Key{giu.KeyX, 9, false}
	keys['y'] = Key{giu.KeyY, 9, false}
	keys['z'] = Key{giu.KeyZ, 9, false}
	keys['A'] = Key{giu.KeyA, 9, true}
	keys['B'] = Key{giu.KeyB, 9, true}
	keys['C'] = Key{giu.KeyC, 9, true}
	keys['D'] = Key{giu.KeyD, 9, true}
	keys['E'] = Key{giu.KeyE, 8, true}
	keys['F'] = Key{giu.KeyF, 8, true}
	keys['G'] = Key{giu.KeyG, 10, true}
	keys['H'] = Key{giu.KeyH, 10, true}
	keys['I'] = Key{giu.KeyI, 5, true}
	keys['J'] = Key{giu.KeyJ, 5, true}
	keys['K'] = Key{giu.KeyK, 9, true}
	keys['L'] = Key{giu.KeyL, 7, true}
	keys['M'] = Key{giu.KeyM, 13, true}
	keys['N'] = Key{giu.KeyN, 9, true}
	keys['O'] = Key{giu.KeyO, 10, true}
	keys['P'] = Key{giu.KeyP, 9, true}
	keys['Q'] = Key{giu.KeyQ, 12, true}
	keys['R'] = Key{giu.KeyR, 8, true}
	keys['S'] = Key{giu.KeyS, 8, true}
	keys['T'] = Key{giu.KeyT, 8, true}
	keys['U'] = Key{giu.KeyU, 11, true}
	keys['V'] = Key{giu.KeyV, 10, true}
	keys['W'] = Key{giu.KeyW, 14, true}
	keys['X'] = Key{giu.KeyX, 10, true}
	keys['Y'] = Key{giu.KeyY, 9, true}
	keys['Z'] = Key{giu.KeyZ, 10, true}
	keys[' '] = Key{giu.KeySpace, 4, false}
	keys['`'] = Key{giu.KeyGraveAccent, 6, false}
	keys['1'] = Key{giu.Key1, 9, false}
	keys['2'] = Key{giu.Key2, 9, false}
	keys['3'] = Key{giu.Key3, 10, false}
	keys['4'] = Key{giu.Key4, 9, false}
	keys['5'] = Key{giu.Key5, 9, false}
	keys['6'] = Key{giu.Key6, 9, false}
	keys['7'] = Key{giu.Key7, 9, false}
	keys['8'] = Key{giu.Key8, 10, false}
	keys['9'] = Key{giu.Key9, 10, false}
	keys['0'] = Key{giu.Key0, 10, false}
	keys['-'] = Key{giu.KeyMinus, 8, false}
	keys['='] = Key{giu.KeyEqual, 9, false}
	keys['['] = Key{giu.KeyLeftBracket, 8, false}
	keys[']'] = Key{giu.KeyRightBracket, 8, false}
	keys['\\'] = Key{giu.KeyBackslash, 8, false}
	keys[';'] = Key{giu.KeySemicolon, 6, false}
	keys['\''] = Key{giu.KeyApostrophe, 8, false}
	keys[','] = Key{giu.KeyComma, 6, false}
	keys['.'] = Key{giu.KeyPeriod, 6, false}
	keys['/'] = Key{giu.KeySlash, 8, false}
	keys['~'] = Key{giu.KeyGraveAccent, 8, true}
	keys['!'] = Key{giu.Key1, 6, true}
	keys['@'] = Key{giu.Key2, 15, true}
	keys['#'] = Key{giu.Key3, 9, true}
	keys['$'] = Key{giu.Key4, 10, true}
	keys['%'] = Key{giu.Key5, 12, true}
	keys['^'] = Key{giu.Key6, 8, true}
	keys['&'] = Key{giu.Key7, 10, true}
	keys['*'] = Key{giu.Key8, 9, true}
	keys['('] = Key{giu.Key9, 8, true}
	keys[')'] = Key{giu.Key0, 8, true}
	keys['_'] = Key{giu.KeyMinus, 8, true}
	keys['+'] = Key{giu.KeyEqual, 10, true}
	keys['{'] = Key{giu.KeyLeftBracket, 8, true}
	keys['}'] = Key{giu.KeyRightBracket, 8, true}
	keys['|'] = Key{giu.KeyBackslash, 6, true}
	keys[':'] = Key{giu.KeySemicolon, 6, true}
	keys['"'] = Key{giu.KeyApostrophe, 6, true}
	keys['<'] = Key{giu.KeyComma, 10, true}
	keys['>'] = Key{giu.KeyPeriod, 10, true}
	keys['?'] = Key{giu.KeySlash, 10, true}
	wnd := giu.NewMasterWindow("keyboard shortcuts", width, height, 0)
	wnd.Run(loop)
}

func keyPress(char rune) {
	str = str + string(char)
}

func getRKS() []giu.WindowShortcut {
	rks := []giu.WindowShortcut{}
	for k, v := range keys {
		if v.shifted {
			rks = append(rks, giu.WindowShortcut{
				Key:      v.key,
				Modifier: giu.ModShift,
				Callback: func() { keyPress(k) }})
		} else {
			rks = append(rks, giu.WindowShortcut{
				Key:      v.key,
				Callback: func() { keyPress(k) }})
		}
	}
	return rks
}

func (c *KeyWidget) Build() {
	pos := image.Pt(8, 8).Add(image.Pt(c.locx, (c.locy * 16)))
	canvas := giu.GetCanvas()
	buildStr := string(c.text)
	if c.text == ' ' {
		buildStr = "."
		pos = pos.Add(image.Pt(0, -4))
	}
	canvas.AddText(pos, color.RGBA{255, 255, 255, 255}, buildStr)
}
