package middleware

import (
	"context"
	"log"

	"example.com/rest-api/mongoose"
	"github.com/gin-gonic/gin"
)

func LogRequests(c *gin.Context) {
	c.Next()
	var urls = c.Request.RequestURI
	var cMethod = c.Request.Method
	var cStatus = c.Writer.Status()

	var db = mongoose.MongoDB
	var collection = db.Collection("logs")
	_, err := collection.InsertOne(context.TODO(), map[string]any{
		"url":    urls,
		"method": cMethod,
		"status": cStatus,
	})
	if err != nil {
		log.Fatal("mongo db error")
	}
}
