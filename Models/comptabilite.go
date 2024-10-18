package models

import "gorm.io/gorm"

type Comptabilite struct {
	gorm.Model
	Montant       float64     `json:"montant" gorm:"not null"`
	TypeOperation string      `json:"type_operation" gorm:"size:50;not null"`
	FournisseurID uint        `json:"fournisseur_id" gorm:"not null"` // FK vers Fournisseur
	Fournisseur   Fournisseur `gorm:"foreignKey:FournisseurID"`
}
