package models

import "gorm.io/gorm"

type Rapport struct {
	gorm.Model
	Type        string `json:"type" gorm:"size:50"`
	Date        string `json:"date" gorm:"not null"`
	Description string `json:"description"`
}
