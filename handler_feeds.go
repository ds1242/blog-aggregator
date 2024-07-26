package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/google/uuid"

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