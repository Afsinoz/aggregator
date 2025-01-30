package main

import (
	"fmt"

	"github.com/Afsinoz/aggregator/internal/config"
	"github.com/Afsinoz/aggregator/internal/database"
)

type State struct {
	db   *database.Queries
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