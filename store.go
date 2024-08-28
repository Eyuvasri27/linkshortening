package main

import (
	"sort"
	"strings"
	"sync"
)

type URLStore struct {
	urls          map[string]string
	domainCounter map[string]int
	mu            sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls:          make(map[string]string),
		domainCounter: make(map[string]int),
	}
}

func (s *URLStore) StoreURL(originalURL, shortURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	domain := getDomain(originalURL)
	s.urls[shortURL] = originalURL
	s.domainCounter[domain]++
}

func (s *URLStore) GetOriginalURL(shortURL string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	originalURL, exists := s.urls[shortURL]
	return originalURL, exists
}

func (s *URLStore) GetShortURL(originalURL string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for shortURL, url := range s.urls {
		if url == originalURL {
			return shortURL, true
		}
	}
	return "", false
}

func (s *URLStore) GetTopDomains(topN int) map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	type domainCount struct {
		domain string
		count  int
	}

	var counts []domainCount
	for domain, count := range s.domainCounter {
		counts = append(counts, domainCount{domain, count})
	}

	// Sort counts by count in descending order
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})

	// Get top N
	if topN > len(counts) {
		topN = len(counts)
	}

	result := make(map[string]int)
	for i := 0; i < topN; i++ {
		result[counts[i].domain] = counts[i].count
	}

	return result
}

func getDomain(url string) string {
	// Simplified domain extraction, should use a proper URL parser in real applications
	parts := strings.Split(url, "/")
	if len(parts) > 2 {
		return parts[2]
	}
	return "unknown"
}
