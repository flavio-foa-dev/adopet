package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Nome      string `json:"nome"`
	Raca      string `json:"raca"`
	Tipo_pet  string `json:"tipo_pet"`
	Idade     int    `json:"idade"`
	Sexo      string `json:"sexo"`
	Situacao  string `json:"situacao"`
	UserID    int    `json:"user_id"`
	ShelterID int    `json:"shelter_id"`
}
