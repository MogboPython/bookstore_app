package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func addApproutes(route *mux.Router) {
	// setStaticFolder(route)
	// route.HandleFunc("/", renderHome)
	route.HandleFunc("/books", getBooks).Methods("GET")
	route.HandleFunc("/books/{id}", getBookHandler).Methods("GET")
	route.HandleFunc("/submit_book", insertBookHandler).Methods("POST")
	route.HandleFunc("/books/{id}", deleteBookHandler).Methods("DELETE")
	// route.HandleFunc("/book", updateBook).Methods("PUT")
	// route.HandleFunc("/book/{id}", borrowBook).Methods("POST")
	fmt.Println("Routes are Loaded.")
}
