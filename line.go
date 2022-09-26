package gorminal

import (
	"strings"
)

func GetLine(text string, line byte, corner byte) (string, string, string, string) {
	cols, _ := getTermSize()
	if cols <= len(text)+1 {
		return "", "", "", text
	}
	var l_width int
	r_width := (cols - len(text) - 2) / 2
	if r_width*2+len(text)+2 == cols {
		l_width = r_width
	} else {
		l_width = r_width + 1
	}
	var rigth strings.Builder
	for i := 0; i < r_width; i++ {
		rigth.WriteByte(line)
	}
	var left strings.Builder
	for i := 0; i < l_width; i++ {
		left.WriteByte(line)
	}
	return string(corner), left.String(), rigth.String(), text
}
