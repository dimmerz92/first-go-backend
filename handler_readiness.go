package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		respondWithJSON(w, http.StatusOK, struct{}{})
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}