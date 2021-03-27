package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// init books var as a slice Book struct
var books []Book

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//create a new book
func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update a book
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["ID"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			// book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
			book.ID = item.ID // Mock ID - not safe
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)

}

// delete a book
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["ID"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init Router
	r := mux.NewRouter()

	// Mock Data -@todo -implement DB
	books = append(books, Book{ID: "1", Isbn: "44876", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Snow"}})
	books = append(books, Book{ID: "2", Isbn: "83476", Title: "Book Two", Author: &Author{Firstname: "Aniket", Lastname: "Kumar"}})
	books = append(books, Book{ID: "3", Isbn: "15676", Title: "Book Thr", Author: &Author{Firstname: "Neha", Lastname: "Rai"}})
	books = append(books, Book{ID: "4", Isbn: "34876", Title: "Book Fou", Author: &Author{Firstname: "Sumesh", Lastname: "M"}})
	books = append(books, Book{ID: "5", Isbn: "54876", Title: "Book Fiv", Author: &Author{Firstname: "Ram", Lastname: "Prit"}})

	// Router Handlers(Endpoints)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{ID}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{ID}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{ID}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
