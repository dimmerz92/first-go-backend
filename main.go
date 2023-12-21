package main

import (
	"database/sql"
	"fmt"
	"github/dimmerz92/go_rss_app/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in .env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("db_URL is not found in .env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Could not establish connection to database")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	http.HandleFunc("/ready", handlerReadiness)
	http.HandleFunc("/err", handleErr)
	http.HandleFunc("/users", apiCfg.handleUser)
	http.HandleFunc("/feeds", apiCfg.handleFeed)
	http.HandleFunc("/feeds/", apiCfg.handleFeed)

	fmt.Printf("Server starting on port: %s\n", port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}