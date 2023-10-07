package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/vargaadam23/rss-project-go/internal/auth"
	"github.com/vargaadam23/rss-project-go/internal/database"
)

func (apiConfig *apiConfig) handleCreateUser(writer http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	err = apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ApiKey:    auth.GenerateApiKey(),
	})

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("User error: %v", err))
		return
	}

	respondWithJson(writer, 201, struct{}{})
}

type UserResponse struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func (apiConfig *apiConfig) handleGetUserByApiKey(writer http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(writer, 200, UserResponse{
		Id:     user.ID,
		Name:   user.Name,
		ApiKey: user.ApiKey,
	})
}

func (apiConfig *apiConfig) handleGetUserPosts(writer http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiConfig.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  5,
	})

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("User posts error: %v", err))
		return
	}

	respondWithJson(writer, 200, posts)
}
