package image

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	googleAPIKey       = "AIzaSyCKHpJWkwATRkHIwSeMNboTbVYTbQeklS4"
	googleSearchEngine = "73938fd31506048b8"
)

type GoogleImageSearchResponse struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}

func GetImageURL(query string) (string, error) {
	baseURL := "https://www.googleapis.com/customsearch/v1"
	queryParam := url.QueryEscape(query)
	searchURL := fmt.Sprintf("%s?key=%s&cx=%s&q=%s&searchType=image", baseURL, googleAPIKey, googleSearchEngine, queryParam)

	resp, err := http.Get(searchURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response GoogleImageSearchResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if len(response.Items) == 0 {
		return "", fmt.Errorf("no images found for query: %s", query)
	}

	return response.Items[0].Link, nil
}
