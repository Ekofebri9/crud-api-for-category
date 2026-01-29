package handlers

import (
	"crud-api-category/internal/models"
	"encoding/json"
	"net/http"
)

func parseBody(w http.ResponseWriter, r *http.Request, v *models.Category) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
}
