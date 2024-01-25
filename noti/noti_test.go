package noti

import (
	"testing"
)

// Shortcut for notify msg with title
// Set notification sound. Default is Glass. Possible options are Basso,
// Blow, Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr, Sosumi,
// Submarine, Tink. Check /System/Library/Sounds for available sounds.
func TestNotification(t *testing.T) {
	Notification("test_msg", "test_title", "Frog")
	Noti("test2")
	Notify("test_msg3", "test_title3")
}
