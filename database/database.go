package database

import (
	"log"
	"time"

	"github.com/carloshdurante/api_golang/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "postgres://root:123456@postgres:5432/api_golang_db?sslmode=disable"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDB() *gorm.DB {
	return db
}
