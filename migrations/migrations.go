package migrations

import (
	"github.com/flavio-foa-dev/adopet/database"
	"github.com/flavio-foa-dev/adopet/models"
)

func Migrate() {
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Shelter{})
	database.DB.AutoMigrate(&models.Pet{})
}
