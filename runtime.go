package wu

import (
	"runtime"
)

const (
	// OS name darwin, linux, windows, freebsd
	GOOS = runtime.GOOS
	// program's architecture target: 386, amd64, or arm.
	GOARCH = runtime.GOARCH
)

// GC runs a garbage collection and blocks the caller until the garbage collection is complete.
// !!! It may also block the entire program.
func GC() { runtime.GC() }

// GOROOT returns the root of the Go tree. It uses the GOROOT environment variable,
// if set, or else the root used during the Go build.
func GOROOT() string { return runtime.GOROOT() }

// Goexit terminates the goroutine that calls it. No other goroutine is affected.
// Goexit runs all deferred calls before terminating the goroutine.
// Because Goexit is not panic, however, any recover calls in those deferred functions will return nil.
// Calling Goexit from the main goroutine terminates that goroutine without func main returning.
// Since func main has not returned, the program continues execution of other goroutines.
// If all other goroutines exit, the program crashes.
func Goexit() { runtime.Goexit() }

// Gosched yields the processor, allowing other goroutines to run.
// It does not suspend the current goroutine, so execution resumes automatically.
func Gosched() { runtime.Gosched() }

// LockOSThread wires the calling goroutine to its current operating system thread.
// Until the calling goroutine exits or calls UnlockOSThread, it will always execute in that thread, and no other goroutine can.
func LockOSThread() { runtime.LockOSThread() }

// UnlockOSThread unwires the calling goroutine from its fixed operating system thread.
// If the calling goroutine has not called LockOSThread, UnlockOSThread is a no-op.
func UnlockOSThread() { runtime.UnlockOSThread() }

// golang version like "go1.5.1"
// Version returns the Go tree's version string.
// It is either the commit hash and date at the time of the build or, when possible, a release tag like "go1.3".
func GoVersion() string { return runtime.Version() }

// NumCPU returns the number of logical CPUs usable by the current process.
func NumCPU() int { return runtime.NumCPU() }

// NumGoroutine returns the number of goroutines that currently exist.
func NumGoroutine() int { return runtime.NumGoroutine() }

// Run Max CPU (since go1.5 default is run max cpus then probably is not useful anymore)
func RunMaxCPUs() { runtime.GOMAXPROCS(runtime.NumCPU()) }
