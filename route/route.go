package route

import (
	"github.com/gorilla/mux"
	"exo1/dictionary"
	"exo1/middleware"
)

// SetupRoutes initialise les routes avec Gorilla Mux.
func SetupRoutes(r *mux.Router) {
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	r.HandleFunc("/add", dictionary.AddHandler).Methods("POST")
	r.HandleFunc("/get/{word}", dictionary.GetHandler).Methods("GET")
	r.HandleFunc("/remove/{word}", dictionary.RemoveHandler).Methods("DELETE")
	r.HandleFunc("/list", dictionary.ListHandler).Methods("GET")
}
