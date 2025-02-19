package main

import (
	"context"
	"fmt"

	"github.com/Afsinoz/aggregator/internal/rss"
)

func scrapeFeeds(s *State) error {
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := rss.FetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}
	for _, feedItem := range rssFeed.Channel.Item {
		fmt.Println("f %v", feedItem.Title)
	}
	return nil

}
