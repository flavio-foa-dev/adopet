package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Nome       string `json:"nome"`
	Raca       string `json:"raca"`
	Idade      int    `json:"idade"`
	Sexo       string `json:"sexo"`
	Situacao   string `json:"situacao"`
	User_id    Users
	Shelter_id Shelter
}
