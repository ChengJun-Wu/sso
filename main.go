package main

import (
	"os"
	"sso/commands"
	_ "sso/statics"
)

var commandMap = map[string]commands.Command {
	"help": &commands.Help{},
	"run": &commands.Run{},
	"route": &commands.Route{},
}

func main()  {
	command := "help"
	for idx, val := range os.Args{
		if idx == 1 {
			command = val
		}
	}
	commandMap[command].Run()
}
