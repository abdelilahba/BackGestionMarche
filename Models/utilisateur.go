package models

import "gorm.io/gorm"

type Utilisateur struct {
	gorm.Model
	Nom        string `json:"nom" gorm:"size:255;not null"`
	Prenom     string `json:"prenom" gorm:"size:255;not null"`
	Email      string `json:"email" gorm:"size:255;unique;not null"`
	MotDePasse string `json:"mot_de_passe" gorm:"size:255;not null"`
	Profil     string `json:"profil" gorm:"size:50;not null"`
	Etat       bool   `json:"etat" gorm:"default:true"`
	Service    string `json:"service" gorm:"size:50"`
	Photo      string `json:"photo"`
}
