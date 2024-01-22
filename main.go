package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"exo1/route"
	"exo1/middleware"
)

func main() {
	// Créer le routeur Gorilla Mux
	router := mux.NewRouter()

	// Utilisation du middleware
	router.Use(middleware.LoggingMiddleware)

	// Initialisation des routes
	route.SetupRoutes(router)

	// Démarre le serveur HTTP
	log.Fatal(http.ListenAndServe(":8080", router))
}
