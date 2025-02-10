package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Afsinoz/aggregator/internal/config"
	"github.com/Afsinoz/aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Read function error:", err)
	}

	var state State

	state.cfgp = &cfg

	// Open the database
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Errorf("connecting db problem", err)
	}
	// Data base queries from the generated database
	dbQueries := database.New(db)

	state.db = dbQueries

	args := os.Args

	var cmdArgs []string

	cmdName := args[1]
	if len(args[2:]) > 0 {
		cmdArgs = args[2:]
	}
	cmds, err := cmdsRegister(cmdArgs)

	cmd := Command{
		name:      cmdName,
		arguments: cmdArgs,
	}

	if err := cmds.Run(&state, cmd); err != nil {
		fmt.Println(err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Read function error:", err)
	}

}
