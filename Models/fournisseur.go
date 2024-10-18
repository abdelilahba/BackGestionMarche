package models

import "gorm.io/gorm"

type Fournisseur struct {
	gorm.Model
	Nom                  string `json:"nom" gorm:"size:255;not null"`
	Adresse              string `json:"adresse" gorm:"size:255"`
	Telephone            string `json:"telephone" gorm:"size:50"`
	Email                string `json:"email" gorm:"size:255"`
	ContactInterlocuteur string `json:"contact_interlocuteur" gorm:"size:255"`
	Fax1                 string `json:"fax1"`
	Fax2                 string `json:"fax2"`
	Remarques            string `json:"remarques"`
	NumLibelle           string `json:"num_libelle"`
	Ice                  string `json:"ice"`
	Rc                   string `json:"rc"`
	FormeJuridique       string `json:"forme_juridique"`
	AffiliationCnss      string `json:"affiliation_cnss" gorm:"size:50"`
}
