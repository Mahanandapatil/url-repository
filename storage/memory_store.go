package storage

var (
	urlMap         = make(map[string]string)
	reverseMap     = make(map[string]string)
	domainCountMap = make(map[string]int)
)

type DomainCount struct {
	Domain string
	Count  int
}

func StoreShortURL(shortCode, longURL string) {
	urlMap[shortCode] = longURL
	reverseMap[longURL] = shortCode
}

func GetLongURL(shortCode string) string {
	return urlMap[shortCode]
}

func GetShortCode(longURL string) (string, bool) {
	code, exists := reverseMap[longURL]
	return code, exists
}

func IncrementDomainCount(domain string) {
	domainCountMap[domain]++
}

func GetTopDomains() []DomainCount {
	var domainCounts []DomainCount
	for domain, count := range domainCountMap {
		domainCounts = append(domainCounts, DomainCount{Domain: domain, Count: count})
	}

	for i := 0; i < len(domainCounts)-1; i++ {
		for j := i + 1; j < len(domainCounts); j++ {
			if domainCounts[j].Count > domainCounts[i].Count {
				domainCounts[i], domainCounts[j] = domainCounts[j], domainCounts[i]
			}
		}
	}

	if len(domainCounts) > 3 {
		domainCounts = domainCounts[:3]
	}
	return domainCounts
}
