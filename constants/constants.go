package constant

import (
	"image/color"

	"github.com/AllenDang/giu"
)

var SC_START = "sc01"
var SC_DISCONNECT = "sc02"
var CC_JOIN = "cc01"
var SPLIT = ":"
var WNAME = "Typing Game"
var ADDR = ":8000"
var NAME = "Guest"
var RED = color.RGBA{150, 25, 25, 225}
var WHITE = color.RGBA{225, 225, 225, 225}
var GRAY = color.RGBA{110, 110, 110, 225}
var DGRAY = color.RGBA{60, 60, 60, 225}
var WIDTH = 1280
var HEIGHT = 640
var WINDOW = giu.NewMasterWindow(WNAME, WIDTH, HEIGHT, 0)
var GUI = giu.SingleWindow()
