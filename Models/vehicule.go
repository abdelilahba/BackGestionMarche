package models

import "gorm.io/gorm"

type Vehicule struct {
	gorm.Model
	Immatriculation string      `json:"immatriculation" gorm:"size:50;unique;not null"`
	Marque          string      `json:"marque" gorm:"size:100;not null"`
	Modele          string      `json:"modele" gorm:"size:100;not null"`
	DateAchat       string      `json:"date_achat" gorm:"not null"`
	Etat            string      `json:"etat" gorm:"size:50;not null"`
	Kilometrage     uint        `json:"kilometrage" gorm:"not null"`
	DateReforme     string      `json:"date_reforme"`
	TypeAcquisition string      `json:"type_acquisition" gorm:"size:50"`
	FournisseurID   uint        `json:"fournisseur_id"` // FK vers Fournisseur
	Fournisseur     Fournisseur `gorm:"foreignKey:FournisseurID" json:"fournisseur"`
}
