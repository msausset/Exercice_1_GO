// dictionary/dictionary.go
package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// WordMap représente la structure qui stocke les mots et leurs définitions.
type WordMap map[string]string

// Chemin du fichier de données du dictionnaire
const filePath = "dictionary_data.txt"

// Add ajoute un mot avec sa définition au fichier.
func Add(word, definition string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = fmt.Fprintf(writer, "%s=%s\n", word, definition)
	if err != nil {
		return err
	}

	return writer.Flush()
}

// Get récupère la définition d'un mot spécifique depuis le fichier.
func Get(word string) (string, bool) {
	entries, err := readEntriesFromFile()
	if err != nil {
		return "", false
	}

	definition, found := entries[word]
	return definition, found
}

// Remove supprime un mot du fichier.
func Remove(word string) error {
	entries, err := readEntriesFromFile()
	if err != nil {
		return err
	}

	delete(entries, word)

	return writeEntriesToFile(entries)
}

// List renvoie la liste triée des mots et de leurs définitions depuis le fichier.
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

// Fonction utilitaire pour lire les entrées depuis le fichier
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
			entries[parts[0]] = parts[1]
		}
	}

	return entries, scanner.Err()
}

// Fonction utilitaire pour écrire les entrées dans le fichier
func writeEntriesToFile(entries WordMap) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for word, definition := range entries {
		_, err := fmt.Fprintf(writer, "%s=%s\n", word, definition)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
