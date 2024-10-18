// routes/router.go
package routes

import (
	"mon-projet-backend/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterDynamicRoutes registers CRUD routes for a given model with optional preloads
func RegisterDynamicRoutes(router *mux.Router, model interface{}, baseRoute string, preloads ...string) {
	// GET: Récupérer tous les enregistrements
	router.HandleFunc("/"+baseRoute, func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAll(model, w, r, preloads...)
	}).Methods("GET")

	// GET: Récupérer un enregistrement par ID
	router.HandleFunc("/"+baseRoute+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetById(model, w, r, preloads...)
	}).Methods("GET")

	// POST: Créer un nouvel enregistrement
	router.HandleFunc("/"+baseRoute, func(w http.ResponseWriter, r *http.Request) {
		controllers.Create(model, w, r)
	}).Methods("POST")

	// PUT: Mettre à jour un enregistrement par ID
	router.HandleFunc("/"+baseRoute+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.Update(model, w, r)
	}).Methods("PUT")

	// DELETE: Supprimer un enregistrement par ID
	router.HandleFunc("/"+baseRoute+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.Delete(model, w, r)
	}).Methods("DELETE")
}
