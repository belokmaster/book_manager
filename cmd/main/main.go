package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/belokmaster/book_manager/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRouter(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
