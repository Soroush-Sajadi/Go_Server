package main

import(
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID        string `json:"id"`
	Isbn      string `json:"isbn"`
	Title     string `json:"title"`
	Author    *Author `json:"author"`
}

type Author struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
}

var books []Book

func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(books)
}
func getbook(w http.ResponseWriter, r *http.Request) {

}
func createbook(w http.ResponseWriter, r *http.Request) {

}
func updatebook(w http.ResponseWriter, r *http.Request) {

}
func deletebook(w http.ResponseWriter, r *http.Request) {

}

func main () {
	route := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn:"44232", Title:"1984", Author: &Author{Firstname:"Goerge", Lastname:"Orvel"}})
	books = append(books, Book{ID: "2", Isbn:"41232", Title:"Jungle", Author: &Author{Firstname:"Tom", Lastname:"Sea"}})
	books = append(books, Book{ID: "3", Isbn:"44122", Title:"Winston", Author: &Author{Firstname:"Jack", Lastname:"Nail"}})

	route.HandleFunc("/api/books", getbooks).Methods("GET")
	route.HandleFunc("/api/book/{id}", getbook).Methods("GET")
	route.HandleFunc("/api/book/add", createbook).Methods("POST")
	route.HandleFunc("/api/book/update/{id}", updatebook).Methods("PUT")
	route.HandleFunc("/api/book/delete/{id}", deletebook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", route))
}