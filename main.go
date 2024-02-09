package main

import (
	"net/http"

	"favorite-artist/service"
)

func main() {
	http.HandleFunc("/artist", service.HandleArtistRequest)
	http.ListenAndServe(":8080", nil)
}
