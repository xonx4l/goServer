package main

import "net/http"

func handlerErr(w http.ResponseWriter, e *http.Request) {
	respondWithError(w, 200, "something went wrong")
}
