package routes

import (
	"github.com/gorilla/mux"
	"github.com/samarth8765/Project_3/pkg/controllers"
)

var RegisterBookStores = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("GET")
	router.HandleFunc("/book", controllers.GetBooks).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("UPDATE")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
