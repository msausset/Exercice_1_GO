package dictionary

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"github.com/gorilla/mux"
)

// Entrée dans le dictionnaire.
type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

// Stockage mot/définition
type WordMap map[string]Entry

// Fichier texte
const filePath = "dictionary_data.txt"

// Channels pour les opérations d'ajout et de suppression
var (
	addChan    = make(chan entryOperation)
	removeChan = make(chan string)
	wg         sync.WaitGroup
)

// Structure pour l'ajout d'un mot
type entryOperation struct {
	word       string
	definition string
}

// Goroutines
func init() {
	go processAddOperations()
	go processRemoveOperations()
}

// Fonction ajout
func processAddOperations() {
	for {
		select {
		case entry := <-addChan:
			addEntryToFile(entry.word, entry.definition)
		}
	}
}

// Fonction suppression
func processRemoveOperations() {
	for {
		select {
		case word := <-removeChan:
			removeEntryFromFile(word)
		}
	}
}

// Ajout avec channel
func Add(entry Entry) {
	addChan <- entryOperation{entry.Word, entry.Definition}
}

// Récupération d'un mot
func Get(word string) (string, bool) {
	entries, err := readEntriesFromFile()
	if err != nil {
		return "", false
	}

	definition, found := entries[word]
	return definition.Definition, found
}

// Suppression avec channel
func Remove(word string) {
	removeChan <- word
}

// Liste triée par ordre
func List() error {
	entries, err := readEntriesFromFile()
	if err != nil {
		return err
	}

	var words []string
	for word := range entries {
		words = append(words, word)
	}

	sort.Strings(words)

	for _, word := range words {
		fmt.Printf("%s: %s\n", word, entries[word])
	}

	return nil
}

// Fonction pour lire les entrées depuis le fichier
func readEntriesFromFile() (WordMap, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	entries := make(WordMap)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			var entry Entry
			err := json.Unmarshal([]byte(parts[1]), &entry)
			if err == nil {
				entries[parts[0]] = entry
			}
		}
	}

	return entries, scanner.Err()
}

// Fonction pour écrire les entrées dans le fichier
func writeEntriesToFile(entries WordMap) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for word, entry := range entries {
		jsonData, err := json.Marshal(entry)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(writer, "%s=%s\n", word, jsonData)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// Fonction pour ajouter une entrée au fichier
func addEntryToFile(word, definition string) {
	wg.Add(1)
	defer wg.Done()

	entries, err := readEntriesFromFile()
	if err != nil {
		log.Println("Erreur lors de la lecture des entrées :", err)
		return
	}

	entries[word] = Entry{Word: word, Definition: definition}

	err = writeEntriesToFile(entries)
	if err != nil {
		log.Println("Erreur lors de l'écriture des entrées :", err)
		return
	}
}

// Fonction pour supprimer une entrée du fichier
func removeEntryFromFile(wordToRemove string) {
	wg.Add(1)
	defer wg.Done()

	entries, err := readEntriesFromFile()
	if err != nil {
		log.Println("Erreur lors de la lecture des entrées :", err)
		return
	}

	delete(entries, wordToRemove)

	err = writeEntriesToFile(entries)
	if err != nil {
		log.Println("Erreur lors de l'écriture des entrées :", err)
		return
	}
}

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

	Add(Entry{Word: entry.Word, Definition: entry.Definition})
	w.WriteHeader(http.StatusCreated)
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
	Remove(word)
	w.WriteHeader(http.StatusNoContent)
}

// Liste triée par ordre
func ListHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := readEntriesFromFile()
	if err != nil {
		log.Println("Erreur lors de la récupération des entrées :", err)
		http.Error(w, "Erreur lors de la récupération des entrées", http.StatusInternalServerError)
		return
	}

	var words []string
	for word := range entries {
		words = append(words, word)
	}

	sort.Strings(words)

	var resultList []map[string]interface{}
	for _, word := range words {
		entry := entries[word]
		resultList = append(resultList, map[string]interface{}{"word": entry.Word, "definition": entry.Definition})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultList)
}
