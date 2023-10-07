package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(writer http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed request")
		writer.WriteHeader(500)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(data)
}

func respondWithError(writer http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 500 level error: %v", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(writer, code, errResponse{
		Error: message,
	})
}
