package api

import "net/http"

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		respondWithJSON(w, http.StatusOK, struct{}{})
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}