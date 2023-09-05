// Package classification REST API with go-swagger.
//
//	BasePath: /
//	Version: 1.0.0
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"github.com/gin-gonic/gin"
	api "golang-gin-rest-service/cmd/api/request"
)

func main() {
	router := createRouter()
	router.Run("localhost:9001")
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/books", api.GetBooks)
	router.GET("/api/v1/books/details", api.GetBookByQueryId)
	router.GET("/api/v1/books/:id", api.GetBookById)

	router.POST("/api/v1/books", api.CreateBook)

	router.Static("/swaggerui/", "cmd/api/swaggerui")
	return router
}