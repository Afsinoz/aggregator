package main

import (
	"context"

	"github.com/Afsinoz/aggregator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, usr database.User) error) func(s *State, cmd Command) error {

	return func(s *State, cmd Command) error {

		ctx := context.Background()

		currentUserName := s.cfgp.CurrentUserName

		currentUser, err := s.db.GetUser(ctx, currentUserName)
		if err != nil {
			return err
		}

		handler(s, cmd, currentUser)
		return nil
	}
}
