package comms

import (
	"net"
	"sync"
	"time"

	"github.com/AllenDang/giu"
	"github.com/google/uuid"
)

var SC_START = "sc01"
var SC_DISCONNECT = "sc02"
var SC_WINNER = "sc03"
var SC_PLAYER = "sc04"
var SC_SPRITES = "sc05"
var CC_UPDATE = "cc01"
var SPLIT = "::"
var EOF = "\n"
var ADDR = ":8787"
var Id = uuid.NewString()
var tick = 1 * time.Second
var playerTick = sync.Map{}

func Write(conn net.Conn, commands ...string) (int, error) {
	str := ""
	for _, command := range commands {
		str = str + command + SPLIT
	}
	return conn.Write([]byte(str + EOF))
}

func SetAddr() []giu.Widget {
	return []giu.Widget{giu.Row(giu.Label("Address : "), giu.InputText(&ADDR))}
}

func Tick() {
	time.Sleep(tick)
}

func UpdatePlayer(id string) {
	playerTick.Store(id, time.Now())
}

func DisconnectedPlayers() []string {
	forgetThese := []string{}
	playerTick.Range(func(id, last interface{}) bool {
		if (last.(time.Time).Add(tick * 2)).Before(time.Now()) {
			forgetThese = append(forgetThese, id.(string))
			playerTick.Delete(id)
		}
		return true
	})
	return forgetThese
}
