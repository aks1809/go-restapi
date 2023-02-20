package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Author struct (Model)
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var Books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for _, item := range Books {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	var emptyBook Book
	json.NewEncoder(w).Encode(emptyBook)
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	Books = append(Books, newBook)
	json.NewEncoder(w).Encode(newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for index, item := range Books {
		if item.ID == id {
			var updatedBook Book
			json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = id
			(Books[index]) = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	var emptyBook Book
	json.NewEncoder(w).Encode(emptyBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for index, item := range Books {
		if item.ID == id {
			var deletedBook Book = item
			Books = append(Books[:index], Books[index+1:]...)
			json.NewEncoder(w).Encode(deletedBook)
			return
		}
	}
	var emptyBook Book
	json.NewEncoder(w).Encode(emptyBook)
}

func main() {
	// Init router
	r := mux.NewRouter()

	book1 := Book{ID: "1", Isbn: "12345", Title: "New Book 1", Author: &Author{FirstName: "Akshay", LastName: "Chaturvedi"}}
	book2 := Book{ID: "2", Isbn: "56789", Title: "New Book 2", Author: &Author{FirstName: "Sonal", LastName: "Chaturvedi"}}
	Books = append(Books, book1)
	Books = append(Books, book2)

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}
