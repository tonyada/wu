package wu

import (
	"fmt"
	"io"
)

type WuPrint struct {
	isOn bool
}

// init WuPrint std; using a package-level variable is generally fine here.
var wuprint = &WuPrint{isOn: true}

// NewWuPrint creates a new WuPrint instance.  Consider if this is really needed.
func NewWuPrint() *WuPrint { return &WuPrint{isOn: true} }

// Control all prints output on/off.  Could be simplified.
func PrintOn()  { wuprint.isOn = true }
func PrintOff() { wuprint.isOn = false }

// SetPrint sets the print output bool.  Could be combined with PrintOn/PrintOff.
func (w *WuPrint) SetPrint(b bool) { w.isOn = b }
func SetPrint(b bool)              { wuprint.SetPrint(b) }
func SetPrintOn()                  { wuprint.SetPrint(true) }
func SetPrintOff()                 { wuprint.SetPrint(false) }

// Print handles both Print and Println functionality.  Simplified.
func (w *WuPrint) Print(println bool, a ...any) {
	if !w.isOn {
		return
	}
	if println {
		fmt.Println(fmt.Sprint(a...))
	} else {
		fmt.Print(fmt.Sprint(a...))
	}
}

func Print(a ...any)     { wuprint.Print(false, a...) }
func MustPrint(a ...any) { fmt.Print(fmt.Sprint(a...)) }

func Println(a ...any)     { wuprint.Print(true, a...) }
func MustPrintln(a ...any) { fmt.Println(fmt.Sprint(a...)) }

// Printf handles formatted printing.  No change needed here.
func (w *WuPrint) Printf(format string, a ...any) {
	if w.isOn {
		fmt.Printf(format, a...)
	}
}
func Printf(format string, a ...any)       { wuprint.Printf(format, a...) }
func MustPrintf(format string, a ...any)   { fmt.Printf(format, a...) }
func Printfln(format string, a ...any)     { wuprint.Printf(format+"\n", a...) }
func MustPrintfln(format string, a ...any) { fmt.Println(fmt.Sprintf(format, a...)) }

// Fprintf handles formatted printing to an io.Writer. No change needed here.
func (w *WuPrint) Fprintf(o io.Writer, format string, a ...any) {
	if w.isOn {
		fmt.Fprintf(o, format, a...)
	}
}
func Fprintf(w io.Writer, format string, a ...any)     { wuprint.Fprintf(w, format, a...) }
func MustFprintf(w io.Writer, format string, a ...any) { fmt.Fprintf(w, format, a...) }

// Newline functions; these are fine as they are.
func Newline()   { println() }
func New2lines() { println("\n") }
func Newlines(lines int) {
	for range lines {
		Newline()
	}
}

// PrintByte prints a byte slice.  Improved for readability and efficiency.
func PrintByte(b []byte) {
	fmt.Printf("[]byte{")
	for i, v := range b {
		fmt.Printf("%v", v)
		if i < len(b)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println("}")
}

// PrintStringArray prints a string array.  No change needed here.
func PrintStringArray(s []string) {
	for i, v := range s {
		fmt.Println(i, v)
	}
}

// Wrapper functions for fmt package functions; these are fine as they are.
func Scanf(format string, a ...any) (int, error) { return fmt.Scanf(format, a...) }
func Scanln(a ...any) (int, error)               { return fmt.Scanln(a...) }
func Sprint(a ...any) string                     { return fmt.Sprint(a...) }
func Sprintf(format string, a ...any) string     { return fmt.Sprintf(format, a...) }
func Sprintln(a ...any) string                   { return fmt.Sprintln(a...) }
func Errorf(format string, a ...any) error       { return fmt.Errorf(format, a...) }
