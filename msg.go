package gorminal

import (
	"fmt"
	"os"

	"github.com/RobertoRojas/golor"
)

func Show(m string, c bool, f *os.File) {
	if !c {
		return
	}
	fmt.Fprint(f, m)
}

func Showln(m string, c bool, f *os.File) {
	if !c {
		return
	}
	fmt.Fprintln(f, m)
}

func ShowMsg(m string, l int) {
	Show(m, l <= messageLevel, os.Stdout)
}

func ShowMsgCond(m string, l int, c bool) {
	Show(m, l <= messageLevel && c, os.Stdout)
}

func ShowMsgLn(m string, l int) {
	Showln(m, l <= messageLevel, os.Stdout)
}

func ShowMsgCondln(m string, l int, c bool) {
	Showln(m, l <= messageLevel && c, os.Stdout)
}

func ShowErrMsg(m string, l int) {
	g := golor.New()
	g.Bell().Red().ColorBit256(196).
		ColorBitRGB(255, 0, 0).Add(m).Reset()
	Show(g.String(), l <= errorLevel, os.Stderr)
}

func ShowErrMsgCond(m string, l int, c bool) {
	g := golor.New()
	g.Bell().Red().ColorBit256(196).
		ColorBitRGB(255, 0, 0).Add(m).Reset()
	Show(g.String(), l <= errorLevel && c, os.Stderr)
}

func ShowErrMsgln(m string, l int) {
	g := golor.New()
	g.Bell().Red().ColorBit256(196).
		ColorBitRGB(255, 0, 0).Add(m).Reset()
	Showln(g.String(), l <= errorLevel, os.Stderr)
}

func ShowErrMsgCondln(m string, l int, c bool) {
	g := golor.New()
	g.Bell().Red().ColorBit256(196).
		ColorBitRGB(255, 0, 0).Add(m).Reset()
	Showln(g.String(), l <= errorLevel && c, os.Stderr)
}

func ShowDbgMsg(m string, l int) {
	g := golor.New()
	g.Magenta().ColorBit256(63).
		ColorBitRGB(153, 51, 255).Bold().
		Add("[DEBUG] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= debugLevel, os.Stdout)
}

func ShowDbgMsgCond(m string, l int, c bool) {
	g := golor.New()
	g.Magenta().ColorBit256(63).
		ColorBitRGB(153, 51, 255).Bold().
		Add("[DEBUG] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= debugLevel && c, os.Stdout)
}

func ShowDbgMsgln(m string, l int) {
	g := golor.New()
	g.Magenta().ColorBit256(63).
		ColorBitRGB(153, 51, 255).Bold().
		Add("[DEBUG] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= debugLevel, os.Stdout)
}

func ShowDbgMsgCondln(m string, l int, c bool) {
	g := golor.New()
	g.Magenta().ColorBit256(63).
		ColorBitRGB(153, 51, 255).Bold().
		Add("[DEBUG] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= debugLevel && c, os.Stdout)
}

func ShowWrnMsg(m string, l int) {
	g := golor.New()
	g.Yellow().ColorBit256(190).
		ColorBitRGB(255, 255, 0).Bold().
		Add("[WARNING] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= warningLevel, os.Stdout)
}

func ShowWrnMsgCond(m string, l int, c bool) {
	g := golor.New()
	g.Yellow().ColorBit256(190).
		ColorBitRGB(255, 255, 0).Bold().
		Add("[WARNING] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= warningLevel && c, os.Stdout)
}

func ShowWrnMsgln(m string, l int) {
	g := golor.New()
	g.Yellow().ColorBit256(190).
		ColorBitRGB(255, 255, 0).Bold().
		Add("[WARNING] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= warningLevel, os.Stdout)
}

func ShowWrnMsgCondln(m string, l int, c bool) {
	g := golor.New()
	g.Yellow().ColorBit256(190).
		ColorBitRGB(255, 255, 0).Bold().
		Add("[WARNING] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= warningLevel && c, os.Stdout)
}

func ShowVerbMsg(m string, l int) {
	g := golor.New()
	g.Cyan().ColorBit256(44).
		ColorBitRGB(0, 255, 255).Bold().
		Add("[VERBOSE] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= verboseLevel, os.Stdout)
}

func ShowVerbMsgCond(m string, l int, c bool) {
	g := golor.New()
	g.Cyan().ColorBit256(44).
		ColorBitRGB(0, 255, 255).Bold().
		Add("[VERBOSE] ").BoldReset().Add(m).Reset()
	Show(g.String(), l <= verboseLevel && c, os.Stdout)
}

func ShowVerbMsgln(m string, l int) {
	g := golor.New()
	g.Cyan().ColorBit256(44).
		ColorBitRGB(0, 255, 255).Bold().
		Add("[VERBOSE] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= verboseLevel, os.Stdout)
}

func ShowVerbMsgCondln(m string, l int, c bool) {
	g := golor.New()
	g.Cyan().ColorBit256(44).
		ColorBitRGB(0, 255, 255).Bold().
		Add("[VERBOSE] ").BoldReset().Add(m).Reset()
	Showln(g.String(), l <= verboseLevel && c, os.Stdout)
}
