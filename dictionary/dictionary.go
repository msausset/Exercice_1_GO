// dictionary/dictionary.go
package dictionary

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
