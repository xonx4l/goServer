package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.printf("Failed to marshal JSON response: %V", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
