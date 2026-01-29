package main

import (
	"crud-api-category/configs"
	"crud-api-category/internal/databases"
	"crud-api-category/internal/handlers"
	"crud-api-category/internal/repositories"
	"crud-api-category/internal/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// initialized config
	config := configs.Init()

	// initialized database
	db, err := databases.Init(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, There! This is Category API")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)

	addr := fmt.Sprintf(":%s", config.Port)
	fmt.Println("Starting server on :", config.Port)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
