package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Song struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /songs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		songs := []Song{
			{ID: 1, Title: "Eyes closed", Author: "Ed Sheeran"},
			{ID: 2, Title: "Perfect", Author: "Ed Sheeran"},
		}

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(songs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	log.Println("Server running on http://0.0.0.0:8080...")

	err := http.ListenAndServe("0.0.0.0:8080", enableCORS(mux))
	if err != nil {
		log.Panicf("failed to start server: %v", err)
	}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
