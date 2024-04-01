package game

import (
	"TypeRace/comms"
	c "TypeRace/constants"
	"TypeRace/stringgen"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var split = ">"
var Players = sync.Map{}
var playersBest_wpm = 0

type PlayerInfo struct {
	Name        string
	KeysCorrect int
	KeysPressed int
	IsPlaying   bool
	IsDead      bool
}

func GetNewPlayer(name string, keysCorrect int, keysPressed int, playing bool, dead bool) PlayerInfo {
	return PlayerInfo{stringgen.SimplifyString(name, c.MAX_CHAR), keysCorrect, keysPressed, playing, dead}
}

func GetMyPlayer() PlayerInfo {
	return GetPlayer(comms.Id)
}

func (p *PlayerInfo) WritePlayer() string {
	return p.Name + split + strconv.Itoa(p.KeysCorrect) + split + strconv.Itoa(p.KeysPressed) + split + strconv.FormatBool(p.IsPlaying) + split + strconv.FormatBool(p.IsDead)
}

func ReadPlayer(str string) PlayerInfo {
	playerData := strings.Split(str, split)
	arg1, _ := strconv.Atoi(playerData[1])
	arg2, _ := strconv.Atoi(playerData[2])
	arg3, _ := strconv.ParseBool(playerData[3])
	arg4, _ := strconv.ParseBool(playerData[4])
	return PlayerInfo{playerData[0], arg1, arg2, arg3, arg4}
}

func MakeMyPlayer() {
	Players.Store(comms.Id, GetNewPlayer(MyName, 0, 0, false, false))
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

func KillPlayer(player PlayerInfo) PlayerInfo {
	return GetNewPlayer(player.Name, player.KeysCorrect, player.KeysPressed, player.IsPlaying, true)
}

func RemovePlayers() {
	for _, id := range comms.DisconnectedPlayers() {
		Players.Delete(id)
	}
}

func PlayersPlaying() bool {
	arePlaying := false
	Players.Range(func(_, data interface{}) bool {
		arePlaying = data.(PlayerInfo).IsPlaying
		return !arePlaying
	})
	return arePlaying
}

func (player *PlayerInfo) getDamage() int {
	damage := ((100 * (player.KeysCorrect + 1)) / (player.KeysPressed + 1)) / 10
	if damage < 6 {
		damage = 6
	}
	if player.IsDead {
		return DEAD_SPRITE
	}
	return damage - 6
}

func (player *PlayerInfo) getDistance() string {
	space := " "
	for u := 0; u < player.KeysCorrect/2; u++ {
		space += " "
	}
	return space
}

func incrKeysCorrect(player PlayerInfo) PlayerInfo {
	return GetNewPlayer(player.Name, player.KeysCorrect+1, player.KeysPressed, player.IsPlaying, player.IsDead)
}

func incrKeysPressed(player PlayerInfo) PlayerInfo {
	return GetNewPlayer(player.Name, player.KeysCorrect, player.KeysPressed+1, player.IsPlaying, player.IsDead)
}
func SortedIds() []string {
	keys := []string{}
	LoopPlayers(func(name string, _ PlayerInfo) { keys = append(keys, name) })
	sort.Strings(keys)
	return keys
}

func SortedStats() []string {
	keys := []string{}
	LoopPlayers(func(name string, _ PlayerInfo) { keys = append(keys, name) })
	sort.Slice(keys, func(i, j int) bool {
		if GetPlayer(keys[i]).KeysCorrect == GetPlayer(keys[j]).KeysCorrect {
			return strings.Compare(keys[i], keys[j]) == 1
		} else {
			return GetPlayer(keys[i]).KeysCorrect > GetPlayer(keys[j]).KeysCorrect
		}
	})
	return keys
}

func (player *PlayerInfo) getWpm(timeElapsed int) int {
	if timeElapsed != 0 && player.KeysPressed != 0 {
		return int(((float64(player.KeysPressed) / 5.0) / (float64(timeElapsed) / 60.0)) * (float64(player.KeysCorrect) / float64(player.KeysPressed)))
	} else {
		return 0
	}
}

func getFitSize(playerCount int) float32 {
	if playerCount < 6 {
		return 1
	} else {
		return 6 / float32(playerCount)
	}
}

func getWinner() string {
	winningId := comms.Id
	winningPlayer := GetPlayer(winningId)
	LoopPlayers(func(id string, player PlayerInfo) {
		if id != MISSILE_ID_DEFAULT {
			if player.getWpm(c.TIMER) > winningPlayer.getWpm(c.TIMER) {
				winningId = id
				winningPlayer = GetPlayer(winningId)
			}
		}
	})
	playersBest_wpm = winningPlayer.getWpm(c.TIMER)
	return GetPlayer(winningId).Name
}
