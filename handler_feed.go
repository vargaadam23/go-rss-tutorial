package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/vargaadam23/rss-project-go/internal/database"
)

func (apiConfig *apiConfig) handleCreateFeed(writer http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	err = apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Feed error: %v", err))
		return
	}

	respondWithJson(writer, 201, FeedResponse{
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       params.Url,
		UserId:    user.ID,
	})
}

func feedTofeedResponse(feed database.Feed) FeedResponse {
	return FeedResponse{
		Id:        feed.ID,
		Name:      feed.Name,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Url:       feed.Url,
		UserId:    feed.UserID,
	}
}

type FeedResponse struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	UserId    int32     `json:"user_id"`
}

func (apiConfig *apiConfig) handleGetFeeds(writer http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := apiConfig.DB.GetFeedsForUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Feed error: %v", err))
		return
	}

	respondWithJson(writer, 200, feeds)
}
