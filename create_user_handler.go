package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ds1242/blog-aggregator.git/internal/database"
	"github.com/google/uuid"
)


func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Params{}
	
	err := decoder.Decode(&params)
	if err != nil {
		ResponseWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	// Generate a new UUID for the user's ID
    id := uuid.New()

    // Get the current time for `created_at` and `updated_at`
    now := time.Now()
	// Get the context from the request
	ctx := r.Context()

	newUser, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID: 		id,
		CreatedAt: 	now,
		UpdatedAt: 	now,
		Name: 		params.Name,
	})
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	

	RespondWithJSON(w, http.StatusCreated, newUser)

}