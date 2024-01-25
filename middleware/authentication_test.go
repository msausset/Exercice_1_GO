// authentication_test.go
package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	// Sauvegarder l'ancienne valeur de validToken pour la restaurer plus tard
	originalValidToken := os.Getenv("validToken")

	// Définir une valeur de validToken pour les tests
	os.Setenv("validToken", "test_token")

	// Restaurer la valeur originale à la fin des tests
	defer func() {
		os.Setenv("validToken", originalValidToken)
	}()

	// Créer un gestionnaire avec le middleware d'authentification
	handler := AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Cas de test avec un jeton valide
	reqValidToken := httptest.NewRequest("GET", "/test", nil)
	reqValidToken.Header.Set("Authorization", "Bearer test_token")
	wValidToken := httptest.NewRecorder()
	handler.ServeHTTP(wValidToken, reqValidToken)

	// Vérifier le statut de la réponse
	assert.Equal(t, http.StatusOK, wValidToken.Code, "Le statut de la réponse attendu doit être OK")

	// Cas de test avec un jeton invalide
	reqInvalidToken := httptest.NewRequest("GET", "/test", nil)
	reqInvalidToken.Header.Set("Authorization", "Bearer invalid_token")
	wInvalidToken := httptest.NewRecorder()
	handler.ServeHTTP(wInvalidToken, reqInvalidToken)

	// Vérifier le statut de la réponse
	assert.Equal(t, http.StatusUnauthorized, wInvalidToken.Code, "Le statut de la réponse attendu doit être Unauthorized")

	// Cas de test sans jeton
	reqNoToken := httptest.NewRequest("GET", "/test", nil)
	wNoToken := httptest.NewRecorder()
	handler.ServeHTTP(wNoToken, reqNoToken)

	// Vérifier le statut de la réponse
	assert.Equal(t, http.StatusUnauthorized, wNoToken.Code, "Le statut de la réponse attendu doit être Unauthorized")
}
