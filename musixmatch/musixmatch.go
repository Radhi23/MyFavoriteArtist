package musixmatch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const musixmatchAPIBaseURL = "https://api.musixmatch.com/ws/1.1/"
const apiKey = "97193c09459725cc2204f1e5d9e74dbd"

type MusixmatchResponse struct {
	Message struct {
		Body struct {
			Lyrics struct {
				LyricsBody string `json:"lyrics_body"`
			} `json:"lyrics"`
		} `json:"body"`
	} `json:"message"`
}

func GetLyrics(track, artist string) (string, error) {
	// Encode track and artist for URL
	track = url.QueryEscape(track)
	artist = url.QueryEscape(artist)

	// Construct URL with API key, track, and artist
	url := fmt.Sprintf("%smatcher.lyrics.get?apikey=%s&q_track=%s&q_artist=%s", musixmatchAPIBaseURL, apiKey, track, artist)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data MusixmatchResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	lyrics := data.Message.Body.Lyrics.LyricsBody
	return lyrics, nil
}
