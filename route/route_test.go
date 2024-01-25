package route

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"exo1/dictionary"
	"strings"
)

func TestIntegrationAddRoute(t *testing.T) {
	// Crée un routeur Gorilla Mux
	router := mux.NewRouter()
	SetupRoutes(router)

	// Crée une requête HTTP POST simulée
	jsonEntry := []byte(`{"word": "test", "definition": "This is a test"}`)
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonEntry))
	if err != nil {
		t.Fatal("Erreur lors de la création de la requête :", err)
	}

	// Défini l'en-tête d'authentification avec un jeton valide
	req.Header.Set("Authorization", "Bearer 2906199918091999")

	// Crée un enregistreur de réponse HTTP
	recorder := httptest.NewRecorder()

	// Passe la requête simulée au routeur
	router.ServeHTTP(recorder, req)

	// Vérifie le code de statut et la réponse
	assert.Equal(t, http.StatusCreated, recorder.Code, "Le code de statut de la réponse doit être Created (201)")
	assert.Equal(t, "{}", recorder.Body.String(), "La réponse attendue doit être une chaîne JSON vide")
}

func TestIntegrationGetRoute(t *testing.T) {
	// Crée un routeur Gorilla Mux
	router := mux.NewRouter()
	SetupRoutes(router)

	// Ajoute le mot au dictionnaire avant d'effectuer la requête GET
	dictionary.Add(dictionary.Entry{Word: "Test2132", Definition: "Definition2132"})

	// Crée une requête HTTP GET simulée
	req, err := http.NewRequest("GET", "/get/Test2132", nil)
	assert.NoError(t, err, "Erreur lors de la création de la requête GET")

	// Ajoute le jeton d'authentification valide
	req.Header.Set("Authorization", "Bearer 2906199918091999")

	// Crée un enregistreur de réponse (response recorder) pour capturer la réponse
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	// Vérifie le code de statut de la réponse
	assert.Equal(t, http.StatusOK, rr.Code, "Le code de statut de la réponse doit être OK")

	// Vérifie le corps de la réponse (json)
	expectedResponse := `{"word":"Test2132","definition":"Definition2132"}`
	actualResponse := strings.TrimSpace(rr.Body.String())

	assert.Equal(t, expectedResponse, actualResponse, "La réponse attendue doit être %s, mais a reçu %s", expectedResponse, actualResponse)
}

func TestIntegrationRemoveRoute(t *testing.T) {
	// Crée un routeur Gorilla Mux
	router := mux.NewRouter()
	SetupRoutes(router)

	// Ajoute le mot au dictionnaire avant d'effectuer la requête DELETE
	dictionary.Add(dictionary.Entry{Word: "WordToRemove", Definition: "DefinitionToRemove"})

	// Crée une requête HTTP DELETE simulée
	req, err := http.NewRequest("DELETE", "/remove/WordToRemove", nil)
	assert.NoError(t, err, "Erreur lors de la création de la requête DELETE")

	// Ajoute le jeton d'authentification valide
	req.Header.Set("Authorization", "Bearer 2906199918091999")

	// Crée un enregistreur de réponse (response recorder) pour capturer la réponse
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	// Vérifie le code de statut de la réponse
	assert.Equal(t, http.StatusNoContent, rr.Code, "Le code de statut de la réponse doit être No Content")

	// Vérifie que le mot a été supprimé du dictionnaire
	_, found := dictionary.Get("WordToRemove")
	assert.False(t, found, "Le mot devrait être supprimé du dictionnaire")
}

