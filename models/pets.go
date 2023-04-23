package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name      string `json:"name"`
	Specie    string `json:"specie"`
	UserID    int    `json:"user_id"`
	ShelterID int    `json:"shelter_id"`
	Shelter   Shelter
	User      User
}
