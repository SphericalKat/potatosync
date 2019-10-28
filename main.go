package main

import (
	"github.com/ATechnoHazard/potatonotes-api/controllers"
	"github.com/ATechnoHazard/potatonotes-api/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication) // use our jwt middleware

	router.HandleFunc("/api/users/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/users/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/users/delete", controllers.Delete).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Listening on port %s\n", port)

	err := http.ListenAndServe(":"+port, router) // launch the middleware, visit localhost:4000/api
	if err != nil {
		log.Fatalln(err)
	}
}
