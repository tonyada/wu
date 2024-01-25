package cli

import (
	"fmt"
	"os"
)

type App struct {
	// The name of the program. Defaults to os.Args[0]
	Name string `json:"-"`
	// The name of the execute binary file. Defaults to os.Args[0]
	Binary string `json:"a"`
	// Version of the program
	Version string `json:"v"`
	// A short description of the usage of this app
	Usage string `json:"-"` //`json:"u"`
	// Description of the app
	Desc string `json:"-"` //`json:"d,omitempty"`
	// List of commands to execute
	Commands []Command `json:"-"`
	// List of flags to parse
	Flags []Flag `json:"-"`
	// The action to execute when no subcommands are specified
	Action func(context *Context) `json:"-"`
	// run result
	ElseMe interface{} `json:"r"`
	// run error
	Error error `json:"e,omitempty"`
}

func NewApp() *App {
	return &App{
		Name:    os.Args[0],
		Binary:  os.Args[0],
		Version: "1.0",
		Usage:   "A new Wu cli app",
		Desc:    "",
		Action:  helpCommand.Action,
		Error:   nil,
		ElseMe:  nil,
	}
}

func (a *App) Run(arguments []string) error {
	// append help to commands
	a.Commands = append(a.Commands, helpCommand, versionCommand)
	// append version to flags
	a.Flags = append(
		a.Flags,
		helpFlag{"Show help"},
		versionFlag{"Show version"},
	)

	// parse flags
	set := flagSet(a.Name, a.Flags)
	err := set.Parse(arguments[1:])
	if err != nil {
		return err
	}

	context := NewContext(a, set, set)
	checkHelp(context)
	checkVersion(context)

	args := context.Args()
	if len(args) > 0 {
		name := args[0]
		for _, c := range a.Commands {
			if c.HasName(name) {
				c.Run(context)
				return nil
			}
		}
	}
	// Run default Action
	a.Action(context)
	return nil
}

// entry point to the cli app
func (a *App) End() {
	_ = a.Run(os.Args)
}

// Another entry point to the cli app, takes care of passing arguments and error handling
func (a *App) EndOnError() {
	if err := a.Run(os.Args); err != nil {
		os.Stderr.WriteString(fmt.Sprintln(err))
		os.Exit(1)
	}
}
