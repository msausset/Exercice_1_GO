package dictionary

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"

	"github.com/gorilla/mux"
)

// Entrée dans le dictionnaire.
type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

// Stockage mot/définition en mémoire
var entriesMap = make(map[string]Entry)
var entriesMutex sync.RWMutex

// Fonction ajout
func Add(entry Entry) {
	entriesMutex.Lock()
	defer entriesMutex.Unlock()

	entriesMap[entry.Word] = entry
}

// Récupération d'un mot
func Get(word string) (string, bool) {
	entriesMutex.RLock()
	defer entriesMutex.RUnlock()

	entry, found := entriesMap[word]
	return entry.Definition, found
}

// Suppression
func Remove(word string) {
	entriesMutex.Lock()
	defer entriesMutex.Unlock()

	delete(entriesMap, word)
}

// Liste triée par ordre
func List() error {
	entriesMutex.RLock()
	defer entriesMutex.RUnlock()

	var words []string
	for word := range entriesMap {
		words = append(words, word)
	}

	sort.Strings(words)

	for _, word := range words {
		fmt.Printf("%s: %s\n", word, entriesMap[word])
	}

	return nil
}

// Fonction pour ajouter une entrée en mémoire
func addEntry(word, definition string) {
	entriesMutex.Lock()
	defer entriesMutex.Unlock()

	entriesMap[word] = Entry{Word: word, Definition: definition}
}

// Fonction pour supprimer une entrée de la mémoire
func removeEntry(wordToRemove string) {
	entriesMutex.Lock()
	defer entriesMutex.Unlock()

	delete(entriesMap, wordToRemove)
}

// Requete d'ajout
// Requete d'ajout
func AddHandler(w http.ResponseWriter, r *http.Request) {
    var entry Entry
    err := json.NewDecoder(r.Body).Decode(&entry)
    if err != nil {
        log.Println("Erreur lors de la lecture du corps JSON :", err)
        http.Error(w, "Données JSON invalides", http.StatusBadRequest)
        return
    }

    // Validation des données
    if err := validateEntry(entry); err != nil {
        log.Println("Erreur de validation des données :", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    Add(entry)

    // Envoyer une réponse JSON vide et définir le code de statut 201
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("{}"))
}



// Validation des données pour l'ajout d'une entrée
func validateEntry(entry Entry) error {
	if len(entry.Word) < 3 || len(entry.Word) > 30 {
		return errors.New("La longueur du mot doit être comprise entre 3 et 30 caractères")
	}
	if len(entry.Definition) < 10 || len(entry.Definition) > 100 {
		return errors.New("La longueur de la définition doit être comprise entre 10 et 100 caractères")
	}
	return nil
}

// Récupération d'un mot
func GetHandler(w http.ResponseWriter, r *http.Request) {
	word := mux.Vars(r)["word"]
	definition, found := Get(word)
	if !found {
		log.Printf("Mot introuvable : %s\n", word)
		http.Error(w, "Mot introuvable", http.StatusNotFound)
		return
	}

	log.Printf("Mot trouvé : %s, Définition : %s\n", word, definition)

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
	Remove(word)
	w.WriteHeader(http.StatusNoContent)
}

// Liste triée par ordre
func ListHandler(w http.ResponseWriter, r *http.Request) {
	entriesMutex.RLock()
	defer entriesMutex.RUnlock()

	var words []string
	for word := range entriesMap {
		words = append(words, word)
	}

	sort.Strings(words)

	var resultList []map[string]interface{}
	for _, word := range words {
		entry := entriesMap[word]
		resultList = append(resultList, map[string]interface{}{"word": entry.Word, "definition": entry.Definition})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultList)
}
