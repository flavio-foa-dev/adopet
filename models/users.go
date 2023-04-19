package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Adress    string `json:"address"`
	Interesse string `json:"interesse"`
	Pets      []Pet
}
