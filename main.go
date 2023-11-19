//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func main() {

// 	fmt.Println("Starting server on port 8080...")
// 	router := mux.NewRouter()

// 	// Define your endpoints and their corresponding handlers
// 	router.HandleFunc("/", RootHandler).Methods("GET")
// 	// router.HandleFunc("/contacts", CreateContact).Methods("POST")
// 	// Add other endpoints for updating, deleting, getting contacts, etc.

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

func RootHandler(w http.ResponseWriter, r *http.Request) {
	message := "Welcome iotas buds!"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
	// fmt.Fprintf(w, "Hello World!")
}

// Define your HTTP handlers here (e.g., createContact, updateContact, etc.)
// These functions will handle CRUD operations for contacts
// func createContact(w http.ResponseWriter, r *http.Request) {
// 	// Implement contact creation logic here
// }

// Add other handler functions for CRUD operations

// Contact struct represents your contact model
// type Contact struct {
// 	ID        string
// 	FirstName string
// 	LastName  string
// 	// Add other fields as needed (Twitter, Avatar, Notes, etc.)
// }
