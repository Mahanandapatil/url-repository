package storage

import (
	"sync"
)

var (
	urlMap         = make(map[string]string)
	domainCountMap = make(map[string]int)
	memoryLock     = sync.RWMutex{}
)

func StoreShortURL(shortCode, longURL string) {
	memoryLock.Lock()
	defer memoryLock.Unlock()
	urlMap[shortCode] = longURL
}

func GetLongURL(shortCode string) string {
	memoryLock.RLock()
	defer memoryLock.RUnlock()
	return urlMap[shortCode]
}

func IncrementDomainCount(domain string) {
	memoryLock.Lock()
	defer memoryLock.Unlock()
	domainCountMap[domain]++
}

func GetTopDomains() []DomainCount {
	memoryLock.RLock()
	defer memoryLock.RUnlock()

	var domainCounts []DomainCount
	for domain, count := range domainCountMap {
		domainCounts = append(domainCounts, DomainCount{Domain: domain, Count: count})
	}

	if len(domainCounts) > 3 {
		domainCounts = domainCounts[:3]
	}
	return domainCounts
}

type DomainCount struct {
	Domain string
	Count  int
}
