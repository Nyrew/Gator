package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nyrew/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// Check if the feed already exists
	existingFeed, err := s.db.GetFeedByURL(context.Background(), url)
	if err == nil {
		fmt.Printf("Feed already exists: %+v\n", existingFeed)

		// Automatically follow the existing feed
		_, followErr := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
			UserID: user.ID,
			FeedID: existingFeed.ID,
		})
		if followErr != nil {
			return fmt.Errorf("error automatically following existing feed: %w", followErr)
		}

		fmt.Printf("Successfully followed existing feed: %s\n", existingFeed.Name)
		return nil
	} else if err == sql.ErrNoRows {
		// Feed doesn't exist, create it
	} else {
		return fmt.Errorf("error checking for existing feed: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("Feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=====================================")
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		printFeed(feed, user)
		fmt.Println("=====================================")
	}

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
	fmt.Printf("* LastFetchedAt: %v\n", feed.LastFetchedAt.Time)
}
