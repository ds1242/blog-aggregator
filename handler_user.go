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
		ResponseWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	// Get the context from the request
	ctx := r.Context()

	newUser, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID: 		uuid.New(),
		CreatedAt: 	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		Name: 		params.Name,
	})
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	

	RespondWithJSON(w, http.StatusCreated, newUser)

}