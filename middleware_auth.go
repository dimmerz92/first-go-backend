package main

import (
	"fmt"
	"github/dimmerz92/go_rss_app/internal/auth"
	"github/dimmerz92/go_rss_app/internal/database"
	"net/http"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(
				w,
				http.StatusForbidden,
				fmt.Sprintf("Auth error: %v", err),
			)
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Could not get user: %v", err),
			)
			return
		}

		handler(w, r, user)
	}
}