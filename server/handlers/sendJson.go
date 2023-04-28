package handlers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/sawyerKent/cli-server/server/models"
)

func SendJson(url string, data models.JsonData) ([]byte, error) {
    jsonData, err := data.MarshalJSON()
    if err != nil {
        return nil, err
    }

	resp, err := http.Post(url, "application/json", io.NopCloser(bytes.NewBuffer(jsonData)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}
