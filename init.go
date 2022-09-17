package gorminal

import (
	"os"
	"strings"

	"github.com/RobertoRojas/golor"
)

func init() {
	if v := os.Getenv("GORMIAL_COLOR_MODE"); len(v) > 0 {
		switch strings.ToLower(v) {
		case "bit8":
			colorMode = golor.Bit8
		case "bit256":
			colorMode = golor.Bit256
		case "bitrgb":
			colorMode = golor.BitRGB
		default:
			colorMode = golor.Plain
		}
		colorEnv = true
	} else {
		colorMode = golor.Plain
		colorEnv = false
	}
	if v := os.Getenv("GORMIAL_SOUND_MODE"); len(v) > 0 {
		switch strings.ToLower(v) {
		case "active":
			soundMode = golor.Active
		default:
			soundMode = golor.Mute
		}
		soundEnv = true
	} else {
		soundMode = golor.Mute
		soundEnv = false
	}
	OverrideEnv = false
}
