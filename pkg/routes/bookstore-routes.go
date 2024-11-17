package routes
// This package is responsible for configuring routes for the web app.
import (
	"github.com/gorrila/mux"
	"./controllers"

)

var RegisterBookStoreRouter = func(router *mux.Router) {
	// the router for create a new book
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	// the router for getting all the books
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	// the router for getting the book by its id
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	// the router for updating the book by its id
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	// the router for deleting the book by its id
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}