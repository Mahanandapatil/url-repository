package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/services"
	"url-shortener/storage"
)

func HandleShortenURL(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		LongURL string `json:"long_url"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	shortCode := services.ShortenURL(requestData.LongURL)

	response := map[string]string{
		"short_url": fmt.Sprintf("/%s", shortCode),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:] // remove the leading "/"
	longURL := storage.GetLongURL(shortCode)

	if longURL == "" {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func HandleMetrics(w http.ResponseWriter, r *http.Request) {
	topDomains := services.GetTopDomains()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topDomains)
}
