package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getArticle(context *gin.Context) {
	fmt.Println("articles")

	context.JSON(http.StatusOK, gin.H{"message": "works"})
}

func RegisterRoutes(server *gin.Engine) {
	routes := server.Group("api/v1")
	routes.GET("/articles", getArticles)
	routes.POST("/article", createArticle)
}
