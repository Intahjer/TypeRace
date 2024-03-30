package game

import (
	c "TypeRace/constants"
	"TypeRace/stringgen"
	"strconv"

	"github.com/AllenDang/giu"
)

func DisplayWinner() {
	GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Style().SetFontSize(40).To(giu.Label(c.CENTER_X + "WINNER : " + getWinner()))))
}

func DisplayPlayers() {
	keys := c.SortStrKeys(Players)
	for _, key := range keys {
		GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(Players[key].Name + " : " + strconv.Itoa(getWPM(Players[key], c.TIMER)))))
	}
}

func DisplayStartScreen(onEnter func()) {
	GUI.Layout(getStartWidgets()...)
	EnterInput(func() {
		onEnter()
		StartScreen = false
	})
}

func DisplayDifficultyOption(difficulty stringgen.Difficulty) stringgen.Difficulty {
	GUI.Layout(giu.Button(string(difficulty)).OnClick(func() {
		difficulty = changeDifficulty(difficulty)
	}))
	return difficulty
}

func changeDifficulty(difficulty stringgen.Difficulty) stringgen.Difficulty {
	switch difficulty {
	case stringgen.Easy:
		return stringgen.Medium
	case stringgen.Medium:
		return stringgen.Hard
	case stringgen.Hard:
		return stringgen.Super
	case stringgen.Super:
		return stringgen.Easy
	default:
		return difficulty
	}
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

func DisplayBest() {
	GUI.Layout(getBestWidget()...)
}
