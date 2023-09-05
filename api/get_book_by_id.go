package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetBookById(c *gin.Context) {
	id := c.Param("id")

	book, err := findBookById(id)
	if err != nil {
		log.Println("no Book found, id: ", id)
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

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
	return nil, errors.New("Book not found")
}
