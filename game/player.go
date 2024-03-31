package game

import (
	"TypeRace/comms"
	c "TypeRace/constants"
	"TypeRace/stringgen"
	"fmt"
	"strconv"
	"strings"
)

var split = ">"

type PlayerInfo struct {
	Name        string
	KeysCorrect int
	KeysPressed int
	Playing     bool
	Dead        bool
}

func MakePlayer(name string, keysCorrect int, keysPressed int, playing bool, dead bool) PlayerInfo {
	return PlayerInfo{stringgen.SimplifyString(name, c.MAX_CHAR), keysCorrect, keysPressed, playing, dead}
}

func GetMyPlayer() PlayerInfo {
	return GetPlayer(comms.Id)
}

func (p *PlayerInfo) WritePlayer() string {
	return p.Name + split + strconv.Itoa(p.KeysCorrect) + split + strconv.Itoa(p.KeysPressed) + split + strconv.FormatBool(RunGame) + split + strconv.FormatBool(p.Dead)
}

func ReadPlayer(str string) PlayerInfo {
	fmt.Println(str)
	playerData := strings.Split(str, split)
	arg1, _ := strconv.Atoi(playerData[1])
	arg2, _ := strconv.Atoi(playerData[2])
	arg3, _ := strconv.ParseBool(playerData[3])
	arg4, _ := strconv.ParseBool(playerData[4])
	return PlayerInfo{playerData[0], arg1, arg2, arg3, arg4}
}

func MakeMyPlayer() {
	Players.Store(comms.Id, MakePlayer(NAME, 0, 0, RunGame, IsDead))
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
