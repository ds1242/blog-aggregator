package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	const filepathRoot = "."

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", healthzHandler)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on PORT: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}