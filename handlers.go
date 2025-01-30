package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Afsinoz/aggregator/internal/database"
	"github.com/google/uuid"
)

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

func handlerRegister(s *State, cmd Command) error {

	ctx := context.Background()

	// check the user is exist or not

	userName := cmd.arguments[3]

	_, err := s.db.GetUser(ctx, userName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			fmt.Println("User does not exist, continue to create")

		} else {
			return fmt.Errorf("Database error", err)
		}
	} else {
		fmt.Printf("User with the name %s exist", userName)
		os.Exit(1)
	}

	uuid := uuid.New()

	currentTime := time.Now()

	_, err = s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      userName,
	})
	if err != nil {
		return err
	}
	return nil

}
