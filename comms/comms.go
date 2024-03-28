package comms

import (
	"fmt"
	"net"
	"time"

	"github.com/AllenDang/giu"
	"github.com/google/uuid"
)

var SC_START = "sc01"
var SC_DISCONNECT = "sc02"
var SC_WINNER = "sc03"
var SC_PLAYER = "sc04"
var CC_UPDATE = "cc01"
var SPLIT = ":"
var EOF = "\n"
var ADDR = ":8787"
var ID = uuid.NewString()
var tick = 1 * time.Second

func Write(conn net.Conn, commands ...string) (int, error) {
	str := ""
	for _, command := range commands {
		str = str + command + SPLIT
	}
	fmt.Println(str)
	return conn.Write([]byte(str + EOF))
}

func SetAddr() []giu.Widget {
	return []giu.Widget{giu.Row(giu.Label("Address : "), giu.InputText(&ADDR))}
}

func Tick() {
	time.Sleep(tick)
}
