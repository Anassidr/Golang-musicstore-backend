package config

import (
	_ "github.com/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	connectionString := "host=localhost port=5432 user=postgres dbname=go-musicstore sslmode=disable"
	d, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
