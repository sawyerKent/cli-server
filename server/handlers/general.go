package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sawyerKent/cli-server/server/models"
)

func GetEndpoint(url string) (models.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if strings.HasSuffix(url, "/HappyLang") {
		var result models.HappyLangResponse
		json.Unmarshal(body, &result)
		return result, nil
	} else {
		var result models.TextResponse
		json.Unmarshal(body, &result)
		return result, nil
	}
}

func PostEndpoint(urlStr string, data models.HappyLangResponse) (models.Response, error) {
	formData := url.Values{}
	formData.Add("FRVRID", data.FRVRID)
	formData.Add("language", data.Language)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if strings.HasSuffix(urlStr, "/HappyLang") {
		var result models.HappyLangResponse
		json.NewDecoder(resp.Body).Decode(&result)
		return result, nil
	} else {
		var result models.TextResponse
		json.NewDecoder(resp.Body).Decode(&result)
		return result, nil
	}
}

func PostJsonEndpoint(urlStr string, data models.HappyLangResponse) (models.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if strings.HasSuffix(urlStr, "/HappyLang") {
		var result models.HappyLangResponse
		json.NewDecoder(resp.Body).Decode(&result)
		return result, nil
	} else {
		var result models.TextResponse
		json.NewDecoder(resp.Body).Decode(&result)
		return result, nil
	}
}
