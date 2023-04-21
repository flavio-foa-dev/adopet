package main

import (
	"fmt"

	"github.com/flavio-foa-dev/adopet/database"
	"github.com/flavio-foa-dev/adopet/migrations"
	"github.com/flavio-foa-dev/adopet/routes"
)

func main() {
	fmt.Println("Starting database...")
	database.StartDB()
	fmt.Println("Starting migration...")
	migrations.Migrate()
	routes.ConfigureRoutes()
}
