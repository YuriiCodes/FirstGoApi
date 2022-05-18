package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbUrl string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
