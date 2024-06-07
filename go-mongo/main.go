package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/etharrra/go-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	fmt.Println("Starting server at localhost:9000")
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	fmt.Println("DB connection Success!")
	return s
}
