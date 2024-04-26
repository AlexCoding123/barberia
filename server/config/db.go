package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := "localhost"
	user := "alexcode123"
	password := "Baseball22"
	dbName := "barbershop"
	port := 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Toronto", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}
