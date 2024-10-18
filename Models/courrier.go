package models

import "gorm.io/gorm"

type Courrier struct {
	gorm.Model
	Numero       string `json:"numero" gorm:"size:50;unique;not null"`
	Objet        string `json:"objet" gorm:"size:255;not null"`
	DateEnvoi    string `json:"date_envoi" gorm:"not null"`
	Destinataire string `json:"destinataire" gorm:"size:255"`
	Type         string `json:"type" gorm:"size:50;not null"`
	FichierJoint string `json:"fichier_joint" gorm:"size:255"`
	Service      string `json:"service" gorm:"size:255;not null"`
	Statut       string `json:"statut" gorm:"size:50;not null"`
}
