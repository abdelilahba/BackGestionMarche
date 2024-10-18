package main

import (
	"log"
	"mon-projet-backend/config"
	"mon-projet-backend/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Connexion à la base de données
	config.ConnectDB()

	// Création d'un nouveau routeur
	router := mux.NewRouter()

	// Enregistrement des routes
	routes.RegisterRoutes(router)

	// Configuration des options CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Lancement du serveur sur le port 8000 avec le middleware CORS
	log.Println("Serveur démarré sur le port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
