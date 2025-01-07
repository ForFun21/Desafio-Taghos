package main

import (
	"desafio_taghos_livros/database"
	"desafio_taghos_livros/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
