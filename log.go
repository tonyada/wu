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
var WuLogger = log.New(os.Stderr, "", LstdFlags)

// init my default wulog for easy use
var wulog = NewWuLog(WuLogger)

// NewWuLog creates a new WuLog instance.
func NewWuLog(logger *log.Logger) *WuLog {
	return &WuLog{
		Logger:    logger,
		prefix:    "",
		calldepth: 4,
		isOutput:  true,
	}
}

// WuLogOn enables all log output.
func WuLogOn() { wulog.isOutput = true }

// WuLogOff disables all log output.
func WuLogOff() { wulog.isOutput = false }

// DebugOn enables debug mode, including file and line information.
func (w *WuLog) DebugOn() { w.SetLogFlags(LstdDebugFlags) }

// LogDebugOn enables debug mode for the default logger.
func LogDebugOn() { wulog.DebugOn() }

// DebugOff disables debug mode, reverting to standard log flags.
func (w *WuLog) DebugOff() { w.SetLogFlags(LstdFlags) }

// LogDebugOff disables debug mode for the default logger.
func LogDebugOff() { wulog.DebugOff() }

// SetLogFlags sets the log flags for the logger.
func (w *WuLog) SetLogFlags(flag int) { w.Logger.SetFlags(flag) }

// SetLogFlags sets the log flags for the default logger.
func SetLogFlags(flag int) { wulog.SetLogFlags(flag) }

// SetLogPrefix sets the log prefix for the logger.
func (w *WuLog) SetLogPrefix(s string) { w.Logger.SetPrefix(s) }

// SetLogPrefix sets the log prefix for the default logger.
func SetLogPrefix(s string) { wulog.SetLogPrefix(s) }

// SetLog enables or disables log output for the logger.
func (w *WuLog) SetLog(b bool) { w.isOutput = b }

// SetLog enables or disables log output for the default logger.
func SetLog(b bool) { wulog.SetLog(b) }

// LogOn enables log output for the default logger.
func LogOn() { wulog.SetLog(true) }

// LogOff disables log output for the default logger.
func LogOff() { wulog.SetLog(false) }

// logColorful logs a message with terminal colors.
func (w *WuLog) logColorful(fg, bg, prefix string, s ...any) {
	if !w.isOutput {
		return
	}

	output := ""
	// get all function names
	allFuncNames := "\n────────────────────────────────────────────────────────────────────────────────────────────────────\n"
	funcCounter := 0
	for i := 3; i < 8; i++ { // Optimized loop range
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
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

// LogWithColor logs a message with specified foreground and background colors.
func LogWithColor(fg, bg string, s ...any) { wulog.logColorful(fg, bg, "", s...) }

// LogWithColorPrefix logs a message with a colored prefix and specified colors.
func LogWithColorPrefix(fg, bg, prefix string, s ...any) {
	wulog.logColorful(fg, bg, prefix, s...)
}

// LogType logs a message with a predefined color scheme based on the log type.
func (w *WuLog) LogType(print_type int, s ...any) {
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

// Log logs a message with default settings.
func Log(s ...any) { wulog.logColorful("", "", "", s...) }

// MustLog logs a message without color and always outputs.
func MustLog(s ...any) { wulog.Logger.Output(wulog.calldepth, fmt.Sprint(s...)) }

// LogNote logs a note message.
func LogNote(s ...any) { wulog.LogType(LOG_NOTE, s...) }

// LogErr logs an error message.
func LogErr(s ...any) { wulog.LogType(LOG_ERROR, s...) }

// LogInfo logs an informational message.
func LogInfo(s ...any) { wulog.LogType(LOG_INFO, s...) }

// LogWarn logs a warning message.
func LogWarn(s ...any) { wulog.LogType(LOG_WARNING, s...) }

// LogDanger logs a danger message.
func LogDanger(s ...any) { wulog.LogType(LOG_DANGER, s...) }

// LogSuccess logs a success message.
func LogSuccess(s ...any) { wulog.LogType(LOG_SUCCESS, s...) }

// LogOK is an alias for LogSuccess.
func LogOK(s ...any) { wulog.LogType(LOG_SUCCESS, s...) }

// LogWithPrefix logs a message with a custom prefix.
func LogWithPrefix(prefix string, s ...any) {
	wulog.logColorful("#000", "191", "["+prefix+"]", s...)
}

// LogInfoPrefix logs an informational message with a "[INFO]" prefix.
func LogInfoPrefix(s ...any) { wulog.logColorful("#fff", "26", "[INFO]", s...) }

// LogWarnPrefix logs a warning message with a "[WARN]" prefix.
func LogWarnPrefix(s ...any) { wulog.logColorful("#000", "226", "[WARN]", s...) }

// LogDangerPrefix logs a danger message with a "[DAGR]" prefix.
func LogDangerPrefix(s ...any) { wulog.logColorful("#fff", "160", "[DAGR]", s...) }

// LogSuccessPrefix logs a success message with a "[SUCC]" prefix.
func LogSuccessPrefix(s ...any) { wulog.logColorful("#000", "118", "[SUCC]", s...) }

// MustLogln logs a message with a newline without color and always outputs.
func MustLogln(s ...any) { wulog.Logger.Output(wulog.calldepth, fmt.Sprintln(s...)) }

// Logf logs a formatted message.
func (w *WuLog) Logf(format string, s ...any) {
	if w.isOutput {
		w.Logger.Output(w.calldepth, fmt.Sprintf(format, s...))
	}
}

// Logf logs a formatted message for the default logger.
func Logf(format string, s ...any) { wulog.Logf(format, s...) }

// MustLogf logs a formatted message without color and always outputs.
func (w *WuLog) MustLogf(format string, s ...any) {
	w.Logger.Output(wulog.calldepth, fmt.Sprintf(format, s...))
}

// MustLogf logs a formatted message for the default logger.
func MustLogf(format string, s ...any) { wulog.MustLogf(format, s...) }

// MustLogfln logs a formatted message with a newline without color and always outputs.
func (w *WuLog) MustLogfln(format string, s ...any) {
	w.Logger.Output(wulog.calldepth, fmt.Sprintf(format+"\n", s...))
}

// Logfln logs a formatted message with a newline.
func Logfln(format string, s ...any) { wulog.Logf(format+"\n", s...) }

// Fatalf logs a formatted message and then exits the program.
func (w *WuLog) Fatalf(format string, s ...any) {
	w.Logger.Output(w.calldepth, fmt.Sprintf(format, s...))
	os.Exit(1)
}

// Fatalf logs a formatted message for the default logger and then exits the program.
func Fatalf(format string, s ...any) { wulog.Fatalf(format, s...) }

// LogStruct logs the details of a struct.
func LogStruct(st any) { wulog.LogStruct(st) }

// LogStruct logs the details of a struct with color.
func (w *WuLog) LogStruct(st any) {
	w.logColorful("#fff", "26", "", "LogStruct:")
	structEx.Explicit(reflect.ValueOf(st), 0)
}
