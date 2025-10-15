package services

import (
	"fmt"
	"strings"
	"url-shortener/storage"
)

func ShortenURL(longURL string) string {
	shortCode := generateShortCode(longURL)
	storage.StoreShortURL(shortCode, longURL)
	domain := extractDomain(longURL)
	storage.IncrementDomainCount(domain)
	return shortCode
}

func generateShortCode(longURL string) string {
	hash := fmt.Sprintf("%x", longURL)

	return hash[:6]
}

func extractDomain(longURL string) string {
	parts := strings.Split(longURL, "://")
	if len(parts) < 2 {
		return ""
	}
	domain := strings.Split(parts[1], "/")[0]
	return domain
}

func GetTopDomains() []storage.DomainCount {
	return storage.GetTopDomains()
}
