package postgresdb

import (
	"fmt"
	"log"
	"time"

	"example.com/rest-api/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDb *gorm.DB

func InitPostgres() {
	var err error

	DB_HOST := env.Get("DB_HOST")
	DB_PORT := env.Get("DB_PORT")
	DB_USER := env.Get("DB_USER")
	DB_NAME := env.Get("DB_NAME")
	DB_PASSWORD := env.Get("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)

	PostgresDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	sqlDB, err := PostgresDb.DB()
	if err != nil {
		log.Fatal("Error getting database connection", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Print("Database connection initialized")
}
