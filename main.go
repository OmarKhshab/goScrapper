package main

import (
	"example.com/rest-api/cronjobs"
	"example.com/rest-api/env"
	"example.com/rest-api/middleware"
	"example.com/rest-api/mongoose"
	postgresdb "example.com/rest-api/postgres_db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	env.LoadEnv()
	mongoose.InitMongoose()
	postgresdb.InitPostgres()
}
func main() {
	server := gin.Default()
	server.Use(middleware.LogRequests)
	routes.RegisterRoutes(server)
	cronjobs.GetArticles()
	server.Run(":8080") // localhost:8080
}
