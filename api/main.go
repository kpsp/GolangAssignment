package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	Id     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Price  string `json:"Price"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleBooks")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, book := range Books {
		if book.Id == key {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewBook")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteBook")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateBook")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedbook Book
	json.Unmarshal(reqBody, &updatedbook)

	vars := mux.Vars(r)
	id := vars["id"]

	var tempbooks []Book
	var index int

	for ind, book := range Books {
		if book.Id == id {
			tempbooks = append(Books[:ind], updatedbook)
			index = ind
		}
	}

	Books = append(tempbooks, Books[index+1:]...)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/books", returnAllBooks)
	myRouter.HandleFunc("/book", createNewBook).Methods("POST")
	myRouter.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	myRouter.HandleFunc("/book/{id}", returnSingleBook)
	log.Fatal(http.ListenAndServe(":6004", myRouter))
}

func main() {
	Books = []Book{
		{Id: "1", Title: "Adventures of Duck and Goose", Author: "Sir Quackalot", Price: "10.99"},
		{Id: "2", Title: "The Return of Duck and Goose", Author: "Sir Quackalot", Price: "19.99"},
		{Id: "3", Title: "More Fun with Duck and Goose", Author: "Sir Quackalot", Price: "12.99"},
		{Id: "4", Title: "Duck and Goose on Holiday", Author: "Sir Quackalot", Price: "11.99"},
		{Id: "5", Title: "My Friend is a Good Duck", Author: "A. Parrot", Price: "14.99"},
		{Id: "6", Title: "Notes on 'Duck and Goose'", Author: "Prof Macaw", Price: "8.99"},
		{Id: "7", Title: "'Duck and Goose' Cheat Sheet", Author: "Polly Parrot", Price: "5.99"},
		{Id: "8", Title: "'Duck and Goose': an allegory", Author: "BorIng Man", Price: "59.99"},
	}
	handleRequests()
}
