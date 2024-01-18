// main.go
package main

import (
    "fmt"
    "exo1/dictionary"
)

func main() {
    // Créer une map pour stocker les mots et les définitions en utilisant le package dictionary.
    wordMap := make(dictionary.WordMap)

    // Utiliser la méthode Add du package dictionary pour ajouter des mots et des définitions.
    wordMap.Add("foo", "bar")
    wordMap.Add("foo-2", "bar-2")
    wordMap.Add("foo-3", "bar-3")

    // Utiliser la méthode Get du package dictionary pour afficher la définition d'un mot spécifique.
    definition, found := wordMap.Get("foo")
    if found {
        fmt.Printf("Definition of 'foo': %s\n", definition)
    } else {
        fmt.Println("Word not found.")
    }

    // Utiliser la méthode Remove du package dictionary pour supprimer un mot de la map.
    wordMap.Remove("foo-2")

    // Appeler la méthode List du package dictionary pour obtenir la liste triée des mots et de leurs définitions.
    wordMap.List()
}
