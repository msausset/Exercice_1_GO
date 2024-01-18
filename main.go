// main.go
package main

import (
	"fmt"
	"exo1/dictionary"
)

func main() {
	// Utiliser la méthode Add du package dictionary pour ajouter des mots et des définitions.
	err := dictionary.Add("foo", "bar")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
	}

	err = dictionary.Add("foo-2", "bar-2")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
	}

	err = dictionary.Add("foo-3", "bar-3")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
	}

	// Utiliser la méthode Get du package dictionary pour afficher la définition d'un mot spécifique.
	definition, found := dictionary.Get("foo")
	if found {
		fmt.Printf("Definition of 'foo': %s\n", definition)
	} else {
		fmt.Println("Word not found.")
	}

	// Utiliser la méthode Remove du package dictionary pour supprimer un mot du fichier.
	err = dictionary.Remove("foo-2")
	if err != nil {
		fmt.Println("Erreur lors de la suppression :", err)
	}

	// Appeler la méthode List du package dictionary pour obtenir la liste triée des mots et de leurs définitions depuis le fichier.
	err = dictionary.List()
	if err != nil {
		fmt.Println("Erreur lors de l'affichage de la liste :", err)
	}
}
