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

func (w *WpmWidget) Build() {
	pos := image.Pt(w.x, w.y)
	canvas := giu.GetCanvas()
	buildStr := strconv.Itoa(w.wpm)
	canvas.AddText(pos, c.WHITE, buildStr)
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
		pos = pos.Add(image.Pt(4, -6))
	}
	canvas.AddText(pos, w.color, buildStr)
}
