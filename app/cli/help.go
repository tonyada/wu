package cli

import (
	"fmt"
	"os"
	"text/tabwriter"
	"text/template"
)

// The text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var AppHelpTemplate = `{{.Name}} {{.Version}} - {{.Usage}}

Usage: {{.Binary}} [options] command [command options] [arguments...]

Options:
   {{range .Flags}}{{.}}
   {{end}}
Commands:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}
`

// The text template for the command help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var CommandHelpTemplate = `{{.Name}} - {{.Usage}}

Usage: command {{.Name}} [command options] [arguments...]

Description:
   {{.Desc}}

Options:
   {{range .Flags}}{{.}}
   {{end}}
`

var helpCommand = Command{
	Name:      "help",
	ShortName: "h",
	Usage:     "Shows a list of commands or help for one command",
	Action: func(c *Context) {
		args := c.Args()
		if len(args) > 0 {
			ShowCommandHelp(c, args[0])
		} else {
			ShowAppHelp(c)
		}
	},
}

// Prints help for the App
func ShowAppHelp(c *Context) {
	printHelp(AppHelpTemplate, c.App)
}

// Prints help for the given command
func ShowCommandHelp(c *Context, command string) {
	for _, c := range c.App.Commands {
		if c.HasName(command) {
			printHelp(CommandHelpTemplate, c)
			return
		}
	}

	fmt.Printf("No help topic for '%v'\n", command)
	os.Exit(1)
}

func printHelp(templ string, data interface{}) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	t := template.Must(template.New("help").Parse(templ))
	_ = t.Execute(w, data)
	w.Flush()
}

func checkHelp(c *Context) {
	if c.GlobalBool("h") || c.GlobalBool("help") {
		ShowAppHelp(c)
		os.Exit(0)
	}
}
