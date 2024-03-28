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
}

func MakePlayer(name string, keysCorrect int, keysPressed int) PlayerInfo {
	return PlayerInfo{c.SimpleName(name), keysCorrect, keysPressed}
}

func GetMyPlayer() PlayerInfo {
	return Players[comms.ADDR]
}

func (p *PlayerInfo) WritePlayer() string {
	return p.Name + ">" + strconv.Itoa(p.KeysCorrect) + ">" + strconv.Itoa(p.KeysPressed)
}

func ReadPlayer(str string) PlayerInfo {
	playerData := strings.Split(str, ">")
	arg2, _ := strconv.Atoi(playerData[1])
	arg3, _ := strconv.Atoi(playerData[2])
	return PlayerInfo{playerData[0], arg2, arg3}
}

func MakeMyPlayer() {
	Players[comms.ADDR] = MakePlayer(NAME, 0, 0)
}
