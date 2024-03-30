package game

import (
	"TypeRace/comms"
	c "TypeRace/constants"
	"TypeRace/stringgen"
	"strconv"
	"strings"
)

type PlayerInfo struct {
	Name        string
	KeysCorrect int
	KeysPressed int
	Playing     bool
}

func MakePlayer(name string, keysCorrect int, keysPressed int, playing bool) PlayerInfo {
	return PlayerInfo{stringgen.SimplifyString(name, c.MAX_CHAR), keysCorrect, keysPressed, playing}
}

func GetMyPlayer() PlayerInfo {
	return GetPlayer(comms.Id)
}

func (p *PlayerInfo) WritePlayer() string {
	return p.Name + ">" + strconv.Itoa(p.KeysCorrect) + ">" + strconv.Itoa(p.KeysPressed) + ">" + strconv.FormatBool(RunGame)
}

func ReadPlayer(str string) PlayerInfo {
	playerData := strings.Split(str, ">")
	arg1, _ := strconv.Atoi(playerData[1])
	arg2, _ := strconv.Atoi(playerData[2])
	arg3, _ := strconv.ParseBool(playerData[3])
	return PlayerInfo{playerData[0], arg1, arg2, arg3}
}

func MakeMyPlayer() {
	Players.Store(comms.Id, MakePlayer(NAME, 0, 0, RunGame))
}

func LoopPlayers(pLoop func(name string, player PlayerInfo)) {
	Players.Range(func(k, v interface{}) bool {
		pLoop(k.(string), v.(PlayerInfo))
		return true
	})
}

func GetPlayer(k string) PlayerInfo {
	val, _ := Players.Load(k)
	return val.(PlayerInfo)
}
