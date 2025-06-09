package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Deirror/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(wr http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Error pasrsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(req.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(wr, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(wr http.ResponseWriter, req *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(req.Context())
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}
	respondWithJSON(wr, http.StatusOK, databaseFeedsToFeeds(feeds))
}

func (apiCfg *apiConfig) handlerDeleteFeed(wr http.ResponseWriter, req *http.Request, user database.User) {
	feedIDStr := chi.URLParam(req, "feedID")
	feedID, err := uuid.Parse(feedIDStr)
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't parse feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeed(req.Context(), database.DeleteFeedParams{
		ID:     feedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't delete feed: %v", err))
		return
	}

	respondWithJSON(wr, http.StatusOK, struct{}{})
}
