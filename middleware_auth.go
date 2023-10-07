package main

import (
	"fmt"
	"net/http"

	"github.com/vargaadam23/rss-project-go/internal/auth"
	"github.com/vargaadam23/rss-project-go/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		}

		result, err := apiConfig.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("User error: %v", err))
			return
		}

		handler(w, r, result)
	}
}
