package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "strconv"
	// "encoding/json"
)

func main() {

	fmt.Println("Server will start at http://localhost:8000/")
	connectDatabase()
	route := mux.NewRouter()
	addApproutes(route)
	log.Fatal(http.ListenAndServe(":8000", route))
}
