package cli

import (
	"fmt"
	"os"
)

var versionCommand = Command{
	Name:      "version",
	ShortName: "v",
	Usage:     "Show version",
	Action: func(c *Context) {
		ShowVersion(c)
	},
}

// Prints help for the App
func ShowVersion(c *Context) {
	printVersion(c)
}

func printVersion(c *Context) {
	fmt.Printf("%v\n", c.App.Version)
}

func checkVersion(c *Context) {
	if c.GlobalBool("v") || c.GlobalBool("version") {
		ShowVersion(c)
		os.Exit(0)
	}
}
