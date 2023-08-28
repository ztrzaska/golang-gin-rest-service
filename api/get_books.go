package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	AuthorId      string `json:"author-id"`
	NumberOfPages int    `json:"number-of-pages"`
}

var books = []book{
	{Id: "1", Title: "Harry Potter", AuthorId: "1", NumberOfPages: 300},
	{Id: "2", Title: "Rich Dad, Poor Dad", AuthorId: "2", NumberOfPages: 180},
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
