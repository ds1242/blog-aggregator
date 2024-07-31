package main

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/ds1242/blog-aggregator.git/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		Limit int32 `json:"limit"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Params{}

	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "cannot decode parameters")
		return
	}

	type PostsParams struct {
		UserID 	uuid.UUID
		Limit 	int32
		Offset	int32
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit: 	params.Limit,
		Offset: 0,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not get posts")
		return
	}
	fmt.Println(posts)
	
	RespondWithJSON(w, http.StatusOK, posts) 
}

