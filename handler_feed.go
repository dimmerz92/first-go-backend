package main

import (
	"encoding/json"
	"fmt"
	"github/dimmerz92/go_rss_app/internal/database"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleFeed(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		apiCfg.middlewareAuth(apiCfg.handlerCreateFeed).ServeHTTP(w, r)
	case http.MethodGet:
		if r.URL.Path == "/feeds" {
			apiCfg.middlewareAuth(apiCfg.handlerGetAllFeeds).ServeHTTP(w, r)
		} else {
			apiCfg.middlewareAuth(apiCfg.handlerGetFeed).ServeHTTP(w, r)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (apiCfg *apiConfig) handlerCreateFeed(
	w http.ResponseWriter, r *http.Request, user database.User) {
		type parameters struct {
			Name string `json:"name"`
			URL string `json:"url"`
		}
		
		decoder := json.NewDecoder(r.Body)

		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			respondWithError(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Error parsing JSON: %v", err),
			)
			return
		}

		feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name: params.Name,
			Url: params.URL,
			UserID: user.ID,
		})
		if err != nil {
			respondWithError(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Could not create feed: %v", err),
			)
			return
		}

		respondWithJSON(w, http.StatusCreated, dbFeedtoFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeed(
	w http.ResponseWriter, r *http.Request, user database.User) {
		id, err := uuid.Parse(strings.TrimPrefix(r.URL.Path, "/feeds/"))
		if err != nil {
			respondWithError(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Not a valid id: %v", err),
			)
			return
		}

		feed, err := apiCfg.DB.GetFeed(r.Context(), database.GetFeedParams{
			ID: id,
			UserID: user.ID,
		})
		if err != nil {
			respondWithError(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Could not get feed: %v", err),
			)
			return
		}

		respondWithJSON(w, http.StatusOK, dbFeedtoFeed(feed))
}

func (apiCfg *apiConfig) handlerGetAllFeeds(
	w http.ResponseWriter, r *http.Request, user database.User) {
		feeds, err := apiCfg.DB.GetAllFeeds(r.Context(), user.ID)
		if err != nil {
			respondWithError(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Could not get feeds: %v", err),
			)
			return
		}

		respondWithJSON(w, http.StatusOK, dbFeedstoFeeds(feeds))
}