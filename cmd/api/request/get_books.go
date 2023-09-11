package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	"golang-gin-rest-service/cmd/mongo"
	"log"
	"net/http"
)

// swagger:model GetBooksResponse
type Book struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	AuthorId      string `json:"author-id"`
	NumberOfPages int    `json:"number-of-pages"`
}

// swagger:route GET /api/v1/books books get_v1_books
//
//	Responses:
//	  200: GetBooksResponse
func GetBooks(c *gin.Context) {
	books, err := findAllBooks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

func findAllBooks() ([]*Book, error) {
	var books []*Book
	cur, err := mongo.Database.Collection(mongo.BooksCollection).Find(context.Background(), bson.D{})

	if err = cur.All(context.Background(), &books); err != nil {
		if err == mongodb.ErrNoDocuments {
			return nil, nil
		}

		log.Fatal("cannot convert books from db", err)
		return nil, err
	}

	return books, nil
}
