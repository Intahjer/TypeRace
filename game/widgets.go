package game

import (
	c "TypeRace/constants"
	"image"
	"image/color"
	"strconv"

	"github.com/AllenDang/giu"
)

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

func createKeyWidget(in string) []KeyWidget {
	layouts := []KeyWidget{}
	for _, key := range in {
		keyWidget := KeyWidget{0, 0, key, c.GRAY}
		layouts = append(layouts, keyWidget)
	}
	return layouts
}
