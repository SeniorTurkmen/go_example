package routes

import (
	controller "example/go_bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

var ResgitserBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/api/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controller.DeleteBook).Methods("DELETE")
}
