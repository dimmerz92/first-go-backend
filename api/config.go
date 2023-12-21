package api

import "github.com/dimmerz92/go_rss_app/internal/database"

type ApiConfig struct {
	DB *database.Queries
}