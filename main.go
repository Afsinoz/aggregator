package main

import (
	"fmt"
	"os"

	"github.com/Afsinoz/aggregator/internal/config"
)

func main() {
	fmt.Println("vim-go")
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Read function error:", err)
	}

	var state State

	state.cfgp = &cfg

	commandList := make(map[string]func(*State, Command) error)

	cmds := Commands{
		listOfCommands: commandList,
	}

	args := os.Args
	fmt.Println(args[1:])
	if len(args) <= 2 {
		fmt.Println("Not Enough arguments!")
	}

	cmd := Command{
		name:      "Login",
		arguments: args,
	}

	cmds.Register(cmd.name, handlerLogin)

	cmds.Run(&state, cmd)

	fmt.Println(cfg.DbURL)

	fmt.Println()

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Read function error:", err)
	}
	fmt.Println(cfg.CurrentUserName)

}
