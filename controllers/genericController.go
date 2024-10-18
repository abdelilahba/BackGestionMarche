package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	models "mon-projet-backend/Models"
	"mon-projet-backend/config"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GetAll fetches all records of a given model with optional preloads
func GetAll(model interface{}, w http.ResponseWriter, r *http.Request, preloads ...string) {
	log.Printf("GetAll: Applying preloads: %v", preloads)

	// Déterminer le type du modèle
	modelType := reflect.TypeOf(model)
	if modelType.Kind() != reflect.Ptr {
		log.Printf("GetAll: model should be a pointer to a struct")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	modelElemType := modelType.Elem()

	// Créer une instance de slice pour stocker les enregistrements
	sliceType := reflect.SliceOf(modelElemType)
	slicePtr := reflect.New(sliceType)
	sliceInterface := slicePtr.Interface()

	// Commencer avec la base de données
	db := config.DB

	// Appliquer les préchargements si spécifiés
	for _, preload := range preloads {
		log.Printf("GetAll: Preloading association: %s", preload)
		db = db.Preload(preload)
	}

	// Récupérer tous les enregistrements avec les préchargements
	if err := db.Find(sliceInterface).Error; err != nil {
		log.Printf("GetAll: Error fetching records: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Log des données récupérées
	data, err := json.Marshal(slicePtr.Elem().Interface())
	if err != nil {
		log.Printf("GetAll: Error marshalling data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("GetAll: Data retrieved: %s", string(data))

	// Encoder les données en JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(slicePtr.Elem().Interface()); err != nil {
		log.Printf("GetAll: Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetById fetches a single record by ID with optional preloads
func GetById(model interface{}, w http.ResponseWriter, r *http.Request, preloads ...string) {
	log.Printf("GetById: Applying preloads: %v", preloads)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("GetById: Invalid ID: %v", err)
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Créer une instance du modèle
	instance := reflect.New(reflect.TypeOf(model).Elem()).Interface()

	// Commencer avec la base de données
	db := config.DB

	// Appliquer les préchargements si spécifiés
	for _, preload := range preloads {
		log.Printf("GetById: Preloading association: %s", preload)
		db = db.Preload(preload)
	}

	// Récupérer l'enregistrement avec les préchargements
	if err := db.First(instance, id).Error; err != nil {
		log.Printf("GetById: Error fetching record by ID: %v", err)
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Enregistrement non trouvé", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Encoder les données en JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(instance); err != nil {
		log.Printf("GetById: Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Create creates a new record
func Create(model interface{}, w http.ResponseWriter, r *http.Request) {
	instance := reflect.New(reflect.TypeOf(model).Elem()).Interface()

	// Décoder le corps de la requête
	if err := json.NewDecoder(r.Body).Decode(instance); err != nil {
		log.Printf("Create: Error decoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validation supplémentaire (exemple pour Vehicule)
	vehicule, ok := instance.(*models.Vehicule)
	if ok {
		if vehicule.FournisseurID == 0 {
			http.Error(w, "FournisseurID est requis", http.StatusBadRequest)
			return
		}
		// Ajoutez d'autres validations si nécessaire
	}

	// Créer l'enregistrement dans la base de données
	if err := config.DB.Create(instance).Error; err != nil {
		log.Printf("Create: Error creating record: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encoder les données en JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(instance); err != nil {
		log.Printf("Create: Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Update updates an existing record by ID
func Update(model interface{}, w http.ResponseWriter, r *http.Request) {
	log.Printf("Update: Applying preloads: []") // Les préchargements ne sont pas utilisés ici

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Update: Invalid ID: %v", err)
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Créer une instance du modèle
	instance := reflect.New(reflect.TypeOf(model).Elem()).Interface()

	// Trouver l'enregistrement existant
	if err := config.DB.First(instance, id).Error; err != nil {
		log.Printf("Update: Error fetching record by ID: %v", err)
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Enregistrement non trouvé", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Décoder la nouvelle donnée
	if err := json.NewDecoder(r.Body).Decode(instance); err != nil {
		log.Printf("Update: Error decoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sauvegarder les modifications
	if err := config.DB.Save(instance).Error; err != nil {
		log.Printf("Update: Error saving record: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encoder les données en JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(instance); err != nil {
		log.Printf("Update: Error encoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Delete deletes a record by ID
func Delete(model interface{}, w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete: Applying preloads: []") // Les préchargements ne sont pas utilisés ici

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Delete: Invalid ID: %v", err)
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Créer une instance du modèle
	instance := reflect.New(reflect.TypeOf(model).Elem()).Interface()

	// Check if the record exists before deleting
	if err := config.DB.First(instance, id).Error; err != nil {
		log.Printf("Delete: Error fetching record by ID: %v", err)
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Enregistrement non trouvé", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Delete the record
	if err := config.DB.Delete(instance, id).Error; err != nil {
		log.Printf("Delete: Error deleting record: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return 204 No Content on successful deletion
	w.WriteHeader(http.StatusNoContent)
}
