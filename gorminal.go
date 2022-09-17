package gorminal

import (
	"github.com/RobertoRojas/golor"
)

var colorMode golor.GolorMode
var colorEnv bool

var soundMode golor.GolorSound
var soundEnv bool

var OverrideEnv bool

func SetColorMode(m golor.GolorMode) {
	if colorEnv && !OverrideEnv {
		return
	}
	colorMode = m
}

func SetSoundMode(m golor.GolorSound) {
	if soundEnv && !OverrideEnv {
		return
	}
	soundMode = m
}
