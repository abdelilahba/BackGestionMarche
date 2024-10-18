// routes/router.go
package routes

import (
	models "mon-projet-backend/Models"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// Fournisseurs
	RegisterDynamicRoutes(router, &models.Fournisseur{}, "fournisseurs")

	// Utilisateurs
	RegisterDynamicRoutes(router, &models.Utilisateur{}, "utilisateurs")

	// Marchés
	RegisterDynamicRoutes(router, &models.Marche{}, "marches")

	// Véhicules avec préchargement de Fournisseur
	RegisterDynamicRoutes(router, &models.Vehicule{}, "vehicules", "Fournisseur")

	// Comptabilité
	RegisterDynamicRoutes(router, &models.Comptabilite{}, "comptabilite")

	// Rapports
	RegisterDynamicRoutes(router, &models.Rapport{}, "rapports")

	// Courriers
	RegisterDynamicRoutes(router, &models.Courrier{}, "courriers")
}
