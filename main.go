package main

import (
	"fmt"
	"sort"
)

// WordMap représente la structure qui stocke les mots et leurs définitions.
type WordMap map[string]string

// Add ajoute un mot avec sa définition à la map.
func (wm WordMap) Add(word, definition string) {
	wm[word] = definition
}

// Get récupère la définition d'un mot spécifique.
func (wm WordMap) Get(word string) (string, bool) {
	definition, found := wm[word]
	return definition, found
}

// Remove supprime un mot de la map.
func (wm WordMap) Remove(word string) {
	delete(wm, word)
}

// List renvoie la liste triée des mots et de leurs définitions.
func (wm WordMap) List() {
	var words []string
	for word := range wm {
		words = append(words, word)
	}

	sort.Strings(words)

	for _, word := range words {
		fmt.Printf("%s: %s\n", word, wm[word])
	}
}

func main() {
	// Créer une map pour stocker les mots et les définitions.
	wordMap := make(WordMap)

	// Utiliser la méthode Add pour ajouter des mots et des définitions.
	wordMap.Add("Foo", "Bar")
	wordMap.Add("Foo-2", "Bar-2")
	wordMap.Add("Foo-3", "Bar-3")

	// Utiliser la méthode Get pour afficher la définition d'un mot spécifique.
	definition, found := wordMap.Get("Foo-2")
	if found {
		fmt.Printf("Definition of 'Foo-2': %s\n", definition)
	} else {
		fmt.Println("Word not found.")
	}

	// Utiliser la méthode Remove pour supprimer un mot de la map.
	wordMap.Remove("Foo-2")

	// Appeler la méthode List pour obtenir la liste triée des mots et de leurs définitions.
	wordMap.List()
}
