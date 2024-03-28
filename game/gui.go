package game

import (
	c "TypeRace/constants"

	"github.com/AllenDang/giu"
)

func DisplayWinner() {
	GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Style().SetFontSize(40).To(giu.Label(c.CENTER_X + "WINNER : " + getWinner()))))
}

func DisplayPlayers() {
	keys := c.SortStrKeys(Players)
	for _, key := range keys {
		GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(Players[key].Name)))
	}
}

func DisplayStartScreen(onEnter func()) {
	GUI.Layout(getStartWidgets()...)
	EnterInput(func() {
		onEnter()
		StartScreen = false
	})
}
func displayCountdown() {
	GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Style().SetFontSize(40).To(getCountdownWidget())))
}

func GameLoop(loop func()) {
	styleLoop := func() {
		giu.PushColorWindowBg(c.DGRAY)
		loop()
		giu.PopStyleColor()
	}
	WINDOW.Run(styleLoop)
}

func EnterInput(onEnter func()) {
	GUI.RegisterKeyboardShortcuts(giu.WindowShortcut{
		Key:      giu.KeyEnter,
		Callback: onEnter,
	})
}

func SetName() []giu.Widget {
	return []giu.Widget{giu.Row(giu.Label("Name : "), giu.InputText(&NAME))}
}

func displayGame() {
	GUI.RegisterKeyboardShortcuts(getKeyInputs()...).Layout(getGameWidgets(keyWidgetStr)...)
}
