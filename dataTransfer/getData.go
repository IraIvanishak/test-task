package dataTransfer

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// GetDataFromAPI fetches breed data from a web API.
func GetDataFromAPI() ApiResponse {
	response, err := http.Get("https://catfact.ninja/breeds")
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Fatalf("Error unpacking JSON: %v", err)
	}

	return apiResponse
}
