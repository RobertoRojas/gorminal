package gorminal

import (
	"os"
	"strconv"
	"strings"

	"github.com/RobertoRojas/golor"
)

func getIntFromEnv(eVar string, def int) (int, bool) {
	if v := os.Getenv(eVar); len(v) > 0 {
		u, err := strconv.ParseInt(v, 10, 32)
		if err == nil {
			return int(u), true
		}
	}
	return def, false
}

func getBoolFromEnv(eVar string, def bool) (bool, bool) {
	if v := os.Getenv(eVar); len(v) > 0 {
		return strings.ToLower(v) == "true", true
	}
	return def, false
}

func init() {
	if v := os.Getenv("GORMIAL_COLOR_MODE"); len(v) > 0 {
		switch strings.ToLower(v) {
		case "bit8":
			golor.Mode = golor.Bit8
		case "bit256":
			golor.Mode = golor.Bit256
		case "bitrgb":
			golor.Mode = golor.BitRGB
		default:
			golor.Mode = golor.Plain
		}
		colorEnv = true
	} else {
		golor.Mode = golor.Plain
		colorEnv = false
	}
	if v := os.Getenv("GORMIAL_SOUND_MODE"); len(v) > 0 {
		switch strings.ToLower(v) {
		case "active":
			golor.Sound = golor.Active
		default:
			golor.Sound = golor.Mute
		}
		soundEnv = true
	} else {
		golor.Sound = golor.Mute
		soundEnv = false
	}
	dynamicSize, dynamicSizeEnv = getBoolFromEnv("GORMIAL_DYNAMIC_SIZE", true)
	defTermSize, defTermSizeEnv = getIntFromEnv("GORMIAL_DEFAULT_TERMINAL_SIZE", 58)
	if defTermSize < 1 {
		defTermSize = 58
	}
	messageLevel, messageLevelEnv = getIntFromEnv("GORMIAL_", 0)
	errorLevel, errorLevelEnv = getIntFromEnv("GORMIAL_MESSAGE_LEVEL", 0)
	debugLevel, debugLevelEnv = getIntFromEnv("GORMIAL_ERROR_LEVEL", -1)
	warningLevel, warningLevelEnv = getIntFromEnv("GORMIAL_WARNING_LEVEL", 0)
	verboseLevel, verboseLevelEnv = getIntFromEnv("GORMIAL_VERBOSE_LEVEL", -1)
	outputLevel, outputLevelEnv = getIntFromEnv("GORMIAL_OUTPUT_LEVEL", -1)
	OverrideEnv = false
}
