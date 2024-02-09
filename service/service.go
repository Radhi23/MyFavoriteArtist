// service/service.go
package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleArtistRequest(w http.ResponseWriter, r *http.Request) {
	// Extract region from request parameters
	query := r.URL.Query()

	// Extract the value of the 'region' parameter
	region := query.Get("region")

	// Check if the 'region' parameter is empty
	if region == "" {
		// If 'region' is empty, return a bad request response
		http.Error(w, "Missing 'region' parameter", http.StatusBadRequest)
		return
	}

	// Print the extracted region
	fmt.Println("Region:", region)
	// Call the service package to get artist information
	artistInfo, err := GetArtistInfo(region)
	if err != nil {
		fmt.Println(err, "err:")
		http.Error(w, "Failed to get artist info", http.StatusInternalServerError)
		return
	}

	// Serialize artistInfo to JSON and write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artistInfo)
}
