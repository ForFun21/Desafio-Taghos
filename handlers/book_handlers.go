package handlers

import (
	"desafio_taghos_livros/database"
	"desafio_taghos_livros/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Entrada inválida.", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO books (title, author, category, synopsis) VALUES ($1, $2, $3, $4) RETURNING id`
	err := database.DB.QueryRow(query, book.Title, book.Author, book.Category, book.Synopsis).Scan(&book.ID)
	if err != nil {
		http.Error(w, "Erro ao criar o livro.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var book models.Book
	query := `SELECT id, title, author, category, synopsis FROM books WHERE id = $1`
	row := database.DB.QueryRow(query, id)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Synopsis); err != nil {
		http.Error(w, "Livro não encontrado.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, title, author, category, synopsis FROM books`
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "Erro ao buscar livros.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Synopsis); err != nil {
			http.Error(w, "Erro ao analisar os dados do livro.", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Entrada inválida.", http.StatusBadRequest)
		return
	}

	query := `UPDATE books SET title=$1, author=$2, category=$3, synopsis=$4 WHERE id=$5`
	_, err := database.DB.Exec(query, book.Title, book.Author, book.Category, book.Synopsis, id)
	if err != nil {
		http.Error(w, "Erro ao atualizar o livro.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := `DELETE FROM books WHERE id=$1`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Erro ao deletar o livro.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
