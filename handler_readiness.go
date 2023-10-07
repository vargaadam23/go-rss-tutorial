package main

import "net/http"

func handleReadiness(writer http.ResponseWriter, r *http.Request) {
	respondWithJson(writer, 200, struct{}{})
}
