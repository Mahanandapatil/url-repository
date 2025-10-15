package main

import (
	"fmt"
	"net/http"
	"url-shortener/handlers"
)

func main() {
	fmt.Println("URL Shortener running on http://localhost:8080")

	http.HandleFunc("/shorten", handlers.HandleShortenURL)
	http.HandleFunc("/metrics/top-domains", handlers.HandleMetrics)
	http.HandleFunc("/", handlers.HandleRedirect)

	http.ListenAndServe(":8080", nil)
}
