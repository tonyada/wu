package net

import (
	"testing"
	. "wu"
)

func TestLocalMACAddr(t *testing.T) {
	mac, err := LocalMACAddr()
	Err(err)
	println("local mac addr:", mac)
}

func TestLocalMACAddrs(t *testing.T) {
	macs, err := LocalMACAddrs()
	Err(err)
	PrintStringArray(macs)
}
