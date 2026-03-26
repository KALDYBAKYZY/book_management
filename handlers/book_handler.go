package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"book_management/models"
	"github.com/gorilla/mux"
)

var books = []models.Book{}
var nextID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	categoryStr := query.Get("category_id")
	categoryID, _ := strconv.Atoi(categoryStr)
	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 4
	}
	filtered := []models.Book{}
	for _, book := range books {
		if categoryID == 0 || book.CategoryID == categoryID {
			filtered = append(filtered, book)
		}
	}
	start := (page - 1) * limit
	end := start + limit
	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}
	json.NewEncoder(w).Encode(filtered[start:end])
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if book.Title == "" || book.Price <= 0 {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	categoryExists := false
	for _, c := range categories {
		if c.ID == book.CategoryID {
			categoryExists = true
			break
		}
	}
	if !categoryExists {
		http.Error(w, "Category not found", http.StatusBadRequest)
		return
	}
	authorExists := false
	for _, a := range authors {
		if a.ID == book.AuthorID {
			authorExists = true
			break
		}
	}
	if !authorExists {
		http.Error(w, "Author not found", http.StatusBadRequest)
		return
	}
	book.ID = nextID
	nextID++
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updated models.Book
	json.NewDecoder(r.Body).Decode(&updated)
	for i, book := range books {
		if book.ID == id {
			if updated.Title != "" {
				books[i].Title = updated.Title
			}
			if updated.Price > 0 {
				books[i].Price = updated.Price
			}
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
