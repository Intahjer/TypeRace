package game

import (
	c "TypeRace/constants"
	"math/rand"
	"time"
)

var missleId = "MISSLE"
var missleStart time.Time
var missleSpeed = float64(0)
var missleMode = 0
var singleMode = false
var headstart = 7
var missleLowest = 25
var missleCurve = 15
var elminateTick float64
var playersEliminated = 0
var bestLastWpm = 0
var LastKilled string

func initMissle() {
	_, missleExists := Players.Load(missleId)
	if missleExists {
		Players.Delete(missleId)
	}
	if missleMode != 0 {
		updateMissleSpeed()
		updateMissleStart()
		if missleMode == 1 {
			singleMode = true
		} else {
			singleMode = false
			a := float64((c.TIMER + c.COUNTDOWN - headstart))
			count := float64(countPlayers())
			elminateTick = (a / count) * 1000
			playersEliminated = 0
		}
	}
}

func countPlayers() int {
	count := 0
	LoopPlayers(func(id string, _ PlayerInfo) {
		if id != missleId {
			count++
		}
	})
	return count
}

func updateMissleStart() {
	missleStart = time.Now().Add(time.Duration(headstart) * time.Second)
}

func updateMissleSpeed() {
	num := bestLastWpm
	if bestLastWpm < missleLowest {
		num = missleLowest
	}
	missleSpeed = float64(num+rand.Intn(missleCurve)-(missleCurve/2.0)) * 5.0
}

func updateMissle() {
	if missleMode != 0 {
		if time.Now().After(missleStart) {
			if singleMode {
				loadSingleMissle()
			} else {
				loadMultiMissle()
			}
		}
	}
	killPlayers()
}

func killPlayers() {
	_, missleExists := Players.Load(missleId)
	if missleMode == 1 && missleExists {
		LoopPlayers(func(id string, player PlayerInfo) {
			if id != missleId && player.KeysCorrect < GetPlayer(missleId).KeysCorrect {
				LastKilled = id
			}
		})
	}
	dead, exists := Players.Load(LastKilled)
	if exists {
		dead = dead.(PlayerInfo)
		Players.Store(LastKilled, MakePlayer(dead.(PlayerInfo).Name, dead.(PlayerInfo).KeysCorrect, dead.(PlayerInfo).KeysPressed, dead.(PlayerInfo).Playing, true))
	}
}

func loadMultiMissle() {
	tick := int(time.Until(timer).Seconds())
	if (c.TIMER-tick)/int(elminateTick) > playersEliminated {
		lowestKeys := 1000
		lowestId := ""
		LoopPlayers(func(id string, player PlayerInfo) {
			if !player.Dead && player.KeysCorrect < lowestKeys {
				lowestKeys = player.KeysCorrect
				lowestId = id
			}
		})
		LastKilled = lowestId
		Players.Store(missleId, MakePlayer("", lowestKeys, lowestKeys, false, false))
		playersEliminated++
	}
}

func loadSingleMissle() {
	tick := int(time.Until(timer).Milliseconds())
	if tick != 0 {
		a := float64((c.TIMER + c.COUNTDOWN - headstart))
		b := float64(tick) / 1000
		keys := int(missleSpeed * ((a - b) / 60.0))
		Players.Store(missleId, MakePlayer("", keys, keys, false, false))

	}
}
