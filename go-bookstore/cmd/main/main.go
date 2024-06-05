package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/etharrra/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server at localhost:1010")
	log.Fatal(http.ListenAndServe("localhost:1010", r))
}
