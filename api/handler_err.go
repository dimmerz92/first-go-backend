package api

import "net/http"

func HandleErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "Something went wrong")
}