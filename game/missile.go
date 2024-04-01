package game

import (
	"TypeRace/comms"
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
var MissileIdCurrent = MISSILE_ID_DEFAULT
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
	ids := SortedIds()
	MissileIdCurrent = ids[rand.Intn(len(ids))]
}

func resetMissile() {
	resetMissileId()
	_, missileExists := Players.Load(MissileIdCurrent)
	if missileExists {
		Players.Delete(MissileIdCurrent)
	}
	updateMissleStart()
	PlayersDead = []string{}
}

func resetMissileId() {
	MissileIdCurrent = MISSILE_ID_DEFAULT
}

func initEliminationMode() {
	eliminationModeTick = ((totalTime_s() + 1) / float64(countPlayers()))
	eliminationModeDeaths = 0
}

func totalTime_s() float64 {
	return float64((c.TIMER + c.COUNTDOWN - MISSILE_DELAY_s))
}

func countPlayers() int {
	count := 0
	LoopPlayers(func(id string, _ PlayerInfo) {
		if id != MissileIdCurrent {
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

func delayMissilePlayer() bool {
	if IsMissle(comms.Id) && !time.Now().After(MissileStart) {
		return true
	}
	return false
}

func updateEliminationMode() {
	tick := time.Until(timer).Seconds()
	if int((totalTime_s())-tick)/int(eliminationModeTick) > eliminationModeDeaths {
		var lowestKeys = 1000
		var lowestId = ""
		LoopPlayers(func(id string, player PlayerInfo) {
			if !IsMissle(id) && !player.IsDead && player.KeysCorrect < lowestKeys {
				lowestKeys = player.KeysCorrect
				lowestId = id
			}
		})
		missleKill(lowestId)
		forceMisslePosition(lowestKeys)
		eliminationModeDeaths++
	}
}

func forceMisslePosition(keys int) {
	Players.Store(MissileIdCurrent, GetNewPlayer("", keys, keys, false, false))
}

func keysPerSecond(wpm float64) int {
	tick := float64(time.Until(timer).Milliseconds())
	if tick != 0 {
		secondsPassed := totalTime_s() - (tick / 1000)
		return int((((5.0) * (secondsPassed)) / 60.0) * wpm)
	}
	return 0
}

func updateChaseMode() {
	if time.Now().After(MissileStart) {
		forceMisslePosition(keysPerSecond(chaseModeSpeed_wpm))
		updateChaseKill()
	}
}

func updateChaseKill() {
	LoopPlayers(func(id string, player PlayerInfo) {
		if id != MissileIdCurrent && player.KeysCorrect < GetPlayer(MissileIdCurrent).KeysCorrect {
			missleKill(id)
		}
	})
}

func IsMissle(id string) bool {
	return id == MissileIdCurrent
}

func missleKill(id string) {
	PlayersDead = append(PlayersDead, id)
	if id == comms.Id {
		Players.Store(id, KillPlayer(GetMyPlayer()))
	}
}
