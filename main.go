package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"
	
	"github.com/joho/godotenv"
	"github.com/ds1242/blog-aggregator.git/internal/database"
)

import _ "github.com/lib/pq"


type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("CONNECTION_STRING")
	if dbURL == "" {
		log.Fatal("no db connection environment variable set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("unable to open db")
	}

	dbQueries := database.New(db)

	cfg := &apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", healthzHandler)
	mux.HandleFunc("GET /v1/err", errorHealthHandler)

	mux.HandleFunc("POST /v1/users", cfg.handlerUsersCreate)
	mux.HandleFunc("GET /v1/users", cfg.middlewareAuth(cfg.handlerUsersGet))
	
	mux.HandleFunc("POST /v1/feeds", cfg.middlewareAuth(cfg.handlerUsersGet))

	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}