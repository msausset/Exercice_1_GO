package route

import (
	"encoding/json"
	"exo1/dictionary"
	"github.com/gorilla/mux"
	"net/http"
)

// Initialisation des routes Gorilla Mux
func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/add", AddHandler).Methods("POST")
	r.HandleFunc("/get/{word}", GetHandler).Methods("GET")
	r.HandleFunc("/remove/{word}", RemoveHandler).Methods("DELETE")
}

// Requete d'ajout
func AddHandler(w http.ResponseWriter, r *http.Request) {
	var entry dictionary.Entry
	err := json.NewDecoder(r.Body).Decode(&entry) 
	if err != nil {
		http.Error(w, "Données JSON invalides", http.StatusBadRequest)
		return
	}

	dictionary.Add(dictionary.Entry{Word: entry.Word, Definition: entry.Definition})
	w.WriteHeader(http.StatusCreated)
}

// Récupération du mot
func GetHandler(w http.ResponseWriter, r *http.Request) {
	word := mux.Vars(r)["word"]
	definition, found := dictionary.Get(word)
	if !found {
		http.Error(w, "Mot introuvable", http.StatusNotFound)
		return
	}

	response := struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}{
		Word:       word,
		Definition: definition,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Suppression
func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	word := mux.Vars(r)["word"]
	dictionary.Remove(word)
	w.WriteHeader(http.StatusNoContent)
}
