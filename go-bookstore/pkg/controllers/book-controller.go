package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/db2103/go-bookstore/pkg/models"
	"github.com/db2103/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func GetBookById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write(res)
	if err2 != nil {
		return
	}
}

func CreateBook(w http.ResponseWriter, req *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(req, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error while parsing")
	}
	deletedBook := models.DeleteBookById(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(deletedBook)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(req, updateBook)

	vars := mux.Vars(req)

	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookById(id)

	if bookDetails.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if bookDetails.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if bookDetails.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
