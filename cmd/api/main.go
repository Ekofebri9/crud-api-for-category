package main

import (
	"crud-api-category/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var categories = []models.Category{
	{ID: 1, Name: "Makanan Ringan", Description: "Berbagai jenis makanan ringan"},
	{ID: 2, Name: "Minuman", Description: "Berbagai jenis minuman segar"},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, There! This is Category API")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(categories)
			if err != nil {
				http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
