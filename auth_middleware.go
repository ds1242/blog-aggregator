package main

import (
	"net/http"

	"github.com/ds1242/blog-aggregator.git/auth"
	"github.com/ds1242/blog-aggregator.git/internal/database"

)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}