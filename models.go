package main

import (
	"time"

	"github.com/ds1242/blog-aggregator.git/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.Apikey,
	}
}


type Feed struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed {
		ID: 		feed.ID,
		CreatedAt: 	feed.CreatedAt,
		UpdatedAt: 	feed.UpdatedAt,
		UserID:		feed.UserID,
		Name:		feed.Name,
		Url: 		feed.Url,
	}
}