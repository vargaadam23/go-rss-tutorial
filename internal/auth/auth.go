package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"strings"
)

func GenerateApiKey() string {
	// Generate a random input
	randomInput := make([]byte, 32) // 32 bytes for SHA-256
	_, err := rand.Read(randomInput)
	if err != nil {
		log.Fatal(err)
	}

	// Compute the SHA-256 hash
	sha256Hash := sha256.Sum256(randomInput)

	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(sha256Hash[:])
}

// Authorization: ApiKey {api_key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentication provided")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part in auth header")
	}

	return vals[1], nil
}
