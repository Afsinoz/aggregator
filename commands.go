package main

import (
	"fmt"

	"github.com/Afsinoz/aggregator/internal/config"
)

type State struct {
	cfgp *config.Config
}

type Command struct {
	name      string
	arguments []string
}

type Commands struct {
	listOfCommands map[string]func(*State, Command) error
}

// Commands methods

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.listOfCommands[name] = f

}

func (c *Commands) Run(s *State, cmd Command) error {
	f := c.listOfCommands[cmd.name]

	if err := f(s, cmd); err != nil {
		return fmt.Errorf("Running error of", err)
	}
	return nil

}

// Handlers
func handlerLogin(s *State, cmd Command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Argument is Empty!")
	}
	username := cmd.arguments[2]
	s.cfgp.CurrentUserName = username

	s.cfgp.SetUser(username)

	fmt.Printf("User %v set!", username)

	return nil

}
