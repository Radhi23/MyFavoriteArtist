package main

import (
	"net/http"

	"github.com/yourusername/yourproject/service"
)

func main() {
	http.HandleFunc("/artist", service.HandleArtistRequest)
	http.ListenAndServe(":8080", nil)
}
