package game

import (
	"TypeRace/comms"
	c "TypeRace/constants"
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
	return PlayerInfo{c.SimpleName(name), keysCorrect, keysPressed, playing}
}

func GetMyPlayer() PlayerInfo {
	return Players[comms.ID]
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
	Players[comms.ID] = MakePlayer(NAME, 0, 0, RunGame)
}
