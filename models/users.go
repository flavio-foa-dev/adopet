package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Endereco  UsersAdress
	Pets      []Pet
}

type UsersAdress struct {
	gorm.Model
	Rua        string `json:"rua"`
	Numero     int64  `json:"numero"`
	Referencia string `json:"referencia"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	UF         string `json:"uf"`
}
