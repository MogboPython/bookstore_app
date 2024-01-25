package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func insertBookHandler(w http.ResponseWriter, request *http.Request) {
	var book Book
	json.NewDecoder(request.Body).Decode(&book)

	// Call the CreateBook function to add a book to the database
	err = CreateBook(db, book.Title, book.Author, book.Publisher, book.Genre, book.Year_published)
	if err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// Call the GetBooks function to fetch the books data from the database
	books, err := GetBooks(db)

	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}

	// Convert the books slice to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	book_id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Call the GetBook function to fetch the book from the database
	user, err := GetBook(db, book_id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Convert the book object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	book_id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book := DeleteBook(db, book_id)
	if book != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Book deleted successfully")
}
