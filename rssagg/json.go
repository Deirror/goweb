package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(wr http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s\n", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(wr, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(wr http.ResponseWriter, code int, payload any) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}
	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(code)
	wr.Write(dat)
}
