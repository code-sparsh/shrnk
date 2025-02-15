package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"sync"
	"net/url"
)

type urlStore struct {
	urls map[string]string
	mu   sync.RWMutex // Mutex for thread-safe access to the map
}

const shortCodeLength = 7

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	sb := strings.Builder{}
	sb.Grow(shortCodeLength)
	for i := 0; i < shortCodeLength; i++ {
		sb.WriteByte(charSet[seededRand.Intn(len(charSet))])
	}
	return sb.String()
}

func main () {
	
	fmt.Println("Enter the URL to shorten:")
	var url string 
	fmt.Scanln(&url)

	store := urlStore{
		urls: make(map[string]string),
	}

	shortCode, err := store.storeURL(url)

	if err != nil {
		fmt.Printf("Failed to store URL: %v\n", err)
		return
	}
	fmt.Printf("Shortened URL: http://short.url/%s\n", shortCode)

}

func (s *urlStore) storeURL(longURL string) (string, error) {

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
