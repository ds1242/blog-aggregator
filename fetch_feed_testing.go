package main

import (
	"testing"
)

func TestFetchFeed(t *testing.T) {
	pageURL := "https://blog.boot.dev/index.xml"
	rssFeed, err := fetchRSSFeed(pageURL)
	if err != nil {
		t.Fatalf("failed to fetch: %v", err)
		return
	} 
	// Perform assertions to validate the fetched RSS feed
	if rssFeed.Channel.Title == "" {
		t.Error("Expected RSS feed channel title to be non-empty")
	}

	if len(rssFeed.Channel.Title) == 0 {
		t.Error("Expected RSS feed to have at least one item")
	}

	// Optional: Print the RSS feed for debugging; not usually recommended in actual tests
	t.Logf("Fetched RSS feed: %v", rssFeed)

}