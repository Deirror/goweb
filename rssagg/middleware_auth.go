package main

import (
	"fmt"
	"net/http"

	"github.com/Deirror/rssagg/internal/auth"
	"github.com/Deirror/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		apiKey, err := auth.GetAPIKey(req.Header)
		if err != nil {
			respondWithError(wr, http.StatusUnauthorized, fmt.Sprintf("Error getting API key: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(req.Context(), apiKey)
		if err != nil {
			respondWithError(wr, http.StatusNotFound, fmt.Sprintf("User not found: %v", err))
			return
		}

		handler(wr, req, user)
	}
}
