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
func handlerHelp(s *State, cmd Command) error {
	cmds, err := cmdsRegister(cmd.arguments)
	if err != nil {
		return err
	}
	for _, desc := range cmds.listOfCommandsNamesDescriptions {
		s := desc
		fmt.Println(s)
	}
	return nil
}
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

// Registering a user
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

// Reset the database tables
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
	// Check the existence of the user
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
	args := cmd.arguments
	if len(args) < 1 {
		return errors.New("Not enough argument")
	}
	timeBetweenRequestsString := args[0]
	timeBetweenRequests, err := time.ParseDuration(timeBetweenRequestsString)
	if err != nil {
		return nil
	}
	fmt.Printf("Collecting feeds every %v", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func handlerAddFeed(s *State, cmd Command, usr database.User) error {
	ctx := context.Background()

	if len(cmd.arguments) < 2 {
		fmt.Printf("Not enough argument for addfeed command")
		os.Exit(1)
	}

	userId := usr.ID

	args := cmd.arguments

	feedName := args[0]
	url := args[1]

	uuid := uuid.New()

	currentTime := time.Now()

	_, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      feedName,
		Url:       url,
		UserID:    userId,
	})

	if err != nil {
		return err
	}

	feed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return err
	}
	feedId := feed.ID

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    userId,
		FeedID:    feedId,
	})
	if err != nil {
		return err
	}
	return nil

}

func handlerFeeds(s *State, cmd Command) error {
	ctx := context.Background()

	feedsList, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}
	for _, feed := range feedsList {
		fmt.Printf("Name: %v URL: %v UserName: %v\n", feed.Name, feed.Url, feed.UsersName)
	}
	//for _, feed := range feedsList {
	//	fmt.Printf("Name: %v, URL:%v, UserName:%v", feed.Name, feed.Url, feed.user_name)
	//}
	return nil

}

func handlerFeedFollows(s *State, cmd Command, usr database.User) error {
	ctx := context.Background()

	cmdArgs := cmd.arguments

	url := cmdArgs[0]

	// Get user first
	// Check the existence/logged in status of the user
	userId := usr.ID
	// Get feed
	feed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return err
	}
	feedId := feed.ID

	currentTime := time.Now()

	uuid := uuid.New()

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    userId,
		FeedID:    feedId,
	})
	if err != nil {
		return err
	}
	return nil
}

func handlerFollowing(s *State, cmd Command, usr database.User) error {
	ctx := context.Background()

	feedFollowings, err := s.db.GetFeedFollowsForUsers(ctx, usr.Name)
	if err != nil {
		return err
	}
	for _, followings := range feedFollowings {
		fmt.Printf("%v\n", followings.FeedName)
	}

	return nil
}

func handlerUnfollow(s *State, cmd Command, usr database.User) error {
	ctx := context.Background()

	url := cmd.arguments[0]

	err := s.db.DeleteFeed(ctx, url)
	if err != nil {
		return err
	}

	return nil

}
