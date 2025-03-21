package termcolor

import (
	"fmt"
	// "github.com/fatih/color"
)

// Escape sequence	Text attributes
// \x1b[0m	All attributes off(color at startup)
// \x1b[1m	Bold on(enable foreground intensity)
// \x1b[4m	Underline on
// \x1b[5m	Blink on(enable background intensity)
// \x1b[21m	Bold off(disable foreground intensity)
// \x1b[24m	Underline off
// \x1b[25m	Blink off(disable background intensity)
// Escape sequence	Foreground colors
// \x1b[30m	Black
// \x1b[31m	Red
// \x1b[32m	Green
// \x1b[33m	Yellow
// \x1b[34m	Blue
// \x1b[35m	Magenta
// \x1b[36m	Cyan
// \x1b[37m	White
// \x1b[39m	Default(foreground color at startup)
// Escape sequence	Background colors
// \x1b[40m	Black
// \x1b[41m	Red
// \x1b[42m	Green
// \x1b[43m	Yellow
// \x1b[44m	Blue
// \x1b[45m	Magenta
// \x1b[46m	Cyan
// \x1b[47m	White
// \x1b[49m	Default(background color at startup)

// Base attributes
// const (
// 	Reset = iota
// 	Bold
// 	Faint
// 	Italic
// 	Underline
// 	BlinkSlow
// 	BlinkRapid
// 	ReverseVideo
// 	Concealed
// 	CrossedOut
// )

// Color is the type of color to be set.
type MyTermColor int

const (
	// No change of color
	None = MyTermColor(iota)
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func ANSI_fgcolor(fg MyTermColor) string {
	return fmt.Sprintf("\x1b[0;%dm", 30+(int)(fg-Black))
}
func ANSI_bgcolor(bg MyTermColor) string {
	return fmt.Sprintf("\x1b[0;%dm", 40+(int)(bg-Black))
}
func TermFgColor(fg MyTermColor) {
	fmt.Print(ANSI_fgcolor(fg))
}
func TermBgColor(bg MyTermColor) {
	fmt.Print(ANSI_bgcolor(bg))
}
func TermResetColor() {
	resetColor()
}
func resetColor() {
	fmt.Print("\x1b[0m")
}
func TermColor(fg MyTermColor, fgBright bool, bg MyTermColor, bgBright bool) {
	if fg == None && bg == None {
		return
	}
	s := ""
	if fg != None {
		s = fmt.Sprintf("%s%d", s, 30+(int)(fg-Black))
		if fgBright {
			s += ";1"
		}
	}
	if bg != None {
		if s != "" {
			s += ";"
		}
		s = fmt.Sprintf("%s%d", s, 40+(int)(bg-Black))
	}

	s = "\x1b[0;" + s + "m"
	print(s)
}
func BoldStr(s string) string {
	return "\x1b[1m" + s + "\x1b[0m"
}
func BlackStr(s string) string {
	return "\x1b[30m" + s + "\x1b[0m"
}
func RedStr(s string) string {
	return "\x1b[31m" + s + "\x1b[0m"
}
func GreenStr(s string) string {
	return "\x1b[32m" + s + "\x1b[0m"
}
func YellowStr(s string) string {
	return "\x1b[33m" + s + "\x1b[0m"
}
func BlueStr(s string) string {
	return "\x1b[34m" + s + "\x1b[0m"
}
func MagentaStr(s string) string {
	return "\x1b[35m" + s + "\x1b[0m"
}
func CyanStr(s string) string {
	return "\x1b[36m" + s + "\x1b[0m"
}
func WhiteStr(s string) string {
	return "\x1b[37m" + s + "\x1b[0m"
}
func UnderlineStr(s string) string {
	return "\x1b[4m" + s + "\x1b[24m"
}
func BlineStr(s string) string {
	return "\x1b[5m" + s + "\x1b[25m"
}
func PrintFgColor(fg MyTermColor, s string, a ...any) {
	ANSI_fgcolor(fg)
	fmt.Print(s, a)
	resetColor()
}
func PrintRed(s string, a ...any) {
	PrintFgColor(Red, s, a)
}
func PrintRedln(s string, a ...any) {
	PrintRed(s+"\n", a)
}
func PrintYellow(s string, a ...any) {
	PrintFgColor(Yellow, s, a)
}
func PrintYellowln(s string, a ...any) {
	PrintYellow(s+"\n", a)
}
func PrintWhite(s string, a ...any) {
	PrintFgColor(White, s, a)
}
func PrintWhiteln(s string, a ...any) {
	PrintWhite(s+"\n", a)
}
