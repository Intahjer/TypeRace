package game

import (
	"fmt"
	"time"

	"TypeRace/comms"
	c "TypeRace/constants"

	"github.com/AllenDang/giu"
)

var NAME string
var keyWidgetStr []KeyWidget
var timer time.Time
var Players = make(map[string]PlayerInfo)
var timerDone = true
var RunGame = false
var countDown = time.Now()
var StartScreen = true
var WINDOW = giu.NewMasterWindow(c.WNAME, c.WIDTH, c.HEIGHT, 0)
var GUI = giu.SingleWindow()

func GameRun(str string) {
	if timerDone {
		resetStats()
		resetTimer()
		keyWidgetStr = createKeyWidget(str)
	} else if !time.Now().After(countDown) {
		displayCountdown()
	} else {
		displayGame()
	}
	giu.Update()
}
func registerKey(char rune) {
	if RunGame {
		newKeyWidgetStr := []KeyWidget{}
		myPlayer := GetMyPlayer()
		myPlayer.KeysPressed++
		for currentIndex, currentChar := range keyWidgetStr {
			if currentIndex == myPlayer.KeysPressed {
				if currentChar.text == char {
					newKeyWidgetStr = append(newKeyWidgetStr, KeyWidget{currentChar.x, currentChar.y, char, c.WHITE})
					myPlayer.KeysCorrect++
				} else {
					newKeyWidgetStr = append(newKeyWidgetStr, KeyWidget{currentChar.x, currentChar.y, char, c.RED})
				}
			} else {
				newKeyWidgetStr = append(newKeyWidgetStr, currentChar)
			}
		}
		keyWidgetStr = newKeyWidgetStr
		Players[comms.ID] = myPlayer
	}
}

func resetStats() {
	for name, player := range Players {
		Players[name] = MakePlayer(player.Name, 0, 0)
	}
}

func resetTimer() {
	timer = time.Now().Add(time.Duration(c.TIMER+c.COUNTDOWN) * time.Second)
	countDown = time.Now().Add(time.Duration(c.COUNTDOWN) * time.Second)
	timerDone = false
}

func getWinner() string {
	winner := comms.ID
	for id, player := range Players {
		if getWPM(player, c.TIMER) > getWPM(Players[winner], c.TIMER) {
			winner = id
		}
	}
	return Players[winner].Name
}

func getSpriteWidgets(playerStats map[string]PlayerInfo) []giu.Widget {
	layouts := []giu.Widget{}
	keys := c.SortStrKeys(playerStats)
	for _, key := range keys {
		player := playerStats[key]
		layouts = append(layouts, giu.Style().SetFontSize(20).To(giu.Row(
			giu.Label(getDistance(player)),
			giu.Label("\n"+player.Name),
			giu.ImageWithFile(getImage(player)))))
	}
	return layouts
}

func getGameWidgets(w []KeyWidget) []giu.Widget {
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
		layouts = append(layouts, giu.Style().SetFontSize(30).To(&WpmWidget{getWPM(GetMyPlayer(), c.TIMER-tick), 8, c.HEIGHT - 40}))
		layouts = append(layouts, getSpriteWidgets(Players)...)
	} else {
		timerDone = true
		RunGame = false
	}

	return layouts
}

func getStartWidgets() []giu.Widget {
	return append(comms.SetAddr(), SetName()...)
}

func getCountdownWidget() giu.Widget {
	left := time.Until(countDown)
	var label giu.Widget
	if left > 3*time.Second {
		label = giu.Label(c.CENTER_X + "3")
	} else if left > 2*time.Second {
		label = giu.Label(c.CENTER_X + "2")
	} else if left > 1*time.Second {
		label = giu.Label(c.CENTER_X + "1")
	} else {
		label = giu.Label(c.CENTER_X + "GO")
	}
	return label
}

func getWPM(player PlayerInfo, timeElapsed int) int {
	if timeElapsed != 0 && player.KeysPressed != 0 {
		return int(((float64(player.KeysPressed) / 5.0) / (float64(timeElapsed) / 60.0)) * (float64(player.KeysCorrect) / float64(player.KeysPressed)))
	} else {
		return 0
	}
}

func getImage(player PlayerInfo) string {
	damage := ((100 * (player.KeysCorrect + 1)) / (player.KeysPressed + 1)) / 10
	if damage < 6 {
		damage = 6
	}
	damage -= 5
	return "sprites\\Jet" + fmt.Sprint(damage) + ".png"
}

func getDistance(player PlayerInfo) string {
	space := " "
	for u := 0; u < player.KeysCorrect/2; u++ {
		space += " "
	}
	return space
}
