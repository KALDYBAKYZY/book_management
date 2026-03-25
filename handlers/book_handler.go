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
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	if book.Title == "" || book.Price <= 0 {
		http.Error(w, "Invalid data", http.StatusBadRequest)
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
