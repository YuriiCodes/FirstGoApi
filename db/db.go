package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUrl := "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
