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

// swagger:route GET /api/v1/books/{bookId} books get_v1_books_book_id
//
//	Parameters:
//	  + name: bookId
//	    in: path
//	Responses:
//	  200:
func GetBookById(c *gin.Context) {
	id := c.Param("id")

	book, err := findBookById(id)
	if err != nil {
		log.Println("no Book found, id: ", id)
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// swagger:route GET /api/v1/books/details books get_v1_books_details
//
//	Responses:
//	  200:
func GetBookByQueryId(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id param"})
		return
	}

	book, err := findBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func findBookById(id string) (*Book, error) {
	var book *Book
	err := mongo.Database.Collection(mongo.BooksCollection).
		FindOne(context.Background(), bson.M{"id": id}).Decode(&book)

	if err != nil {
		if err == mongodb.ErrNoDocuments {
			return nil, nil

		}
		log.Println("cannot fetch books", err)
		return nil, err
	}

	return book, nil
}
