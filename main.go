package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

import _ "github.com/lib/pq"

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", healthzHandler)
	mux.HandleFunc("GET /v1/err", errorHealthHandler)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}