package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

// Fonction d'authentification
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Charge la variable d'environnement du token depuis le fichier .env
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Erreur lors du chargement du fichier .env :", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Récupération du jeton dans le fichier .env
		validToken := os.Getenv("validToken")

		// Vérifie le jeton dans l'en-tête
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("Token manquant")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Extraire le jeton du format "Bearer <token>"
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 || splitToken[1] != validToken {
			fmt.Println("Token invalide")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Si le jeton est valide, passer à l'handler suivant
		next.ServeHTTP(w, r)
	})
}
