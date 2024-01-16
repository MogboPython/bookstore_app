package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func addApproutes(route *mux.Router) {
	// setStaticFolder(route)
	// route.HandleFunc("/", renderHome)
	// route.HandleFunc("/user", getUsers).Methods("GET")
	// route.HandleFunc("/user", insertUser).Methods("POST")
	// route.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	// route.HandleFunc("/user", updateUser).Methods("PUT")
	fmt.Println("Routes are Loded.")
}
