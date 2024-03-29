package constants

import (
	"image/color"
	"sort"
)

var RED = color.RGBA{150, 25, 25, 225}
var WHITE = color.RGBA{225, 225, 225, 225}
var GRAY = color.RGBA{110, 110, 110, 225}
var DGRAY = color.RGBA{60, 60, 60, 225}
var EOF = "\n"
var CENTER_X = "\n\n\n\n\n\n\n"
var DEFAULT_NAME = "Guest"
var TIMER = 30
var COUNTDOWN = 4
var WNAME = "Typing Game"
var WIDTH = 1280
var HEIGHT = 640
var MAX_CHAR = 10

func SortStrKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func SimpleName(str string) string {
	corrLett := []rune{}
	for idx, lett := range str {
		if idx <= MAX_CHAR-1 {
			if (lett > 64 && lett < 91) || (lett > 96 && lett < 123) {
				corrLett = append(corrLett, lett)
			}
		}
	}
	simpleName := string(corrLett)
	if simpleName == "" {
		return DEFAULT_NAME
	}
	return simpleName
}
