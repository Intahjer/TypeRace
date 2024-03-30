package game

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"TypeRace/comms"
	c "TypeRace/constants"

	"github.com/AllenDang/giu"
)

var NAME string
var keyWidgetStr []KeyWidget
var timer time.Time
var Players = sync.Map{}
var RunGame = false
var countDown = time.Now()
var StartScreen = true
var WINDOW = giu.NewMasterWindow(c.WNAME, c.WIDTH, c.HEIGHT, 0)
var GUI = giu.SingleWindow()
var lastBest = 0

func updateBest() {
	thisBest := getWPM(GetMyPlayer(), c.TIMER)
	if thisBest > lastBest {
		lastBest = thisBest
	}
}

func GameRun() {
	if !time.Now().After(countDown) {
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
		myPlayer.KeysPressed++
		keyWidgetStr = newKeyWidgetStr
		Players.Store(comms.Id, myPlayer)
	}
}

func StartGame(str string) {
	resetStats()
	resetTimer()
	keyWidgetStr = createKeyWidget(str)
	RunGame = true
}

func resetStats() {
	LoopPlayers(func(id string, player PlayerInfo) {
		Players.Store(id, MakePlayer(player.Name, 0, 0, false))
	})
}

func resetTimer() {
	timer = time.Now().Add(time.Duration(c.TIMER+c.COUNTDOWN) * time.Second)
	countDown = time.Now().Add(time.Duration(c.COUNTDOWN) * time.Second)
}

func getWinner() string {
	winner := comms.Id
	LoopPlayers(func(id string, player PlayerInfo) {
		if getWPM(player, c.TIMER) > getWPM(GetPlayer(winner), c.TIMER) {
			winner = id
		}
	})
	return GetPlayer(winner).Name
}

func getSpriteWidgets() []giu.Widget {
	layouts := []giu.Widget{}
	keys := []string{}
	LoopPlayers(func(name string, _ PlayerInfo) { keys = append(keys, name) })
	sort.Strings(keys)
	percent := getFitSize(len(keys))
	for _, key := range keys {
		player := GetPlayer(key)
		layouts = append(layouts, giu.Style().SetFontSize(17*percent).To(giu.Row(
			giu.Label(getDistance(player)),
			giu.ImageWithFile(getImage(player)).Size(75*percent, 50*percent),
			giu.Label("\n"+player.Name))))
	}
	return layouts
}

func getFitSize(playerCount int) float32 {
	if playerCount < 6 {
		return 1
	} else {
		return 6 / float32(playerCount)
	}
}

func getGameWidgets(w []KeyWidget) []giu.Widget {
	layouts := []giu.Widget{}
	if int(time.Until(timer).Seconds()) > 0 {
		tick := int(time.Until(timer).Seconds())
		layouts = append(layouts, getKeyWidgets(w)...)
		layouts = append(layouts, giu.Style().SetFontSize(30).To(&WpmWidget{tick, c.WIDTH - 40, c.HEIGHT - 40}))
		layouts = append(layouts, giu.Style().SetFontSize(30).To(&WpmWidget{getWPM(GetMyPlayer(), c.TIMER-tick), 8, c.HEIGHT - 40}))
		layouts = append(layouts, getSpriteWidgets()...)
	} else {
		RunGame = false
	}

	return layouts
}

func getBestWidget() []giu.Widget {
	layouts := []giu.Widget{}
	updateBest()
	layouts = append(layouts, giu.Style().SetFontSize(30).To(&WpmWidget{lastBest, 8, c.HEIGHT - 40}))
	return layouts
}

func getKeyWidgets(w []KeyWidget) []giu.Widget {
	layouts := []giu.Widget{}
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

func RemovePlayers() {
	for _, id := range comms.DisconnectedPlayers() {
		Players.Delete(id)
	}
}

func ClientsPlaying() bool {
	arePlaying := false
	Players.Range(func(_, data interface{}) bool {
		arePlaying = data.(PlayerInfo).Playing
		return !arePlaying
	})
	return arePlaying
}
