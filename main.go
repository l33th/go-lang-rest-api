package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Book Struct (Model)
type Book struct {
	ID		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}

// Author Struct
type Author struct {
	Firstname	string	`json:"firstname"`
	Lastname	string `json:"lastname"`
}

// Init books variable as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Fatalln(err, "Error")
	}
}
// Get A Book
func getBook(w http.ResponseWriter, r *http.Request) {}
// Create A Book
func createBook(w http.ResponseWriter, r *http.Request) {}
// Update A Book
func updateBook(w http.ResponseWriter, r *http.Request) {}
// Delete A Book
func deleteBook(w http.ResponseWriter, r *http.Request) {}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "681565", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
