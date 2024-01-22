package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"log"
)

// LoggingMiddleware enregistre chaque requête dans un fichier log
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logEntry := fmt.Sprintf("%s %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		LogToFile(logEntry)
		next.ServeHTTP(w, r)
	})
}

// LogToFile enregistre l'entrée dans un fichier log
func LogToFile(logEntry string) {
	file, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Erreur lors de l'ouverture du fichier journal :", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(logEntry)
	if err != nil {
		log.Println("Erreur lors de l'écriture dans le fichier journal :", err)
	}
}
