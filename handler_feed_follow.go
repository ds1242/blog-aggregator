package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/google/uuid"
	// "fmt"

	"github.com/ds1242/blog-aggregator.git/internal/database"
)

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

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request) {
	// TODO: Does this need to throw an error if no record exists?
	feedFollowId:= r.PathValue("feedFollowID")

	feedIdUUID, err := uuid.Parse(feedFollowId)
	
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "bad feed follow id")
		return 
	}

	unfollowErr := cfg.DB.UnfollowFeed(r.Context(), feedIdUUID)
	if unfollowErr != nil {
		RespondWithError(w, http.StatusBadRequest, "unable to unfollow that feed")
		return
	}

	RespondWithJSON(w, http.StatusOK, "deleted")
}