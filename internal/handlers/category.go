package handlers

import (
	"crud-api-category/internal/models"
	"crud-api-category/internal/services"
	"encoding/json"
	"net/http"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategories - GET /api/categories
func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	parseBody(w, r, &newCategory)
	err := h.service.Create(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newCategory)
	if err != nil {
		http.Error(w, "Failed to encode category", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
