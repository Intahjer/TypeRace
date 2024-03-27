package game

import (
	"fmt"
	"image"
	"image/color"
	"strconv"
	"strings"
	"time"

	"github.com/AllenDang/giu"
)

var SC_START = "sc01"
var SC_DISCONNECT = "sc02"
var CC_JOIN = "cc01"
var SPLIT = ":"
var WNAME = "Typing Game"
var ADDR = ":8000"
var NAME = "Guest"
var CENTER_X = "\n\n\n\n\n\n\n"
var TIMER = 30
var COUNTDOWN = 4
var RED = color.RGBA{150, 25, 25, 225}
var WHITE = color.RGBA{225, 225, 225, 225}
var GRAY = color.RGBA{110, 110, 110, 225}
var DGRAY = color.RGBA{60, 60, 60, 225}
var WIDTH = 1280
var HEIGHT = 640
var WINDOW = giu.NewMasterWindow(WNAME, WIDTH, HEIGHT, 0)
var GUI = giu.SingleWindow()
var keyWidgetStr []KeyWidget
var timer time.Time
var characterIndex = 0
var keysPressed = 0
var keysCorrect = 0
var wpm = 0
var Players = make(map[string]PlayerInfo)
var keys = make(map[rune]Key)
var timerDone = true
var RunGame = false
var countDown = time.Now()

func MakePlayer(name string, keysCorrect int, keysPressed int) PlayerInfo {
	return PlayerInfo{name, keysCorrect, keysPressed}
}

type PlayerInfo struct {
	name        string
	keysCorrect int
	keysPressed int
}

type Key struct {
	key     giu.Key
	size    int
	shifted bool
}

type WpmWidget struct {
	wpm int
	x   int
	y   int
}

type KeyWidget struct {
	x     int
	y     int
	text  rune
	color color.RGBA
}

func keyPress(char rune) {
	newKeyWidgetStr := []KeyWidget{}
	keysPressed++
	for currentIndex, currentChar := range keyWidgetStr {
		if currentIndex == characterIndex {
			if currentChar.text == char {
				newKeyWidgetStr = append(newKeyWidgetStr, KeyWidget{currentChar.x, currentChar.y, char, WHITE})
				keysCorrect++
			} else {
				newKeyWidgetStr = append(newKeyWidgetStr, KeyWidget{currentChar.x, currentChar.y, char, RED})
			}
		} else {
			newKeyWidgetStr = append(newKeyWidgetStr, currentChar)
		}
	}
	keyWidgetStr = newKeyWidgetStr
	characterIndex++
	Players[ADDR] = PlayerInfo{NAME, keysCorrect, keysPressed}
}

func createKeyWidget(in string) []KeyWidget {
	layouts := []KeyWidget{}
	for _, key := range in {
		keyWidget := KeyWidget{0, 0, key, GRAY}
		layouts = append(layouts, keyWidget)
	}
	return layouts
}

func getKeyWidget(w []KeyWidget) []giu.Widget {
	layouts := []giu.Widget{}
	if int(time.Until(timer).Seconds()) > 0 {
		widgetLocX := 0
		widgetLocY := 0
		for _, key := range w {
			if widgetLocX/(WIDTH-40) != 0 {
				widgetLocY++
				widgetLocX = 0
			}
			keyWidget := KeyWidget{widgetLocX, widgetLocY, key.text, key.color}
			layouts = append(layouts, giu.Style().SetFontSize(30).To(&keyWidget))
			widgetLocX += keys[key.text].size
		}
		tick := int(time.Until(timer).Seconds())
		layouts = append(layouts, giu.Style().SetFontSize(30).To(&WpmWidget{tick, WIDTH - 40, HEIGHT - 40}))
		wpm = getWPM(TIMER - tick)
		layouts = append(layouts, giu.Style().SetFontSize(30).To(&WpmWidget{wpm, 8, HEIGHT - 40}))
		layouts = append(layouts, getSprites(Players)...)
	} else {
		timerDone = true
		RunGame = false
	}

	return layouts
}

func getSprites(playerStats map[string]PlayerInfo) []giu.Widget {
	layouts := []giu.Widget{}
	for _, info := range playerStats {
		space := " "
		for u := 0; u < info.keysCorrect/2; u++ {
			space += " "
		}
		jet := ((100 * (info.keysCorrect + 1)) / (info.keysPressed + 1)) / 10
		if jet < 6 {
			jet = 6
		}
		jet -= 5
		layouts = append(layouts, giu.Style().SetFontSize(20).To(giu.Row(giu.Label(space), giu.Label("\n"+info.name), giu.ImageWithFile("sprites\\Jet"+fmt.Sprint(jet)+".png"))))
	}
	return layouts
}

func getRKS() []giu.WindowShortcut {
	keys['a'] = Key{giu.KeyA, 12, false}
	keys['b'] = Key{giu.KeyB, 15, false}
	keys['c'] = Key{giu.KeyC, 12, false}
	keys['d'] = Key{giu.KeyD, 14, false}
	keys['e'] = Key{giu.KeyE, 14, false}
	keys['f'] = Key{giu.KeyF, 8, false}
	keys['g'] = Key{giu.KeyG, 13, false}
	keys['h'] = Key{giu.KeyH, 14, false}
	keys['i'] = Key{giu.KeyI, 5, false}
	keys['j'] = Key{giu.KeyJ, 5, false}
	keys['k'] = Key{giu.KeyK, 12, false}
	keys['l'] = Key{giu.KeyL, 5, false}
	keys['m'] = Key{giu.KeyM, 22, false}
	keys['n'] = Key{giu.KeyN, 14, false}
	keys['o'] = Key{giu.KeyO, 14, false}
	keys['p'] = Key{giu.KeyP, 15, false}
	keys['q'] = Key{giu.KeyQ, 14, false}
	keys['r'] = Key{giu.KeyR, 10, false}
	keys['s'] = Key{giu.KeyS, 11, false}
	keys['t'] = Key{giu.KeyT, 9, false}
	keys['u'] = Key{giu.KeyU, 15, false}
	keys['v'] = Key{giu.KeyV, 14, false}
	keys['w'] = Key{giu.KeyW, 22, false}
	keys['x'] = Key{giu.KeyX, 13, false}
	keys['y'] = Key{giu.KeyY, 13, false}
	keys['z'] = Key{giu.KeyZ, 12, false}
	keys['A'] = Key{giu.KeyA, 16, true}
	keys['B'] = Key{giu.KeyB, 15, true}
	keys['C'] = Key{giu.KeyC, 15, true}
	keys['D'] = Key{giu.KeyD, 17, true}
	keys['E'] = Key{giu.KeyE, 13, true}
	keys['F'] = Key{giu.KeyF, 13, true}
	keys['G'] = Key{giu.KeyG, 17, true}
	keys['H'] = Key{giu.KeyH, 16, true}
	keys['I'] = Key{giu.KeyI, 8, true}
	keys['J'] = Key{giu.KeyJ, 8, true}
	keys['K'] = Key{giu.KeyK, 15, true}
	keys['L'] = Key{giu.KeyL, 13, true}
	keys['M'] = Key{giu.KeyM, 22, true}
	keys['N'] = Key{giu.KeyN, 18, true}
	keys['O'] = Key{giu.KeyO, 18, true}
	keys['P'] = Key{giu.KeyP, 15, true}
	keys['Q'] = Key{giu.KeyQ, 19, true}
	keys['R'] = Key{giu.KeyR, 16, true}
	keys['S'] = Key{giu.KeyS, 14, true}
	keys['T'] = Key{giu.KeyT, 15, true}
	keys['U'] = Key{giu.KeyU, 18, true}
	keys['V'] = Key{giu.KeyV, 17, true}
	keys['W'] = Key{giu.KeyW, 25, true}
	keys['X'] = Key{giu.KeyX, 16, true}
	keys['Y'] = Key{giu.KeyY, 15, true}
	keys['Z'] = Key{giu.KeyZ, 14, true}
	keys[' '] = Key{giu.KeySpace, 6, false}
	keys['`'] = Key{giu.KeyGraveAccent, 7, false}
	keys['1'] = Key{giu.Key1, 14, false}
	keys['2'] = Key{giu.Key2, 14, false}
	keys['3'] = Key{giu.Key3, 15, false}
	keys['4'] = Key{giu.Key4, 14, false}
	keys['5'] = Key{giu.Key5, 14, false}
	keys['6'] = Key{giu.Key6, 14, false}
	keys['7'] = Key{giu.Key7, 14, false}
	keys['8'] = Key{giu.Key8, 15, false}
	keys['9'] = Key{giu.Key9, 15, false}
	keys['0'] = Key{giu.Key0, 15, false}
	keys['-'] = Key{giu.KeyMinus, 9, false}
	keys['='] = Key{giu.KeyEqual, 13, false}
	keys['['] = Key{giu.KeyLeftBracket, 11, false}
	keys[']'] = Key{giu.KeyRightBracket, 11, false}
	keys['\\'] = Key{giu.KeyBackslash, 12, false}
	keys[';'] = Key{giu.KeySemicolon, 9, false}
	keys['\''] = Key{giu.KeyApostrophe, 7, false}
	keys[','] = Key{giu.KeyComma, 8, false}
	keys['.'] = Key{giu.KeyPeriod, 8, false}
	keys['/'] = Key{giu.KeySlash, 11, false}
	keys['~'] = Key{giu.KeyGraveAccent, 12, true}
	keys['!'] = Key{giu.Key1, 8, true}
	keys['@'] = Key{giu.Key2, 25, true}
	keys['#'] = Key{giu.Key3, 14, true}
	keys['$'] = Key{giu.Key4, 15, true}
	keys['%'] = Key{giu.Key5, 21, true}
	keys['^'] = Key{giu.Key6, 13, true}
	keys['&'] = Key{giu.Key7, 17, true}
	keys['*'] = Key{giu.Key8, 12, true}
	keys['('] = Key{giu.Key9, 10, true}
	keys[')'] = Key{giu.Key0, 10, true}
	keys['_'] = Key{giu.KeyMinus, 11, true}
	keys['+'] = Key{giu.KeyEqual, 15, true}
	keys['{'] = Key{giu.KeyLeftBracket, 11, true}
	keys['}'] = Key{giu.KeyRightBracket, 11, true}
	keys['|'] = Key{giu.KeyBackslash, 9, true}
	keys[':'] = Key{giu.KeySemicolon, 9, true}
	keys['"'] = Key{giu.KeyApostrophe, 9, true}
	keys['<'] = Key{giu.KeyComma, 14, true}
	keys['>'] = Key{giu.KeyPeriod, 14, true}
	keys['?'] = Key{giu.KeySlash, 13, true}
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

func (w *KeyWidget) Build() {
	pos := image.Pt(8, 8+(HEIGHT/2)).Add(image.Pt(w.x, (w.y * 32)))
	canvas := giu.GetCanvas()
	buildStr := string(w.text)
	if w.text == ' ' {
		buildStr = "."
		pos = pos.Add(image.Pt(0, -6))
	}
	canvas.AddText(pos, w.color, buildStr)
}

func (w *WpmWidget) Build() {
	pos := image.Pt(w.x, w.y)
	canvas := giu.GetCanvas()
	buildStr := strconv.Itoa(w.wpm)
	canvas.AddText(pos, WHITE, buildStr)
}

func (p *PlayerInfo) Write() string {
	return p.name + ">" + strconv.Itoa(p.keysCorrect) + ">" + strconv.Itoa(p.keysPressed)
}

func getWPM(timeElapsed int) int {
	if timeElapsed != 0 && keysPressed != 0 {
		return int(((float64(keysPressed) / 5.0) / (float64(timeElapsed) / 60.0)) * (float64(keysCorrect) / float64(keysPressed)))
	} else {
		return 0
	}
}

func GameRun(str string) {
	if timerDone {
		timer = time.Now().Add(time.Duration(TIMER+COUNTDOWN) * time.Second)
		countDown = time.Now().Add(time.Duration(COUNTDOWN) * time.Second)
		timerDone = false
		keyWidgetStr = createKeyWidget(str)
	} else if !time.Now().After(countDown) {
		left := time.Until(countDown)
		var label giu.Widget
		if left > 3*time.Second {
			label = giu.Label(CENTER_X + "3")
		} else if left > 2*time.Second {
			label = giu.Label(CENTER_X + "2")
		} else if left > 1*time.Second {
			label = giu.Label(CENTER_X + "1")
		} else {
			label = giu.Label(CENTER_X + "GO")
		}
		GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Style().SetFontSize(40).To(label)))
		giu.Update()
	} else {
		giu.Update()
		GUI.RegisterKeyboardShortcuts(getRKS()...).Layout(getKeyWidget(keyWidgetStr)...)
		giu.Update()
	}
}

func SimpleName(str string) string {
	corrLett := []rune{}
	for _, lett := range str {
		if (lett > 64 && lett < 91) || (lett > 96 && lett < 123) {
			corrLett = append(corrLett, lett)
		}
	}
	simpleName := string(corrLett)
	if simpleName == "" {
		simpleName = "Guest"
	}
	return simpleName
}

func ReadPlayer(str string) PlayerInfo {
	playerData := strings.Split(str, ">")
	arg2, _ := strconv.Atoi(playerData[1])
	arg3, _ := strconv.Atoi(playerData[2])
	return PlayerInfo{playerData[0], arg2, arg3}
}
