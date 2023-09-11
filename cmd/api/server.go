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
	"context"
	"github.com/gin-gonic/gin"
	api "golang-gin-rest-service/cmd/api/request"
	"golang-gin-rest-service/cmd/mongo"
)

func main() {
	router := createRouter()
	db := mongo.Connect()
	mongo.Database = db

	defer db.Client().Disconnect(context.Background())

	router.Run("localhost:9001")
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/books", api.GetBooks)
	router.GET("/api/v1/books/:id", api.GetBookById)
	router.POST("/api/v1/books", api.CreateBook)
	router.GET("/api/v1/books/details", api.GetBookByQueryId)

	router.Static("/swaggerui/", "cmd/api/swaggerui")
	return router
}
