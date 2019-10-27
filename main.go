package main

import (
	"github.com/ATechnoHazard/potatonotes-api/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // use our jwt middleware

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Using port %s\n", port)

	err := http.ListenAndServe(":"+port, router) // launch the app, visit localhost:4000/api
	if err != nil {
		log.Fatalln(err)
	}
}
