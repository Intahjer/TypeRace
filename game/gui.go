package game

import (
	c "TypeRace/constants"
	"TypeRace/stringgen"
	"image"
	"strconv"

	"github.com/AllenDang/giu"
)

var WINDOW = giu.NewMasterWindow(c.WNAME, c.WIDTH, c.HEIGHT, 0)
var GUI = giu.SingleWindow()
var StartScreen = true
var Sprites = []image.Image{}
var DEAD_SPRITE = 6
var MISSILE_SPRITE = 5

func displayWinner() {
	GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Style().SetFontSize(40).To(giu.Label(c.CENTER_X + "WINNER : " + getWinner()))))
	giu.Update()
}

func displayPlayers() {
	ids := SortedStats()
	size := getFitSize(countPlayers())
	for _, id := range ids {
		player := GetPlayer(id)
		GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Style().SetFontSize(15 * size).To(giu.Label(player.Name + " : " + strconv.Itoa(player.getWpm(c.TIMER))))))
	}
	giu.Update()
}

func DisplayStartScreen(onEnter func()) {
	GUI.Layout(getStartWidgets()...)
	EnterInput(func() {
		onEnter()
		StartScreen = false
	})
	giu.Update()
}

func DisplayDifficultyOption(difficulty stringgen.Difficulty) stringgen.Difficulty {
	GUI.Layout(giu.Button(string(difficulty)).OnClick(func() {
		difficulty = changeDifficulty(difficulty)
	}))
	giu.Update()
	return difficulty
}

func DisplayMissileMode() {
	GUI.Layout(giu.Button(string(MissileMode)).OnClick(func() {
		MissileMode = changeMissileMode(MissileMode)
	}))
	giu.Update()
}

func changeMissileMode(missileMode MissileModeEnum) MissileModeEnum {
	switch missileMode {
	case ChaseMode:
		return EliminationMode
	case EliminationMode:
		return PvpMode
	case PvpMode:
		return NoMode
	case NoMode:
		return ChaseMode
	default:
		return missileMode
	}

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
	giu.Update()
}

func GameLoop(loop func()) {
	styleLoop := func() {
		giu.PushColorWindowBg(c.DGRAY)
		loop()
		giu.PopStyleColor()
		giu.Update()
	}
	WINDOW.Run(styleLoop)
}

func EnterInput(onEnter func()) {
	GUI.RegisterKeyboardShortcuts(giu.WindowShortcut{
		Key:      giu.KeyEnter,
		Callback: onEnter,
	})
}

func displayGame() {
	updateMissle()
	GUI.RegisterKeyboardShortcuts(getKeyInputs()...).Layout(getGameWidgets(keyWidgets)...)
	giu.Update()
}

func displayBest() {
	GUI.Layout(getBestWidget())
	giu.Update()
}

func DisplayStats() {
	displayWinner()
	displayPlayers()
	displayBest()
}

func DisplayWaitingForHost() {
	GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Waiting for host...")))
	giu.Update()
}

func DisplayWaitingForOthers() {
	GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Waiting for other players to finish...")))
	giu.Update()
}

func DisplayHostScreen() {
	GUI.Layout(giu.Align(giu.AlignCenter).To(giu.Label(c.CENTER_X + "Press enter to play")))
	giu.Update()
}
