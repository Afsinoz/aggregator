package main

import (
	"errors"
	"fmt"

	"github.com/Afsinoz/aggregator/internal/config"
	"github.com/Afsinoz/aggregator/internal/database"
)

type State struct {
	db   *database.Queries
	cfgp *config.Config
}

type Command struct {
	name        string
	arguments   []string
	description string
}

type Commands struct {
	listOfCommandsNamesDescriptions map[string]string
	listOfCommands                  map[string]func(*State, Command) error
}

// Commands methods

func (c *Commands) Register(cmd Command, f func(*State, Command) error) {
	c.listOfCommands[cmd.name] = f
	s := fmt.Sprintf("%v: %v", cmd.name, cmd.description)
	c.listOfCommandsNamesDescriptions[cmd.name] = s
}

func (c *Commands) Run(s *State, cmd Command) error {
	f, ok := c.listOfCommands[cmd.name]

	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)

}

// Command registration

func cmdsRegister(args []string) (Commands, error) {
	// here is the kind a the manual entry of the possible command
	commandList := make(map[string]func(*State, Command) error)

	descList := make(map[string]string)

	var cmds Commands

	cmds.listOfCommands = commandList

	cmds.listOfCommandsNamesDescriptions = descList

	cmdHelp := Command{
		name:        "help",
		arguments:   args,
		description: "List of commmands",
	}

	cmds.Register(cmdHelp, handlerHelp)
	// register login

	cmdLogin := Command{
		name:        "login",
		arguments:   args,
		description: "logging the user",
	}

	cmds.Register(cmdLogin, handlerLogin)

	// register Register :D,
	// TODO: change the name

	cmdRegister := Command{
		name:        "register",
		arguments:   args,
		description: "Registering user into the database",
	}
	cmds.Register(cmdRegister, handlerRegister)

	cmdReset := Command{
		name:        "reset",
		arguments:   args,
		description: "Resetting users list.",
	}
	cmds.Register(cmdReset, handlerReset)

	cmdAgg := Command{
		name:        "agg",
		arguments:   args,
		description: "Aggregating the URL",
	}
	cmds.Register(cmdAgg, handlerAgg)

	cmdUsers := Command{
		name:        "users",
		arguments:   args,
		description: "Printing users list",
	}
	cmds.Register(cmdUsers, handlerUsers)

	cmdAddFeed := Command{
		name:        "addfeed",
		arguments:   args,
		description: "Adding the feed to user.",
	}
	cmds.Register(cmdAddFeed, MiddlewareLoggedIn(handlerAddFeed))

	cmdFeeds := Command{
		name:        "feeds",
		arguments:   args,
		description: "Prints all the feeds of a user.",
	}
	cmds.Register(cmdFeeds, handlerFeeds)

	cmdFeedFollows := Command{
		name:        "follow",
		arguments:   args,
		description: "Creates a new feed from the user.",
	}
	cmds.Register(cmdFeedFollows, MiddlewareLoggedIn(handlerFeedFollows))

	cmdFollowing := Command{
		name:        "following",
		arguments:   args,
		description: "Return all the feed follow by the user",
	}
	cmds.Register(cmdFollowing, MiddlewareLoggedIn(handlerFollowing))

	cmdUnfollow := Command{
		name:        "unfollow",
		arguments:   args,
		description: "Unfollow a feed by url from a user.",
	}

	cmds.Register(cmdUnfollow, MiddlewareLoggedIn(handlerUnfollow))

	return cmds, nil
}
