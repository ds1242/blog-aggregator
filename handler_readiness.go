package main

import (
	"net/http"

	"github.com/ds1242/blog-aggregator.git/helpers"
)


func healthzHandler(w http.ResponseWriter, r *http.Request) {
	type statusOKStruct struct {
		Status string `json:"status"`
	}
	helpers.RespondWithJSON(w, http.StatusOK, statusOKStruct{
		Status: "ok",
	})	
}