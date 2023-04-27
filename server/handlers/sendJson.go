package handlers

import (
	"os"
	"io"
	"net/http"
	"bytes"
)

func SendJson(url, filepath string) ([]byte, error) {
	// Read the JSON file
	jsonData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Send POST request with JSON data
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
