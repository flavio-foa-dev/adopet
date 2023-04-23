package models

import "gorm.io/gorm"

type Shelter struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Pets    []Pet  // Relacionamento um-para-muitos
}
