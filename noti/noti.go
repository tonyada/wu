package noti

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Shortcut for notify msg with title
// Set notification sound. Default is Glass. Possible options are Basso,
// Blow, Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr, Sosumi,
// Submarine, Tink. Check /System/Library/Sounds for available sounds.
func Noti(msg string) {
	Notification(msg, "Message", "Ping")
}
func Msg(msg string) {
	Noti(msg)
}

// notify title + msg
func Notify(title, msg string) {
	Notification(title, msg, "Ping")
}

// notify title + sub title + msg
func NotifyWithSubTitle(title, subTitle, msg string) {
	NotificationWithSubTitle(title, subTitle, msg, "Ping")
}

// full notification func with msg title and sound
func Notification(title, msg, sound string) {
	if runtime.GOOS == "darwin" {
		_, _ = exec.Command("osascript", "-e", fmt.Sprintf("display notification \"%s\" with title \"%s\" sound name \"%s\"", msg, title, sound)).Output()
	}
}

// full notification func with msg title and sound
func NotificationWithSubTitle(title, subTitle, msg, sound string) {
	if runtime.GOOS == "darwin" {
		_, _ = exec.Command("osascript", "-e", fmt.Sprintf("display notification \"%s\" with title \"%s\" subtitle \"%s\" sound name \"%s\"", msg, title, subTitle, sound)).Output()
	}
}

// osascript -e 'display notification "ha" with title "Title" subtitle "subtitle" sound name "Tink"'
