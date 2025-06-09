package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Deirror/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(wr http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Error pasrsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(wr, http.StatusCreated, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(wr http.ResponseWriter, req *http.Request, user database.User) {
	respondWithJSON(wr, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(wr http.ResponseWriter, req *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(req.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(wr, http.StatusBadRequest, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJSON(wr, http.StatusOK, databasePostsToPosts(posts))
}
