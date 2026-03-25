package handlers

import (
	"encoding/json"
	"net/http"

	"book_management/models"
)

var categories = []models.Category{}
var nextCategoryID = 1

func GetCategories(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)
	if category.Name == "" {
		http.Error(w, "Name required", http.StatusBadRequest)
		return
	}
	category.ID = nextCategoryID
	nextCategoryID++
	categories = append(categories, category)
	json.NewEncoder(w).Encode(category)
}
