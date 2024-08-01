package main

import (
	"context"
	"database/sql"
	// "fmt"
	"log"
	"sync"
	"time"

	"github.com/ds1242/blog-aggregator.git/internal/database"
	"github.com/google/uuid"
)


func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines...", timeBetweenRequest, concurrency)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Couldn't get next feeds to fetch", err)
			continue
		}
		log.Printf("found %v feeds to fetch!", len(feeds))
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}


func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err :=  fetchRSSFeed(feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	
	
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))

	for _, item := range feedData.Channel.Item {

		var descriptionNullable sql.NullString

		if item.Description != "" {
			descriptionNullable.String = item.Description
			descriptionNullable.Valid = true
		} else {
			descriptionNullable.Valid = false
		}

		// List of potential layouts
		layouts := []string{
			time.RFC3339,
			time.UnixDate,
			// Add custom layouts as needed
		}

		var timeConverted time.Time
		var err error

		// Attempt parsing with each layout
		for _, layout := range layouts {
			timeConverted, err = time.Parse(layout, item.PubDate)
			if err == nil {
				break
			}
		}
		var publishedAtNullable sql.NullTime

		if !timeConverted.IsZero() {
			publishedAtNullable.Time = timeConverted.UTC()
			publishedAtNullable.Valid = true
		} else {
			publishedAtNullable.Valid = false
		}

		db.AddPost(context.Background(), database.AddPostParams{
			ID: 			uuid.New(),
			CreatedAt: 		time.Now().UTC(),
			UpdatedAt: 		time.Now().UTC(),
			Title: 			item.Title,
			Url: 			item.Link,
			Description: 	descriptionNullable,	
			PublishedAt: 	publishedAtNullable,
			FeedID: 		feed.ID,	
		})
	}

}
