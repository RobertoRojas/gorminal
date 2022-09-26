package gorminal

import (
	"os"

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

var cols int
var rows int

func updateTermSize() {
	cols, rows = consolesize.GetConsoleSize()
}

func getTermSize() (int, int) {
	if !dynamicSize {
		return defTermSize, 0
	}
	if cols == 0 && rows == 0 {
		updateTermSize()
	}
	return cols, rows
}

var quiet bool
var quietEnv bool

func SetQuiet(d bool) {
	if quietEnv && !OverrideEnv {
		return
	}
	quiet = d
}

var output bool
var outputEnv bool

func SetOutput(d bool) {
	if outputEnv && !OverrideEnv {
		return
	}
	output = d
}

var messageLevel int
var messageLevelEnv bool

func SetMessageLevel(l int) {
	if messageLevelEnv && !OverrideEnv {
		return
	}
	messageLevel = l
}

var messageSTD os.File
var messageSTDEnv bool

func SetMessageSTD(f os.File) {
	if messageSTDEnv && !OverrideEnv {
		return
	}
	messageSTD = f
}

var errorLevel int
var errorLevelEnv bool

func SetErrorLevel(l int) {
	if errorLevelEnv && !OverrideEnv {
		return
	}
	errorLevel = l
}

var errorSTD os.File
var errorSTDEnv bool

func SetErrorSTD(f os.File) {
	if errorSTDEnv && !OverrideEnv {
		return
	}
	errorSTD = f
}

var debugLevel int
var debugLevelEnv bool

func SetDebugLevel(l int) {
	if debugLevelEnv && !OverrideEnv {
		return
	}
	debugLevel = l
}

var debugSTD os.File
var debugSTDEnv bool

func SetDebugSTD(f os.File) {
	if debugLevelEnv && !OverrideEnv {
		return
	}
	debugSTD = f
}

var warningLevel int
var warningLevelEnv bool

func SetWarningLevel(l int) {
	if warningLevelEnv && !OverrideEnv {
		return
	}
	warningLevel = l
}

var warningSTD os.File
var warningSTDEnv bool

func SetWarningSTD(f os.File) {
	if warningSTDEnv && !OverrideEnv {
		return
	}
	warningSTD = f
}

var verboseLevel int
var verboseLevelEnv bool

func SetVerboseLevel(l int) {
	if verboseLevelEnv && !OverrideEnv {
		return
	}
	verboseLevel = l
}

var verboseSTD os.File
var verboseSTDEnv bool

func SetVerboseSTD(f os.File) {
	if verboseSTDEnv && !OverrideEnv {
		return
	}
	verboseSTD = f
}

var outputLevel int
var outputLevelEnv bool

func SetOutputLevel(l int) {
	if outputLevelEnv && !OverrideEnv {
		return
	}
	outputLevel = l
}

var outputSTD os.File
var outputSTDEnv bool

func SetOutputSTD(f os.File) {
	if outputSTDEnv && !OverrideEnv {
		return
	}
	outputSTD = f
}
