package wu

import (
	"errors"
	"fmt"
	"os"
)

type WuErr struct {
	isOn bool //`json:"error_on"` // trigger print err or not
}

// init wuerr std
var wuerr = &WuErr{isOn: true}

// default is print error [isOn = true]
func NewWuErr() *WuErr { return &WuErr{isOn: true} }

// control all errors output on/off
func ErrorOn()  { wuerr.isOn = true }
func ErrorOff() { wuerr.isOn = false }

// set err output bool
func SetErr(b bool)            { wuerr.SetErr(b) }
func SetErrOn()                { wuerr.SetErr(true) }
func SetErrOff()               { wuerr.SetErr(false) }
func SeWuerrOn()               { wuerr.SetErr(true) }
func SeWuerrOff()              { wuerr.SetErr(false) }
func (w *WuErr) SetErr(b bool) { w.isOn = b }

// check err & print
// error with description
func (w *WuErr) Err(err error, desc string) bool {
	if err != nil {
		// if desc is empty then use standard Check Err:
		if desc == "" {
			desc = "CErr: "
		} else {
			desc = desc + " CErr: "
		}
		if w.isOn {
			wulog.logColorful("#ccc", "88", "", desc, err)
		}
		return true
	}
	return false
}

// New Error
func ErrNew(s string) error                      { return errors.New(s) }
func Errf(format string, a ...interface{}) error { return ErrNew(fmt.Sprintf(format, a...)) }

// check error just return bool
// if error != nil then has error
func Err(err error) bool { return wuerr.Err(err, "") }

// func isErr(err error) bool    { return Err(err) }
func CheckErr(err error) bool { return Err(err) }

// if err == nil then it is no error
func OK(err error) bool     { return !wuerr.Err(err, "") }
func NotErr(err error) bool { return OK(err) }
func NoErr(err error) bool  { return OK(err) }
func Noerr(err error) bool  { return OK(err) }

// error with desc
func ErrDesc(err error, desc string) bool { return wuerr.Err(err, desc) }
func Err2(err error, desc string) bool    { return ErrDesc(err, desc) }

// error with exit
func (w *WuErr) ErrFatal(err error, desc string) bool {
	if err != nil {
		if desc == "" {
			desc = "ErrFatal:"
		} else {
			desc = desc + " ErrFatal:"
		}
		if w.isOn {
			wulog.logColorful("#ccc", "124", "", desc, err)
		}
		os.Exit(1)
		return true
	}
	return false
}

// err with exit
func ErrFatal(err error) bool { return wuerr.ErrFatal(err, "") }

// err + desc with exit
func ErrFatalDesc(err error, desc string) bool { return wuerr.ErrFatal(err, desc) }
func ErrFatal2(err error, desc string) bool    { return ErrFatalDesc(err, desc) }

// check err and only print desc
func ErrFatalExit(err error, desc string) {
	if err != nil {
		wulog.logColorful("#ccc", "124", "", desc, err)
		os.Exit(1)
	}
}
