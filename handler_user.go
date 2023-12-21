package main

import (
	"encoding/json"
	"fmt"
	"github/dimmerz92/go_rss_app/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfg.middlewareAuth(apiCfg.handlerGetUser).ServeHTTP(w, r)
	case http.MethodPost:
		apiCfg.handlerCreateUser(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
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

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Couldn't create user: %v", err),
		)
		return
	}
	
	respondWithJSON(w, http.StatusCreated, dbUsertoUser(user))
}
	

func (apiCfg *apiConfig) handlerGetUser(
	w http.ResponseWriter, r *http.Request, user database.User) {
		respondWithJSON(w, http.StatusOK, dbUsertoUser(user))
}