package models

import "gorm.io/gorm"

type Marche struct {
	gorm.Model
	NumeroMarche      string      `json:"numero_marche" gorm:"size:50;unique;not null"`
	AnneeMarche       uint        `json:"annee_marche" gorm:"not null"`
	ObjetMarche       string      `json:"objet_marche" gorm:"not null"`
	Demandeur         string      `json:"demandeur"`
	Service           string      `json:"service"`
	MontantMarche     float64     `json:"montant_marche" gorm:"not null"`
	TitulaireMarcheID uint        `json:"titulaire_marche_id" gorm:"not null"` // FK vers Fournisseur
	TitulaireMarche   Fournisseur `gorm:"foreignKey:TitulaireMarcheID"`
}
