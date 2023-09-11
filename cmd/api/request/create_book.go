package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang-gin-rest-service/cmd/mongo"
	"log"
	"net/http"
)

// swagger:route POST /api/v1/books books post_v1_books
//
//	Responses:
//	  200:
func CreateBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		return
	}

	_, err := mongo.Database.Collection(mongo.BooksCollection).InsertOne(context.Background(), book)
	if err != nil {
		log.Println("error inserting new book", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, book)
}
