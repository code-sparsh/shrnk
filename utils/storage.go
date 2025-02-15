package utils

import (
	"fmt"
	"sync"
	"net/url"
)


type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex // Mutex for thread-safe access to the map
}


func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]string),
	}
}


func (s *URLStore) StoreURL(longURL string) (string, error) {

	// if the URL is invalid, return an error
	if _, err := url.ParseRequestURI(longURL); err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	shortURL := generateShortCode()
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortURL] = longURL

	return shortURL, nil
}

func (s *URLStore) RetrieveURL (shortURL string) (string, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()
	longURL, ok := s.urls[shortURL]
	if !ok {
		return "", fmt.Errorf("short URL not found")
	}
	return longURL, nil
}