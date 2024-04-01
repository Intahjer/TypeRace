package game

import (
	"TypeRace/comms"
	c "TypeRace/constants"
	"strconv"
	"time"

	"github.com/AllenDang/giu"
)

func getStartWidgets() []giu.Widget {
	return append(comms.GetAddrWidget(), getNameWidget())
}

func getNameWidget() giu.Widget {
	return giu.Row(giu.Label("Name : "), giu.InputText(&MyName))
}

func getCountdownWidget() giu.Widget {
	left := time.Until(countDown)
	var label giu.Widget
	if left > 3*time.Second {
		label = giu.Label(c.CENTER_X + "3")
	} else if left > 2*time.Second {
		label = giu.Label(c.CENTER_X + "2")
	} else if left > 1*time.Second {
		label = giu.Label(c.CENTER_X + "1")
	} else {
		label = giu.Label(c.CENTER_X + "GO")
	}
	return label
}

func getKeyWidget(in string) []KeyWidget {
	layouts := []KeyWidget{}
	for _, key := range in {
		keyWidget := KeyWidget{0, 0, key, c.GRAY}
		layouts = append(layouts, keyWidget)
	}
	return layouts
}

func getGameWidgets(w []KeyWidget) []giu.Widget {
	layouts := []giu.Widget{}
	tick := int(time.Until(timer).Seconds())
	myPlayer := GetMyPlayer()
	if !myPlayer.IsDead && !delayMissilePlayer() {
		layouts = append(layouts, getKeyWidgets(w)...)
	}
	layouts = append(layouts, getTimeWidget(tick))
	layouts = append(layouts, getMyWpmWidget(tick))
	layouts = append(layouts, getSpriteWidgets()...)
	return layouts
}

func getMyWpmWidget(tick int) giu.Widget {
	return giu.Style().SetFontSize(30).To(&WpmWidget{getMyWpm(tick), 8, c.HEIGHT - 40})
}

func getTimeWidget(tick int) giu.Widget {
	return giu.Style().SetFontSize(30).To(&WpmWidget{tick, c.WIDTH - 40, c.HEIGHT - 40})
}

func getSpriteWidget(id string, size float32) giu.Widget {
	player := GetPlayer(id)
	sprite := player.getDamage()
	if id == MissileIdCurrent {
		sprite = MISSILE_SPRITE
	}
	return giu.Style().SetFontSize(17 * size).To(giu.Row(
		giu.Label(player.getDistance()),
		giu.ImageWithRgba(Sprites[sprite]).ID(strconv.Itoa(sprite)).Size(75*size, 50*size),
		giu.Label("\n"+player.Name)))
}

func getSpriteWidgets() []giu.Widget {
	layouts := []giu.Widget{}
	ids := SortedIds()
	for _, id := range ids {
		layouts = append(layouts, getSpriteWidget(id, getFitSize(len(ids))))
	}
	return layouts
}

func getCorrectKeyWidget(key KeyWidget) KeyWidget {
	return KeyWidget{key.x, key.y, key.text, c.WHITE}
}

func getIncorrectKeyWidget(key KeyWidget, char rune) KeyWidget {
	return KeyWidget{key.x, key.y, char, c.RED}
}

func getBestWidget() giu.Widget {
	updateBest()
	return giu.Style().SetFontSize(30).To(&WpmWidget{myBest_wpm, 8, c.HEIGHT - 40})
}

func getKeyWidgets(w []KeyWidget) []giu.Widget {
	layouts := []giu.Widget{}
	widgetLocX := 0
	widgetLocY := 0
	for _, key := range w {
		if widgetLocX/(c.WIDTH-40) != 0 {
			widgetLocY++
			widgetLocX = 0
		}
		keyWidget := KeyWidget{widgetLocX, widgetLocY, key.text, key.color}
		layouts = append(layouts, giu.Style().SetFontSize(30).To(&keyWidget))
		widgetLocX += keys[key.text].size
	}
	return layouts
}
