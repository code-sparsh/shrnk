package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/code-sparsh/shrnk/utils"
)

type ShortenRequest struct{
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortURL"`
	Error string `json:"error,omitempty"`
}

func shortenHandler(DB *utils.DB) http.HandlerFunc {
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

		shortCode, err := DB.StoreURL(req.URL)
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


func redirectHandler(DB *utils.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path

		shortCode := strings.Split(shortURL, "/")[1]

		longURL, err := DB.RetrieveURL(shortCode)
		if err != nil {
			http.Error(w, "URL Not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, longURL, http.StatusMovedPermanently)
	}
}

func main() {

	DB := utils.NewDB()
	defer DB.SQL.Close()

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello World") })
	http.HandleFunc("/shorten", shortenHandler(DB))
	http.HandleFunc("/", redirectHandler(DB))


	fmt.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
