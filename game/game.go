package game

import (
	"time"

	"TypeRace/comms"
	c "TypeRace/constants"
)

var timer time.Time
var countDown time.Time

func GameRun() {
	if int(time.Until(timer).Seconds()) < 0 {
		stopPlaying()
	}
	if !time.Now().After(countDown) {
		displayCountdown()
	} else {
		displayGame()
	}
}

func StartGame(str string) {
	resetStats()
	resetTimer()
	initMissile()
	startPlaying()
	keyWidgets = getKeyWidget(str)
}

func resetStats() {
	LoopPlayers(func(id string, player PlayerInfo) {
		Players.Store(id, GetNewPlayer(player.Name, 0, 0, false, false))
	})
}

func resetTimer() {
	timer = time.Now().Add(time.Duration(c.TIMER+c.COUNTDOWN) * time.Second)
	countDown = time.Now().Add(time.Duration(c.COUNTDOWN) * time.Second)
}

func startPlaying() {
	myPlayer := GetMyPlayer()
	myPlayer.IsPlaying = true
	Players.Store(comms.Id, myPlayer)
}

func stopPlaying() {
	myPlayer := GetMyPlayer()
	myPlayer.IsPlaying = false
	Players.Store(comms.Id, myPlayer)
}
