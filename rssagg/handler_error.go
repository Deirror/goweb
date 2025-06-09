package main

import "net/http"

func handlerError(wr http.ResponseWriter, req *http.Request) {
	respondWithError(wr, http.StatusBadRequest, "Something went wrong")
}
