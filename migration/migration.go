package main

import (
	"log"

	"example.com/rest-api/env"
	"example.com/rest-api/models"
	"example.com/rest-api/mongoose"
	postgresdb "example.com/rest-api/postgres_db"
)

func init() {
	env.LoadEnv()
	mongoose.InitMongoose()
	postgresdb.InitPostgres()
}
func main() {
	err := postgresdb.PostgresDb.AutoMigrate(&models.Article{})
	if err != nil {
		log.Fatal("Error migrating to database", err)
	}
	log.Print("Migrated Succesfully")
}
