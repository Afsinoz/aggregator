package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Afsinoz/aggregator/internal/database"
	"github.com/Afsinoz/aggregator/internal/rss"
	"github.com/google/uuid"
)

// Handlers
func handlerLogin(s *State, cmd Command) error {

	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Argument is Empty!")
	}
	userName := cmd.arguments[0]

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, userName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errors.New("User doesn't exist")
			os.Exit(1)
		} else {
			return fmt.Errorf("another GetUser issue %v", err)
		}
	}

	s.cfgp.CurrentUserName = userName

	s.cfgp.SetUser(userName)

	fmt.Printf("User %v set!", userName)

	return nil

}

func handlerRegister(s *State, cmd Command) error {

	ctx := context.Background()

	// check the user is exist or not

	userName := cmd.arguments[0]

	_, err := s.db.GetUser(ctx, userName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			fmt.Println("User does not exist, continue to create")

		} else {
			return fmt.Errorf("database error: %v", err)
		}
	} else {
		fmt.Printf("User with the name %s exist, ", userName)
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

	fmt.Printf("User %s is created succesfully!", userName)

	s.cfgp.CurrentUserName = userName

	s.cfgp.SetUser(userName)

	fmt.Printf("Creation time: %v, id: %v ", currentTime, uuid)
	return nil

}

func handlerReset(s *State, cmd Command) error {
	ctx := context.Background()

	err := s.db.DeleteUsers(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Users list succesfully deleted!")
	return nil

}

func handlerUsers(s *State, cmd Command) error {
	ctx := context.Background()

	usrList, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, usrName := range usrList {
		if usrName == s.cfgp.CurrentUserName {
			fmt.Printf("* %v (current)\n", usrName)
		} else {
			fmt.Printf("* %v\n", usrName)
		}
	}
	return nil
}

func handlerAgg(s *State, cmd Command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	rssFeed, err := rss.FetchFeed(ctx, feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("%v/n", rssFeed)
	return nil
}
