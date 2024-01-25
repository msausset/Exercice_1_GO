// dictionary_test.go
package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Ajout d'une entrée
	entry := Entry{Word: "TestWord", Definition: "TestDefinition"}
	Add(entry)

	// Récupération de l'entrée ajoutée
	retrievedDefinition, found := Get(entry.Word)

	// Utilisation de assert pour les vérifications
	assert.True(t, found, "La fonction Add n'a pas ajouté correctement l'entrée")
	assert.Equal(t, entry.Definition, retrievedDefinition, "La fonction Add n'a pas ajouté correctement l'entrée")
}

func TestGet(t *testing.T) {
	// Ajout d'une entrée
	entry := Entry{Word: "TestWord", Definition: "TestDefinition"}
	Add(entry)

	// Récupération de l'entrée ajoutée
	retrievedDefinition, found := Get(entry.Word)

	// Utilisation de assert pour les vérifications
	assert.True(t, found, "La fonction Get n'a pas récupéré correctement l'entrée existante")
	assert.Equal(t, entry.Definition, retrievedDefinition, "La fonction Get n'a pas récupéré correctement l'entrée.")
}

func TestRemove(t *testing.T) {
	// Ajout d'une entrée
	entry := Entry{Word: "TestWord", Definition: "TestDefinition"}
	Add(entry)

	// Suppression de l'entrée ajoutée
	Remove(entry.Word)

	// Utilisation de assert pour les vérifications
	_, found := Get(entry.Word)
	assert.False(t, found, "La fonction Remove n'a pas supprimé correctement l'entrée")
}

