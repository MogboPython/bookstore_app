package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func insertBook(w http.ResponseWriter, request *http.Request) {
	var book Book
	json.NewDecoder(request.Body).Decode(&book)

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

//  func getUserHandler(w http.ResponseWriter, r *http.Request) {
//     db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
//     if err != nil {
//       panic(err.Error())
//     }
//     defer db.Close()

//     // Get the 'id' parameter from the URL
//     vars := mux.Vars(r)
//     idStr := vars["id"]

//     // Convert 'id' to an integer
//     userID, err := strconv.Atoi(idStr)

//     // Call the GetUser function to fetch the user data from the database
//     user, err := GetUser(db, userID)
//     if err != nil {
//       http.Error(w, "User not found", http.StatusNotFound)
//       return
//     }

//     // Convert the user object to JSON and send it in the response
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(user)
//  }
