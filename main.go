package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dimmerz92/go_rss_app/api"
	"github.com/dimmerz92/go_rss_app/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

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

	apiCfg := api.ApiConfig{
		DB: database.New(conn),
	}

	http.HandleFunc("/ready", api.HandlerReadiness)
	http.HandleFunc("/err", api.HandleErr)
	http.HandleFunc("/users", apiCfg.HandleUser)
	http.HandleFunc("/feeds", apiCfg.HandleFeed)
	http.HandleFunc("/feeds/", apiCfg.HandleFeed)

	fmt.Printf("Server starting on port: %s\n", port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}