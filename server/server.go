package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"net"
	"strconv"
	"strings"
	"time"

	c "TypeRace/constants"

	"github.com/AllenDang/giu"
)

var keyWidgetStr []KeyWidget
var timer time.Time
var listener net.Listener
var players = make(map[string]PlayerInfo)
var keys = make(map[rune]Key)
var playerSpace = 5
var characterIndex = 0
var keysPressed = 0
var keysCorrect = 0
var wpm = 0
var test = "ASDASDASDASDASDASDASDASDASDASDASDASDADS"
var timerDone = true
var hasSet = false

func main() {
	c.WINDOW.Run(loop)
}

func handleConnection() {
	defer listener.Close()
	playerSpace--
	for {
		conn, _ := listener.Accept()
		remoteAddr := conn.RemoteAddr().String()
		scanner := bufio.NewScanner(conn)
		for {
			ok := scanner.Scan()
			if !ok {
				break
			}
			handleMessage(scanner.Text(), conn)
		}
		delete(players, remoteAddr)
		playerSpace++
	}
}

func handleMessage(message string, conn net.Conn) {
	currentClient := conn.RemoteAddr().String()
	command := strings.Split(message, c.SPLIT)
	switch {
	case command[0] == c.CC_JOIN:
		if playerSpace <= 0 {
			conn.Write([]byte(c.SC_DISCONNECT))
			return
		}
		players[currentClient] = PlayerInfo{command[1], 0, 0.0}
		return
	}
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

func loop() {
	giu.PushColorWindowBg(c.DGRAY)
	giu.PopStyleColor()
	if !hasSet {
		c.GUI.Layout(giu.Row(giu.Label("Address : "), giu.InputText(&c.ADDR)),
			giu.Row(giu.Label("Name : "), giu.InputText(&c.NAME)),
			giu.Button("Ok").Size(200, 50).OnClick(func() {
				hasSet = true
				listener, _ = net.Listen("tcp", c.ADDR)
				go handleConnection()
				players[c.ADDR] = PlayerInfo{c.NAME, 0, 0}
				playerSpace--
			}))
	} else if timerDone {
		c.GUI.Layout(giu.Style().SetFontSize(30).To(giu.Button("Play game").Size(float32(400), float32(50)).OnClick(func() { play() })))
	} else {
		c.GUI.RegisterKeyboardShortcuts(getRKS()...).Layout(getKeyWidget(keyWidgetStr)...)
	}
	giu.Update()
}

func play() {
	timer = time.Now().Add(30 * time.Second)
	timerDone = false
	keyWidgetStr = createKeyWidget(test)
}

func keyPress(char rune) {
	newKeyWidgetStr := []KeyWidget{}
	keysPressed++
	for currentIndex, currentChar := range keyWidgetStr {
		if currentIndex == characterIndex {
			if currentChar.text == char {
				newKeyWidgetStr = append(newKeyWidgetStr, KeyWidget{currentChar.x, currentChar.y, char, c.WHITE})
				keysCorrect++
			} else {
				newKeyWidgetStr = append(newKeyWidgetStr, KeyWidget{currentChar.x, currentChar.y, char, c.RED})
			}
		} else {
			newKeyWidgetStr = append(newKeyWidgetStr, currentChar)
		}
	}
	keyWidgetStr = newKeyWidgetStr
	characterIndex++
	players[c.ADDR] = PlayerInfo{c.NAME, keysCorrect, keysPressed}
}

func createKeyWidget(in string) []KeyWidget {
	layouts := []KeyWidget{}
	for _, key := range in {
		keyWidget := KeyWidget{0, 0, key, c.GRAY}
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
			if widgetLocX/(c.WIDTH-40) != 0 {
				widgetLocY++
				widgetLocX = 0
			}
			keyWidget := KeyWidget{widgetLocX, widgetLocY, key.text, key.color}
			layouts = append(layouts, giu.Style().SetFontSize(30).To(&keyWidget))
			widgetLocX += keys[key.text].size
		}
		tick := int(time.Until(timer).Seconds())
		layouts = append(layouts, giu.Style().SetFontSize(30).To(&WpmWidget{tick, c.WIDTH - 40, c.HEIGHT - 40}))
		wpm = getWPM(30 - tick)
		layouts = append(layouts, giu.Style().SetFontSize(30).To(&WpmWidget{wpm, 8, c.HEIGHT - 40}))
		layouts = append(layouts, getSprites(players)...)
	} else {
		timerDone = true
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
	pos := image.Pt(8, 8+(c.HEIGHT/2)).Add(image.Pt(w.x, (w.y * 32)))
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
	canvas.AddText(pos, c.WHITE, buildStr)
}

func getWPM(timeElapsed int) int {
	if timeElapsed != 0 && keysPressed != 0 {
		return int(((float64(keysPressed) / 5.0) / (float64(timeElapsed) / 60.0)) * (float64(keysCorrect) / float64(keysPressed)))
	} else {
		return 0
	}
}
