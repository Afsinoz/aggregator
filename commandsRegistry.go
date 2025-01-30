package main

// Command registration

func cmdsRegister(args []string) (Commands, error) {
	// here is the kind a the manual entry of the possible command
	commandList := make(map[string]func(*State, Command) error)

	var cmds Commands

	cmds.listOfCommands = commandList

	// register login

	cmdLogin := Command{
		name:      "Login",
		arguments: args,
	}

	cmds.Register(cmdLogin.name, handlerLogin)

	// register Register :D,
	// TODO: change the name

	cmdRegister := Command{
		name:      "Register",
		arguments: args,
	}
	cmds.Register(cmdRegister.name, handlerRegister)
	return cmds, nil

}
