package comms

import (
	"net"

	"github.com/AllenDang/giu"
)

var SC_START = "sc01"
var SC_DISCONNECT = "sc02"
var SC_WINNER = "sc03"
var CC_UPDATE = "cc01"
var SPLIT = ":"
var EOF = "\n"
var ADDR = ":8787"

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
