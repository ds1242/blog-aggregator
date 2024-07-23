package main

import (
	"net/http"

	"github.com/ds1242/blog-aggregator.git/helpers"
)

func errorHealthHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ResponseWithError(w, http.StatusInternalServerError, "Internal Server Error")
}