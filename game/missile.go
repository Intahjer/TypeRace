package game

import (
	c "TypeRace/constants"
	"math/rand"
	"time"
)

type MissileModeEnum string

const (
	ChaseMode       MissileModeEnum = "Chase Mode"
	EliminationMode MissileModeEnum = "Elimination Mode"
	PvpMode         MissileModeEnum = "PvP Mode"
	NoMode          MissileModeEnum = "Off"
)

var MissileMode = NoMode
var missileIdCurrent = MISSILE_ID_DEFAULT
var MissileStart time.Time
var MISSILE_ID_DEFAULT = "MISSILE"
var MISSILE_DELAY_s = 7
var MISSILE_LOWEST_wpm = 25
var MISSILE_VARIATION_wpm = 15
var chaseModeSpeed_wpm = float64(0)
var eliminationModeTick float64
var eliminationModeDeaths = 0
var PlayersDead = []string{}

func initMissile() {
	resetMissile()
	initMode()
}

func initMode() {
	switch MissileMode {
	case ChaseMode:
		initChaseMode()
	case EliminationMode:
		initEliminationMode()
	case PvpMode:
		initPvpMode()
	case NoMode:
		return
	}
}

func initPvpMode() {

}

func resetMissile() {
	resetMissileId()
	_, missileExists := Players.Load(missileIdCurrent)
	if missileExists {
		Players.Delete(missileIdCurrent)
	}
	updateMissleStart()
	PlayersDead = []string{}
}

func resetMissileId() {
	missileIdCurrent = MISSILE_ID_DEFAULT
}

func initEliminationMode() {
	eliminationModeTick = (totalTime() / float64(countPlayers())) * c.MILLIS
	eliminationModeDeaths = 0
}

func totalTime() float64 {
	return float64((c.TIMER + c.COUNTDOWN - MISSILE_DELAY_s))
}

func countPlayers() int {
	count := 0
	LoopPlayers(func(id string, _ PlayerInfo) {
		if id != missileIdCurrent {
			count++
		}
	})
	return count
}

func updateMissleStart() {
	MissileStart = time.Now().Add(time.Duration(MISSILE_DELAY_s) * time.Second)
}

func initChaseMode() {
	num := playersBest_wpm
	if playersBest_wpm < MISSILE_LOWEST_wpm {
		num = MISSILE_LOWEST_wpm
	}
	chaseModeSpeed_wpm = float64(num + rand.Intn(MISSILE_VARIATION_wpm) - (MISSILE_VARIATION_wpm / 2.0))
}

func updateMissle() {
	switch MissileMode {
	case ChaseMode:
		updateChaseMode()
	case EliminationMode:
		updateEliminationMode()
	case PvpMode:
		updateChaseKill()
	case NoMode:
		return
	}
}

func updateEliminationMode() {
	tick := int(time.Until(timer).Seconds())
	if (c.TIMER-tick)/int(eliminationModeTick) > eliminationModeDeaths {
		lowestKeys := 1000
		lowestId := ""
		LoopPlayers(func(id string, player PlayerInfo) {
			if !player.IsDead && player.KeysCorrect < lowestKeys {
				lowestKeys = player.KeysCorrect
				lowestId = id
			}
		})
		PlayersDead = append(PlayersDead, lowestId)
		Players.Store(missileIdCurrent, GetNewPlayer("", lowestKeys, lowestKeys, false, false))
		eliminationModeDeaths++
	}
}

func keysPerMilli() float64 {
	tick := int(time.Until(timer).Milliseconds())
	if tick != 0 {
		secondsPassed := float64((c.TIMER + c.COUNTDOWN - MISSILE_DELAY_s) - tick)
		return ((c.MILLIS * 5.0) * (secondsPassed)) / 60.0
	}
	return 0
}

func updateChaseMode() {
	if time.Now().After(MissileStart) {
		keys := int(chaseModeSpeed_wpm * keysPerMilli())
		Players.Store(missileIdCurrent, GetNewPlayer("", keys, keys, false, false))
		updateChaseKill()
	}
}

func updateChaseKill() {
	LoopPlayers(func(id string, player PlayerInfo) {
		if id != missileIdCurrent && player.KeysCorrect < GetPlayer(missileIdCurrent).KeysCorrect {
			PlayersDead = append(PlayersDead, id)
		}
	})
}

func IsMissle(id string) bool {
	return id == missileIdCurrent
}
