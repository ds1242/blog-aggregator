package main

import (
	"encoding/json"
	"net/http"

	"github.com/ds1242/blog-aggregator.git/internal/database"
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
	
	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit: 	params.Limit,
		Offset: 0,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not get posts")
		return
	}
	var postSlice []Post
	for _, post := range posts {
		postSlice = append(postSlice, databasePostToPost(post))
	}
	RespondWithJSON(w, http.StatusOK, postSlice) 
}

