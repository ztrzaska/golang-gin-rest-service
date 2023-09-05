package api

import (
	"github.com/gin-gonic/gin"
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

	books = append(books, book)
	c.IndentedJSON(http.StatusCreated, book)
}
