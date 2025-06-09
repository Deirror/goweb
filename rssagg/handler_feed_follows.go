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

func (apiCfg *apiConfig) handlerCreateFeedFollow(wr http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Error pasrsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(req.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(wr, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(wr http.ResponseWriter, req *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(req.Context(), user.ID)
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}
	respondWithJSON(wr, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(wr http.ResponseWriter, req *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(req, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't parse feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(req.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	respondWithJSON(wr, http.StatusOK, struct{}{})
}
