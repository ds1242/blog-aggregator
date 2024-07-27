package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/google/uuid"
	// "fmt"

	"github.com/ds1242/blog-aggregator.git/internal/database"
)


func (cfg *apiConfig) handlerFeedsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		Name string `json:"name"`
		URL string 	`json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Params{}
	
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	// Get context from request again
	ctx := r.Context()

	feed, err := cfg.DB.AddToFeed(ctx, database.AddToFeedParams{
		ID: 		uuid.New(),
		CreatedAt: 	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		UserID: 	user.ID,	
		Name: 		params.Name,
		Url: 		params.URL,

	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not create feed")
		return
	}

	RespondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (cfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "cound not get feeds")
		return
	}
	var feedSlice []Feed
	for _, feed := range(feeds) {
		feedSlice = append(feedSlice, databaseFeedToFeed(feed))
	}
	RespondWithJSON(w, http.StatusOK, feedSlice)
}


func (cfg *apiConfig) handlerFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		Feed string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Params{}
	
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	feedIdUUID, err := uuid.Parse(params.Feed)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode feed id")
		return 
	}

	followFeed, err := cfg.DB.FeedFollow(r.Context(), database.FeedFollowParams{
		ID: 		uuid.New(),
		CreatedAt: 	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		FeedID: 	feedIdUUID,
		UserID: 	user.ID,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not follow feed")
		return
	}

	RespondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(followFeed))
}