package main

import (
	"image"
	"image/color"

	g "github.com/AllenDang/giu"
)

type CircleButtonWidget struct {
	id  rune
	pos int
}

func CircleButton(id rune, pos int) *CircleButtonWidget {
	return &CircleButtonWidget{
		id:  id,
		pos: pos,
	}
}

func (c *CircleButtonWidget) Build() {
	pos := g.GetCursorScreenPos()
	canvas := g.GetCanvas()
	canvas.AddText(pos.Add(image.Pt(c.pos*10, 0)), color.RGBA{255, 255, 255, 255}, string(c.id))
}

func loop() {
	g.SingleWindow().Layout(
		g.Row(CircleButton('H', 1), CircleButton('e', 2)),
		CircleButton('l', 3), CircleButton('l', 4), CircleButton('o', 5),
	)
}

func main() {
	wnd := g.NewMasterWindow("Custom Widget", 400, 300, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
