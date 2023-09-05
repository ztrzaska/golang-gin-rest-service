package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		return
	}

	books = append(books, book)
	c.IndentedJSON(http.StatusCreated, book)
}
