package main

import (
	"crud-api-category/configs"
	"crud-api-category/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var categories = []models.Category{
	{ID: 1, Name: "Makanan Ringan", Description: "Berbagai jenis makanan ringan"},
	{ID: 2, Name: "Minuman", Description: "Berbagai jenis minuman segar"},
}

func parseBody(w http.ResponseWriter, r *http.Request, v *models.Category) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
}

func main() {
	// initialized config
	config := configs.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, There! This is Category API")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			for _, category := range categories {
				if category.ID == id {
					w.Header().Set("Content-Type", "application/json")
					err := json.NewEncoder(w).Encode(category)
					if err != nil {
						http.Error(w, "Failed to encode category", http.StatusInternalServerError)
						return
					}
					w.WriteHeader(http.StatusOK)
					return
				}
			}
			http.Error(w, "Category not found", http.StatusNotFound)
		case http.MethodPut:
			var updatedCategory models.Category
			parseBody(w, r, &updatedCategory)
			for i, category := range categories {
				if category.ID == id {
					updatedCategory.ID = id
					categories[i] = updatedCategory
					w.Header().Set("Content-Type", "application/json")
					_ = json.NewEncoder(w).Encode(updatedCategory)
					w.WriteHeader(http.StatusOK)
					return
				}
			}
			http.Error(w, "Category not found", http.StatusNotFound)
		case http.MethodDelete:
			for i, category := range categories {
				if category.ID == id {
					categories = append(categories[:i], categories[i+1:]...)
					w.Header().Set("Content-Type", "application/json")
					_ = json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted"})
					w.WriteHeader(http.StatusOK)
					return
				}
			}
			http.Error(w, "Category not found", http.StatusNotFound)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(categories)
			if err != nil {
				http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		case http.MethodPost:
			var newCategory models.Category
			parseBody(w, r, &newCategory)
			newCategory.ID = len(categories) + 1
			categories = append(categories, newCategory)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(newCategory)
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	addr := fmt.Sprintf(":%s", config.Port)
	fmt.Println("Starting server on :", config.Port)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
