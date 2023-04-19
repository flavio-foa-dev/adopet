package main

import (
	"github.com/flavio-foa-dev/adopet/database"
	"github.com/flavio-foa-dev/adopet/migrations"
	"github.com/flavio-foa-dev/adopet/server"
)

func main() {
	database.StartDB()

	migrations.Migrate()

	server := server.NewServer()
	server.Run()
}
