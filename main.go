package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/code-sparsh/shrnk/utils"
)

type ShortenRequest struct{
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortURL"`
	Error string `json:"error,omitempty"`
}

func shortenHandler(store *utils.URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ShortenRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error": "Invalid JSON"}`, http.StatusBadRequest)
			return
		}

		if req.URL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		shortCode, err := store.StoreURL(req.URL)
		if err != nil {
			http.Error(w, "Failed to store URL", http.StatusInternalServerError)
			return
		}

		shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ShortenResponse{ShortURL: shortURL})
	}
}

func main() {

	store := utils.NewURLStore()

	http.HandleFunc("/shorten", shortenHandler(store))

	fmt.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

	fmt.Println("Enter the URL to shorten:")
	var url string
	fmt.Scanln(&url)

	shortCode, err := store.StoreURL(url)

	if err != nil {
		fmt.Printf("Failed to store URL: %v\n", err)
		return
	}
	fmt.Printf("Shortened URL: http://short.url/%s\n", shortCode)
}
