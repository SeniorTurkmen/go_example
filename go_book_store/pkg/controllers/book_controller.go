package controller

import (
	"example/go_bookstore/pkg/config"
	"example/go_bookstore/pkg/models"
	"example/go_bookstore/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		utils.ResponsWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponsWithJson(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, bookError := strconv.ParseInt(params["id"], 0, 0)
	if bookError != nil {
		utils.ResponsWithError(w, http.StatusBadRequest, "Id is not valid")
		return
	}
	book, err := models.GetBookById(id)
	if err != nil {
		utils.ResponsWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	utils.ResponsWithJson(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var NewBook models.Book

	if err := utils.ParseBody(r, &NewBook); err != nil {
		utils.ResponsWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	book, err := NewBook.CreateBook()
	if err != nil {
		utils.ResponsWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponsWithJson(w, http.StatusCreated, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, bookError := strconv.ParseInt(params["id"], 0, 0)
	if bookError != nil {
		utils.ResponsWithError(w, http.StatusBadRequest, "Id is not valid")
		return
	}
	var tempBook models.Book

	if err := utils.ParseBody(r, &tempBook); err != nil {
		utils.ResponsWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		utils.ResponsWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	if tempBook.Name != "" {
		book.Name = tempBook.Name
	}
	if tempBook.Author != "" {
		book.Author = tempBook.Author
	}
	if tempBook.Publication != "" {
		book.Publication = tempBook.Publication
	}

	config.GetDB().Save(&book)

	utils.ResponsWithJson(w, http.StatusOK, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, bookError := strconv.ParseInt(params["id"], 0, 0)
	if bookError != nil {
		utils.ResponsWithError(w, http.StatusBadRequest, "Id is not valid")
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		utils.ResponsWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	_, err = models.DeleteBook(id)
	if err != nil {
		utils.ResponsWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	currentTime := time.Now()
	book.DeletedAt = &currentTime
	utils.ResponsWithJson(w, http.StatusOK, book)
}
