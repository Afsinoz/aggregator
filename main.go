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

	var commands Commands

	args := os.Args

	if len(args) <= 2 {
		fmt.Println("Not Enough arguments!")
	}

	cmd := Command{
		name:      "Login",
		arguments: args,
	}

	commands.Register(cmd.name, handlerLogin)

	fmt.Println(cfg.DbURL)

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Read function error:", err)
	}
	fmt.Println(cfg.CurrentUserName)

}
