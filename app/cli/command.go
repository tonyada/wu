package cli

type Command struct {
	// The name of the command
	Name string
	// short name of the command. Typically one character
	ShortName string
	// A short description of the usage of this command
	Usage string
	// A longer explanation of how the command works
	Desc string
	// The function to call when this command is invoked
	Action func(context *Context)
	// List of flags to parse
	Flags []Flag
}

func (c Command) Run(ctx *Context) {
	set := flagSet(c.Name, c.Flags)
	err := set.Parse(ctx.Args()[1:])
	if err != nil {
		println(err)
	}
	c.Action(NewContext(ctx.App, set, ctx.globalSet))
}

func (c Command) HasName(name string) bool {
	return c.Name == name || c.ShortName == name
}
