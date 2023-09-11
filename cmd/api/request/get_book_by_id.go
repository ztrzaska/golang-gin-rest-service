package api

import (
	"errors"
	"github.com/gin-gonic/gin"
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
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func findBookById(id string) (*Book, error) {
	for _, book := range books {
		if id == book.Id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}
