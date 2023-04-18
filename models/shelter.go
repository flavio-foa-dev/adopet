package models

import "gorm.io/gorm"

type Shelter struct {
	gorm.Model
	Nome          string `json:"nome"`
	Telefone      string `json:"telefone"`
	Administrador string `json:"administrador"`
	Tipo          string `json:"tipo"`
	Endereco_id   int    `json:"endereco_id"`
	Site          string `json:"site"`
	Rede_social   string `json:"rede_social"`
	Pets          []Pet  `json:"pets"`
}

type ShelterAdress struct {
	gorm.Model
	Rua        string `json:"rua"`
	Numero     int64  `json:"numero"`
	Referencia string `json:"referencia"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	UF         string `json:"uf"`
}
