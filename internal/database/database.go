package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dbAddress := "host=localhost user=postgres password=postgres dbname=go_restaurant port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbSeed(db)
	return db
}
