package handlers

import (
	"encoding/json"
	"net/http"

	"book_management/models"
)

var authors = []models.Author{}
var nextAuthorID = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {

	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)

	if author.Name == "" {
		http.Error(w, "Name required", http.StatusBadRequest)
		return
	}

	author.ID = nextAuthorID
	nextAuthorID++

	authors = append(authors, author)

	json.NewEncoder(w).Encode(author)
}
