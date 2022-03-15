package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDatabase() {

	connectionString := "host=localhost port=5432 user=postgres dbname=books sslmode=disable password=mypostgres"

	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(100)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}
