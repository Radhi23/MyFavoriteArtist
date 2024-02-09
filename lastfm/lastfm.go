package lastfm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const lastFMAPIBaseURL = "http://ws.audioscrobbler.com/2.0/"
const apiKey = "a1bf6bdc94482601838cd9f964dcee83"

type LastFMTopTrackResponse struct {
	Tracks struct {
		Track []struct {
			Name   string `json:"name"`
			Artist struct {
				Name string `json:"name"`
			} `json:"artist"`
		} `json:"track"`
	} `json:"tracks"`
}

func GetTopTrackAndArtist(region string) (string, string, error) {
	// Encode region for URL
	region = url.QueryEscape(region)
	fmt.Println("ab", region)

	// Construct URL with API key and parameters
	url := fmt.Sprintf("%s?method=geo.gettoptracks&country=%s&api_key=%s&format=json", lastFMAPIBaseURL, region, apiKey)

	// Perform GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// Decode JSON response
	var data LastFMTopTrackResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", "", err
	}
	fmt.Println("top", data)
	// Extract top track and artist name
	if len(data.Tracks.Track) > 0 {
		topTrack := data.Tracks.Track[0].Name
		topArtist := data.Tracks.Track[0].Artist.Name
		fmt.Println("top", topArtist, topTrack)
		return topTrack, topArtist, nil
	}

	return "", "", fmt.Errorf("no top track found for region: %s", region)
}
