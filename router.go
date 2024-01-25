package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func addApproutes(route *mux.Router) {
	// setStaticFolder(route)
	// route.HandleFunc("/", renderHome)
	route.HandleFunc("/books", getBooks).Methods("GET")
	// route.HandleFunc("/books/{id}", getBook).Methods("GET")
	route.HandleFunc("/submit_book", insertBook).Methods("POST")
	// route.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	// route.HandleFunc("/book", updateBook).Methods("PUT")
	// route.HandleFunc("/book/{id}", borrowBook).Methods("POST")
	fmt.Println("Routes are Loaded.")
}
