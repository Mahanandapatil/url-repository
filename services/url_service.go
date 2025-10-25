package services

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"url-shortener/storage"
)

func ShortenURL(longURL string) string {
	// Check if this URL is already shortened
	if code, exists := storage.GetShortCode(longURL); exists {
		return code
	}

	// Generate new short code
	shortCode := generateShortCode(longURL)
	storage.StoreShortURL(shortCode, longURL)

	// Increment domain count for metrics
	domain := extractDomain(longURL)
	storage.IncrementDomainCount(domain)

	return shortCode
}

func generateShortCode(longURL string) string {
	hash := md5.Sum([]byte(longURL))
	return hex.EncodeToString(hash[:])[:6]
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
