package gorminal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RobertoRojas/golor"
)

func Show(m string, c bool, f os.File) {
	if !c {
		return
	}
	fmt.Fprint(&f, m)
}

func Showln(m string, c bool, f os.File) {
	if !c {
		return
	}
	fmt.Fprintln(&f, m)
}

func ShowMsg(m string, l int) {
	Show(m, l <= messageLevel && !quiet, messageSTD)
}

func ShowMsgCond(m string, l int, c bool) {
	Show(m, l <= messageLevel && c && !quiet, messageSTD)
}

func ShowMsgLn(m string, l int) {
	Showln(m, l <= messageLevel && !quiet, messageSTD)
}

func ShowMsgCondln(m string, l int, c bool) {
	Showln(m, l <= messageLevel && c && !quiet, messageSTD)
}

func ShowErrMsg(m string, l int) {
	g := golor.New()
	g.Bell().Red().ColorBit256(196).
		ColorBitRGB(255, 0, 0).Add(m).Reset()
	Show(g.String(), l <= errorLevel && !quiet, errorSTD)
}

func ShowErrMsgCond(m string, l int, c bool) {
	g := golor.New()
	g.Bell().Red().ColorBit256(196).
		ColorBitRGB(255, 0, 0).Add(m).Reset()
	Show(g.String(), l <= errorLevel && c && !quiet, errorSTD)
}

func ShowErrMsgln(m string, l int) {
	g := golor.New()
	g.Bell().Red().ColorBit256(196).
		ColorBitRGB(255, 0, 0).Add(m).Reset()
	Showln(g.String(), l <= errorLevel && !quiet, errorSTD)
}

func ShowErrMsgCondln(m string, l int, c bool) {
	g := golor.New()
	g.Bell().Red().ColorBit256(196).
		ColorBitRGB(255, 0, 0).Add(m).Reset()
	Showln(g.String(), l <= errorLevel && c && !quiet, errorSTD)
}

func ShowDbgMsg(m string, l int) {
	g := golor.New()
	g.Magenta().ColorBit256(63).
		ColorBitRGB(153, 51, 255).Bold().
		Add("[DEBUG] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= debugLevel && !quiet, debugSTD)
}

func ShowDbgMsgCond(m string, l int, c bool) {
	g := golor.New()
	g.Magenta().ColorBit256(63).
		ColorBitRGB(153, 51, 255).Bold().
		Add("[DEBUG] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= debugLevel && c && !quiet, debugSTD)
}

func ShowDbgMsgln(m string, l int) {
	g := golor.New()
	g.Magenta().ColorBit256(63).
		ColorBitRGB(153, 51, 255).Bold().
		Add("[DEBUG] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= debugLevel && !quiet, debugSTD)
}

func ShowDbgMsgCondln(m string, l int, c bool) {
	g := golor.New()
	g.Magenta().ColorBit256(63).
		ColorBitRGB(153, 51, 255).Bold().
		Add("[DEBUG] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= debugLevel && c && !quiet, debugSTD)
}

func ShowWrnMsg(m string, l int) {
	g := golor.New()
	g.Yellow().ColorBit256(190).
		ColorBitRGB(255, 255, 0).Bold().
		Add("[WARNING] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= warningLevel && !quiet, warningSTD)
}

func ShowWrnMsgCond(m string, l int, c bool) {
	g := golor.New()
	g.Yellow().ColorBit256(190).
		ColorBitRGB(255, 255, 0).Bold().
		Add("[WARNING] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= warningLevel && c && !quiet, warningSTD)
}

func ShowWrnMsgln(m string, l int) {
	g := golor.New()
	g.Yellow().ColorBit256(190).
		ColorBitRGB(255, 255, 0).Bold().
		Add("[WARNING] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= warningLevel && !quiet, warningSTD)
}

func ShowWrnMsgCondln(m string, l int, c bool) {
	g := golor.New()
	g.Yellow().ColorBit256(190).
		ColorBitRGB(255, 255, 0).Bold().
		Add("[WARNING] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= warningLevel && c && !quiet, warningSTD)
}

func ShowVerbMsg(m string, l int) {
	g := golor.New()
	g.Cyan().ColorBit256(44).
		ColorBitRGB(0, 255, 255).Bold().
		Add("[VERBOSE] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= verboseLevel && !quiet, verboseSTD)
}

func ShowVerbMsgCond(m string, l int, c bool) {
	g := golor.New()
	g.Cyan().ColorBit256(44).
		ColorBitRGB(0, 255, 255).Bold().
		Add("[VERBOSE] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= verboseLevel && c && !quiet, verboseSTD)
}

func ShowVerbMsgln(m string, l int) {
	g := golor.New()
	g.Cyan().ColorBit256(44).
		ColorBitRGB(0, 255, 255).Bold().
		Add("[VERBOSE] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= verboseLevel && !quiet, verboseSTD)
}

func ShowVerbMsgCondln(m string, l int, c bool) {
	g := golor.New()
	g.Cyan().ColorBit256(44).
		ColorBitRGB(0, 255, 255).Bold().
		Add("[VERBOSE] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= verboseLevel && c && !quiet, verboseSTD)
}

func ShowOutput(o string, l int) {
	Showln(o, l == outputLevel && output, outputSTD)
}

func ShowOutputStruc(o interface{}, l int) {
	j, err := json.Marshal(o)
	if err != nil {
		ShowErrMsg(err.Error(), 0)
		return
	}
	ShowOutput(string(j), l)
}

func ShowOutputStrucIndent(o interface{}, prefix string, indent string, l int) {
	j, err := json.MarshalIndent(o, prefix, indent)
	if err != nil {
		ShowErrMsg(err.Error(), 0)
		return
	}
	ShowOutput(string(j), l)
}
