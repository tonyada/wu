package wu

import (
	"fmt"
	"io"
)

type WuPrint struct {
	isOn bool
}

// init WuPrint std
var wuprint = &WuPrint{isOn: true}

func NewWuPrint() *WuPrint { return &WuPrint{isOn: true} }

// control all prints output on/off
func PrintOn()  { wuprint.isOn = true }
func PrintOff() { wuprint.isOn = false }

// set print output bool
func (w *WuPrint) SetPrint(b bool) { w.isOn = b }
func SetPrint(b bool)              { wuprint.SetPrint(b) }
func SetPrintOn()                  { wuprint.SetPrint(true) }
func SetPrintOff()                 { wuprint.SetPrint(false) }

// fmt.print
func (w *WuPrint) Print(print_type string, a ...interface{}) {
	if !w.isOn {
		return
	}
	switch print_type {
	case "print":
		fmt.Print(fmt.Sprint(a...))
	case "println":
		fmt.Println(fmt.Sprint(a...))
	default:
		fmt.Println(fmt.Sprint(a...))
	}
}
func Print(a ...interface{})     { wuprint.Print("print", a...) }
func MustPrint(a ...interface{}) { fmt.Print(fmt.Sprint(a...)) }

func Println(a ...interface{})     { wuprint.Print("println", a...) }
func MustPrintln(a ...interface{}) { fmt.Println(fmt.Sprint(a...)) }

// fmt.Printf
func (w *WuPrint) Printf(format string, a ...interface{}) {
	if w.isOn {
		fmt.Printf(format, a...)
	}
}
func Printf(format string, a ...interface{})       { wuprint.Printf(format, a...) }
func MustPrintf(format string, a ...interface{})   { fmt.Printf(format, a...) }
func Printfln(format string, a ...interface{})     { wuprint.Printf(format+"\n", a...) }
func MustPrintfln(format string, a ...interface{}) { fmt.Println(fmt.Sprintf(format, a...)) }

// fmt.Fprintf
func (w *WuPrint) Fprintf(o io.Writer, format string, a ...interface{}) {
	if w.isOn {
		fmt.Fprintf(o, format, a...)
	}
}
func Fprintf(w io.Writer, format string, a ...interface{})     { wuprint.Fprintf(w, format, a...) }
func MustFprintf(w io.Writer, format string, a ...interface{}) { fmt.Fprintf(w, format, a...) }

func Newline()   { println("") }
func New2lines() { println("\n") }
func Newlines(lines int) {
	for i := 0; i < lines; i++ {
		Newline()
	}
}

func PrintByte(b []byte) {
	print("[]byte{")
	for i := 0; i < len(b); i++ {
		if i != len(b)-1 {
			print(Sprintf("%v,", b[i]))
		} else {
			// last byte without ,
			print(Sprintf("%v", b[i]))
		}
	}
	println("}")
}

func PrintStringArray(s []string) {
	for i, v := range s {
		println(i, v)
	}
}

// fmt.Scanf
func Scanf(format string, a ...interface{}) (int, error) { return fmt.Scanf(format, a...) }

// Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
func Scanln(a ...interface{}) (int, error) { return fmt.Scanln(a...) }

// fmt.Sprint
func Sprint(a ...interface{}) string { return fmt.Sprint(a...) }

// fmt.Sprintf
func Sprintf(format string, a ...interface{}) string { return fmt.Sprintf(format, a...) }

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func Sprintln(a ...interface{}) string { return fmt.Sprintln(a...) }

// fmt.Errorf
func Errorf(format string, a ...interface{}) error { return fmt.Errorf(format, a...) }
