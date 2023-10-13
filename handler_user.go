package main

import "net/http"

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.request) {
	respondWithJSON(w, 200, struct{}{})
}
