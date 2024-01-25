package wu

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"wu/structEx"

	"github.com/muesli/termenv"
)

const (
	LOG_PRINT = iota
	LOG_NOTE
	LOG_ERROR
	LOG_INFO
	LOG_WARNING
	LOG_DANGER
	LOG_SUCCESS

	// default log flags
	LstdFlags = log.Ldate | log.Ltime
	// more: log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lshortfile
	LstdDebugFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
)

// log.* fmt.* err != nil checker's bool
type WuLog struct {
	Logger    *log.Logger // go default log
	prefix    string      // log's prefix
	calldepth int         // log 深度 默认是2，但我又重载了所以是3
	isOutput  bool        // 是否输出的bool
	// isDebug
}

// retrieve color profile supported by terminal
var p = termenv.ColorProfile()

// init go log with default flags
var TLogger = log.New(os.Stderr, "", LstdFlags)

// init my default WuLog for easy use
var wulog = &WuLog{Logger: TLogger, prefix: "", calldepth: 4, isOutput: true}

func NewWuLog(logger *log.Logger) *WuLog {
	return &WuLog{Logger: logger, prefix: "", calldepth: 4, isOutput: true}
}

// control all log output on/off
func WuLogOn()  { wulog.isOutput = true }
func WuLogOff() { wulog.isOutput = false }

// debug, output file + line
func (w *WuLog) DebugOn() { w.SetLogFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile) }
func LogDebugOn()         { wulog.SetLogFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile) }

func (w *WuLog) DebugOff() { w.SetLogFlags(log.Ldate | log.Ltime) }
func LogDebugOff()         { wulog.SetLogFlags(log.Ldate | log.Ltime) }

// set log flags with: log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lshortfile
func (w *WuLog) SetLogFlags(flag int) { w.Logger.SetFlags(flag) }
func SetLogFlags(flag int)            { wulog.SetLogFlags(flag) }

// set log prefix
func (w *WuLog) SetLogPrefix(s string) { w.Logger.SetPrefix(s) }
func SetLogPrefix(s string)            { wulog.SetLogPrefix(s) }

// set log output bool
func (w *WuLog) SetLog(b bool) { w.isOutput = b }
func SetLog(b bool)            { wulog.SetLog(b) }
func LogOn()                   { wulog.SetLog(true) }
func LogOff()                  { wulog.SetLog(false) }

// root custom logger
// log with terminal colors
// fg supports hex values will automatically degrade colors on terminals not supporting RGB
// bg supports ANSI colors (0-255)
func (w *WuLog) logColorful(fg, bg, prefix string, s ...interface{}) {
	if !w.isOutput {
		return
	}
	output := ""
	// get all function names
	allFuncNames := "\n────────────────────────────────────────────────────────────────────────────────────────────────────\n"
	funcCounter := 0
	for i := 5; i > 0; i-- {
		pc, file, line, _ := runtime.Caller(i)
		funcName := runtime.FuncForPC(pc).Name()
		// runtime.main or runtime.goexit
		if funcName == "" || funcName == "main.main" || strings.Contains(funcName, "reflect.") || strings.Contains(funcName, "wu.") || strings.Contains(funcName, "runtime.") {
			continue
		}
		allFuncNames += fmt.Sprintf("%d [%s:%d]\t -> %s()\n", funcCounter+1, file, line, funcName)
		funcCounter++
	}

	if prefix == "" { // no prefix then print out log with color
		colorOut := termenv.String(" " + fmt.Sprint(s...) + " ")
		colorOut = colorOut.Foreground(p.Color(fg))
		colorOut = colorOut.Background(p.Color(bg))
		output = fmt.Sprintf("%s %s", allFuncNames, fmt.Sprintf(" %v", colorOut))

	} else { // only print out color prefix
		colorPrefix := termenv.String(prefix)
		colorPrefix = colorPrefix.Foreground(p.Color(fg))
		colorPrefix = colorPrefix.Background(p.Color(bg))
		// print out with colorful prefix
		output = fmt.Sprintf(" %v %v", colorPrefix, fmt.Sprint(s...))
	}
	output += "\n────────────────────────────────────────────────────────────────────────────────────────────────────\n"
	w.Logger.Output(w.calldepth, output)
}
func LogWithColor(fg, bg string, s ...interface{}) { wulog.logColorful(fg, bg, "", s...) }
func LogWithColorPrefix(fg, bg, prefix string, s ...interface{}) {
	wulog.logColorful(fg, bg, prefix, s...)
}

// log
func (w *WuLog) LogType(print_type int, s ...interface{}) {
	if !w.isOutput {
		return
	}
	switch print_type {
	case LOG_NOTE:
		w.logColorful("#000", "#eee", "", s...)
	case LOG_ERROR:
		w.logColorful("#fff", "124", "", s...)
	case LOG_INFO:
		w.logColorful("#fff", "26", "", s...)
	case LOG_WARNING:
		w.logColorful("#000", "227", "", s...)
	case LOG_DANGER:
		w.logColorful("#fff", "160", "", s...)
	case LOG_SUCCESS:
		w.logColorful("#000", "118", "", s...)
	default:
		w.logColorful("", "", "", s...)
	}
}
func Log(s ...interface{})     { wulog.logColorful("", "", "", s...) }
func MustLog(s ...interface{}) { wulog.Logger.Output(wulog.calldepth, fmt.Sprint(s...)) }

// logs
func LogNote(s ...interface{})    { wulog.LogType(LOG_NOTE, s...) }
func LogErr(s ...interface{})     { wulog.LogType(LOG_ERROR, s...) }
func LogInfo(s ...interface{})    { wulog.LogType(LOG_INFO, s...) }
func LogWarn(s ...interface{})    { wulog.LogType(LOG_WARNING, s...) }
func LogDanger(s ...interface{})  { wulog.LogType(LOG_DANGER, s...) }
func LogSuccess(s ...interface{}) { wulog.LogType(LOG_SUCCESS, s...) }
func LogOK(s ...interface{})      { wulog.LogType(LOG_SUCCESS, s...) }

func LogWithPrefix(prefix string, s ...interface{}) {
	wulog.logColorful("#000", "191", "["+prefix+"]", s...)
}

func LogInfoPrefix(s ...interface{})    { wulog.logColorful("#fff", "26", "[INFO]", s...) }
func LogWarnPrefix(s ...interface{})    { wulog.logColorful("#000", "226", "[WARN]", s...) }
func LogDangerPrefix(s ...interface{})  { wulog.logColorful("#fff", "160", "[DAGR]", s...) }
func LogSuccessPrefix(s ...interface{}) { wulog.logColorful("#000", "118", "[SUCC]", s...) }
func MustLogln(s ...interface{})        { wulog.Logger.Output(wulog.calldepth, fmt.Sprintln(s...)) }

// Logf
func (w *WuLog) Logf(format string, s ...interface{}) {
	if w.isOutput {
		w.Logger.Output(w.calldepth, fmt.Sprintf(format, s...))
	}
}
func Logf(format string, s ...interface{}) { wulog.Logf(format, s...) }

// MustLogf
func (w *WuLog) MustLogf(format string, s ...interface{}) {
	w.Logger.Output(wulog.calldepth, fmt.Sprintf(format, s...))
}
func MustLogf(format string, s ...interface{}) { wulog.MustLogf(format, s...) }

// Logfln
func (w *WuLog) MustLogfln(format string, s ...interface{}) {
	w.Logger.Output(wulog.calldepth, fmt.Sprintf(format+"\n", s...))
}
func Logfln(format string, s ...interface{}) { wulog.Logf(format+"\n", s...) }

// log.Fatalf
func (w *WuLog) Fatalf(format string, s ...interface{}) {
	w.Logger.Output(w.calldepth, fmt.Sprintf(format, s...))
	os.Exit(1)
}
func Fatalf(format string, s ...interface{}) { wulog.Fatalf(format, s...) }

func LogStruct(st interface{}) { wulog.LogStruct(st) }

func (w *WuLog) LogStruct(st interface{}) {
	w.logColorful("#fff", "26", "", "LogStruct:")
	// w.Logger.Output(w.calldepth-1, "\n")
	structEx.Explicit(reflect.ValueOf(st), 0)
}
