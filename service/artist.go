// service/artist.go
package service

import (
	"favorite-artist/image"
	"favorite-artist/lastfm"
	"favorite-artist/musixmatch"
	"fmt"
)

type ArtistInfo struct {
	TopTrack   string `json:"top_track"`
	Lyrics     string `json:"lyrics"`
	ArtistName string `json:"artist_name"`
	ImageURL   string `json:"image_url"`
}

func GetArtistInfo(region string) (ArtistInfo, error) {
	// Call Last.fm package to get top track and artist info
	topTrack, artistName, err := lastfm.GetTopTrackAndArtist(region)
	fmt.Println("topTrack:", topTrack, "artistName:", artistName)
	if err != nil {
		return ArtistInfo{}, err
	}

	// Call Musixmatch package to get lyrics
	lyrics, err := musixmatch.GetLyrics(topTrack, artistName)
	if err != nil {
		fmt.Println("Error:", err)
		return ArtistInfo{}, err
	}
	fmt.Println("Lyrics:", lyrics)
	// Call Image package to get artist image URL
	imageURL, err := image.GetImageURL(artistName)
	if err != nil {
		fmt.Println("Error:", err)
		return ArtistInfo{}, err
	}

	fmt.Println("Image URL:", imageURL)

	return ArtistInfo{
		TopTrack:   topTrack,
		Lyrics:     lyrics,
		ArtistName: artistName,
		ImageURL:   imageURL,
	}, nil
}
