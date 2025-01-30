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
	fmt.Println("vim-go")
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

	args := os.Args[1:]
	fmt.Println(args[1:])
	if len(args) <= 2 {
		fmt.Println("Not Enough arguments!")
	}

	cmds, err := cmdsRegister(args)

	cmds.Run(&state, cmd[args[2]])

	fmt.Println(cfg.DbURL)

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Read function error:", err)
	}
	fmt.Println(cfg.CurrentUserName)

}
