package game

import (
	c "TypeRace/constants"
)

var myBest_wpm = 0
var MyName = "Guest"

func updateBest() {
	myPlayer := GetMyPlayer()
	thisBest := myPlayer.getWpm(c.TIMER)
	if thisBest > myBest_wpm {
		myBest_wpm = thisBest
	}
}

func getMyWpm(tick int) int {
	myPlayer := GetMyPlayer()
	return myPlayer.getWpm(c.TIMER - tick)
}
