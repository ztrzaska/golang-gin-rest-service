package main

import (
	"github.com/gin-gonic/gin"
	"golang-gin-rest-service/api"
)

func main() {
	router := createRouter()
	router.Run("localhost:9001")
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", api.GetBooks)
	router.GET("/books/details", api.GetBookByQueryId)
	router.GET("/books/:id", api.GetBookById)

	router.POST("/books", api.CreateBook)

	return router
}
