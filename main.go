package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"exo1/route"
)

func main() {
	// Création du routeur 
	router := mux.NewRouter()

	// Initialisation des routes
	route.SetupRoutes(router)

	// Démarre le serveur
	log.Fatal(http.ListenAndServe(":8080", router))
}
