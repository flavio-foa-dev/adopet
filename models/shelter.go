package models

import "gorm.io/gorm"

type Shelter struct {
	gorm.Model
	Nome          string `json:"nome"`
	Telefone      string `json:"telefone"`
	Endereco      string `json:"endereco"`
	Administrador string `json:"administrador"`
	Tipo_abrigo   string `json:"tipo_abrigo"`
	Site          string `json:"site"`
	Rede_social   string `json:"rede_social"`
	Pets          []Pet
}
