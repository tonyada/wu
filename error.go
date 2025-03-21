package wu

import (
	"fmt"
	"os"
)

type WuErr struct {
	isOn bool
}

// NewWuErr creates a new WuErr with error output enabled.
func NewWuErr() *WuErr { return &WuErr{isOn: true} }

// ErrorOn enables error output.
func ErrorOn() { wuerr.isOn = true }

// ErrorOff disables error output.
func ErrorOff() { wuerr.isOn = false }

// SetErr sets the error output state.
func (w *WuErr) SetErr(b bool) { w.isOn = b }

// Err checks for an error and prints it if enabled.  Returns true if an error occurred.
func (w *WuErr) Err(err error, desc string) bool {
	if err == nil {
		return false
	}
	desc = desc + " CErr: " //Simplified conditional logic
	if w.isOn {
		wulog.logColorful("#ccc", "88", "", desc, err)
	}
	return true
}

// ErrNew creates a new error.
func ErrNew(s string) error { return fmt.Errorf(s) } //Use fmt.Errorf for better error wrapping

// Err checks for an error and prints it if enabled. Returns true if an error occurred.
func Err(err error) bool { return wuerr.Err(err, "") }

// CheckErr is an alias for Err.
func CheckErr(err error) bool { return Err(err) }

// OK checks for the absence of an error. Returns true if no error occurred.
func OK(err error) bool { return !Err(err) }

// NotErr, NoErr, Noerr are aliases for OK.
func NotErr(err error) bool { return OK(err) }
func NoErr(err error) bool  { return OK(err) }
func Noerr(err error) bool  { return OK(err) }

// ErrDesc checks for an error and prints it with a description if enabled. Returns true if an error occurred.
func ErrDesc(err error, desc string) bool { return wuerr.Err(err, desc) }

// Err2 is an alias for ErrDesc.
func Err2(err error, desc string) bool { return ErrDesc(err, desc) }

// ErrFatal checks for an error, prints it, and exits if enabled. Returns true if an error occurred.
func (w *WuErr) ErrFatal(err error, desc string) bool {
	if err == nil {
		return false
	}
	desc = desc + " ErrFatal:" //Simplified conditional logic
	if w.isOn {
		wulog.logColorful("#ccc", "124", "", desc, err)
	}
	os.Exit(1)
	return true
}

// ErrFatal checks for an error, prints it, and exits if enabled. Returns true if an error occurred.
func ErrFatal(err error) bool { return wuerr.ErrFatal(err, "") }

// ErrFatalDesc checks for an error, prints it with a description, and exits if enabled. Returns true if an error occurred.
func ErrFatalDesc(err error, desc string) bool { return wuerr.ErrFatal(err, desc) }

// ErrFatal2 is an alias for ErrFatalDesc.
func ErrFatal2(err error, desc string) bool { return ErrFatalDesc(err, desc) }

// ErrFatalExit checks for an error, prints a description, and exits.
func ErrFatalExit(err error, desc string) {
	if err != nil {
		wulog.logColorful("#ccc", "124", "", desc, err)
		os.Exit(1)
	}
}

// Global variable initialization moved to top for clarity.
var wuerr = NewWuErr()
