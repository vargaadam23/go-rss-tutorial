package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/vargaadam23/rss-project-go/internal/database"
)

func (apiConfig *apiConfig) handleCreateFeedFollow(writer http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId int32 `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	err = apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Feed error: %v", err))
		return
	}

	respondWithJson(writer, 201, FeedFollowResponse{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    params.FeedId,
		UserId:    user.ID,
	})
}

func feedTofeedFollowResponse(feed database.FeedFollow) FeedFollowResponse {
	return FeedFollowResponse{
		Id:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		FeedID:    feed.FeedID,
		UserId:    feed.UserID,
	}
}

type FeedFollowResponse struct {
	Id        int32     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    int32     `json:"feed_id"`
	UserId    int32     `json:"user_id"`
}

func (apiConfig *apiConfig) handleGetFeedFollows(writer http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiConfig.DB.GetFeedFollowByUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Feed error: %v", err))
		return
	}

	respondWithJson(writer, 200, feedFollows)
}

func (apiConfig *apiConfig) handleDeleteFeedFollow(writer http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId := chi.URLParam(r, "feedFollowId")
	feedFollowIdInt, err := strconv.ParseInt(feedFollowId, 10, 32)

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Invalid feed follow id given: %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     int32(feedFollowIdInt),
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(writer, 400, fmt.Sprintf("Feed follow delete error error: %v", err))
		return
	}

	respondWithJson(writer, 200, struct{}{})
}
