package gorminal

import (
	"github.com/RobertoRojas/golor"
	"github.com/nathan-fiscaletti/consolesize-go"
)

var OverrideEnv bool

var colorEnv bool

func SetColorMode(m golor.GolorMode) {
	if colorEnv && !OverrideEnv {
		return
	}
	golor.Mode = m
}

var soundEnv bool

func SetSoundMode(m golor.GolorSound) {
	if soundEnv && !OverrideEnv {
		return
	}
	golor.Sound = m
}

var dynamicSize bool
var dynamicSizeEnv bool

func SetDynamicSize(d bool) {
	if dynamicSizeEnv && !OverrideEnv {
		return
	}
	dynamicSizeEnv = d
}

var defTermSize int
var defTermSizeEnv bool

func SetDefTermSize(s int) {
	if s < 1 {
		return
	}
	defTermSize = s
}

func getTermSize() (int, int) {
	if !dynamicSize {
		return int(defTermSize), 0
	}
	return consolesize.GetConsoleSize()
}

var messageLevel int
var messageLevelEnv bool

func SetMessageLevel(l int) {
	if messageLevelEnv && !OverrideEnv {
		return
	}
	messageLevel = l
}

var errorLevel int
var errorLevelEnv bool

func SetErrorLevel(l int) {
	if errorLevelEnv && !OverrideEnv {
		return
	}
	errorLevel = l
}

var debugLevel int
var debugLevelEnv bool

func SetDebugLevel(l int) {
	if debugLevelEnv && !OverrideEnv {
		return
	}
	debugLevel = l
}

var warningLevel int
var warningLevelEnv bool

func SetWarningLevel(l int) {
	if warningLevelEnv && !OverrideEnv {
		return
	}
	warningLevel = l
}

var verboseLevel int
var verboseLevelEnv bool

func SetVerboseLevel(l int) {
	if verboseLevelEnv && !OverrideEnv {
		return
	}
	verboseLevel = l
}

var outputLevel int
var outputLevelEnv bool

func SetOutputLevel(l int) {
	if outputLevelEnv && !OverrideEnv {
		return
	}
	outputLevel = l
}
