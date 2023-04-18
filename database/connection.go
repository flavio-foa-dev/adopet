package database

import "gorm.io/gorm"

var (
	db  *gorm.DB
	dsn string
)

func connections() {
	dsn = "host=localhost user=root password=docker dbname=alunos port=5432 sslmode=disable"
}
