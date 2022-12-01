package cli

type DotCommand struct {
	Name string
	Func func([]string)
}

var DotCommands = []DotCommand{
	{"configure", ConfigureCmd},
	{"help", HelpCmd},
}
