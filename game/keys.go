package game

import "github.com/AllenDang/giu"

var keys = make(map[rune]Key)

type Key struct {
	key     giu.Key
	size    int
	shifted bool
}

func assignKeys() {
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
	keys[' '] = Key{giu.KeySpace, 8, false}
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
	keys[';'] = Key{giu.KeySemicolon, 11, false}
	keys['\''] = Key{giu.KeyApostrophe, 7, false}
	keys[','] = Key{giu.KeyComma, 10, false}
	keys['.'] = Key{giu.KeyPeriod, 10, false}
	keys['/'] = Key{giu.KeySlash, 11, false}
	keys['~'] = Key{giu.KeyGraveAccent, 12, true}
	keys['!'] = Key{giu.Key1, 8, true}
	keys['@'] = Key{giu.Key2, 25, true}
	keys['#'] = Key{giu.Key3, 14, true}
	keys['$'] = Key{giu.Key4, 15, true}
	keys['%'] = Key{giu.Key5, 21, true}
	keys['^'] = Key{giu.Key6, 13, true}
	keys['&'] = Key{giu.Key7, 17, true}
	keys['*'] = Key{giu.Key8, 13, true}
	keys['('] = Key{giu.Key9, 10, true}
	keys[')'] = Key{giu.Key0, 10, true}
	keys['_'] = Key{giu.KeyMinus, 11, true}
	keys['+'] = Key{giu.KeyEqual, 15, true}
	keys['{'] = Key{giu.KeyLeftBracket, 11, true}
	keys['}'] = Key{giu.KeyRightBracket, 11, true}
	keys['|'] = Key{giu.KeyBackslash, 9, true}
	keys[':'] = Key{giu.KeySemicolon, 11, true}
	keys['"'] = Key{giu.KeyApostrophe, 9, true}
	keys['<'] = Key{giu.KeyComma, 14, true}
	keys['>'] = Key{giu.KeyPeriod, 14, true}
	keys['?'] = Key{giu.KeySlash, 15, true}

}

func getKeyInputs() []giu.WindowShortcut {
	assignKeys()
	rks := []giu.WindowShortcut{}
	for k, v := range keys {
		if v.shifted {
			rks = append(rks, giu.WindowShortcut{
				Key:      v.key,
				Modifier: giu.ModShift,
				Callback: func() { registerKey(k) }})
		} else {
			rks = append(rks, giu.WindowShortcut{
				Key:      v.key,
				Callback: func() { registerKey(k) }})
		}
	}
	return rks
}
