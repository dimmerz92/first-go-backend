package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in .env")
	}

	http.HandleFunc("/ready", handlerReadiness)
	http.HandleFunc("/err", handleErr)

	fmt.Printf("Server starting on port: %s\n", port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}