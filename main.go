package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Alchemist", Author: "Paulo Coelho", Quantity: 10},
	{ID: "2", Title: "The Little Prince", Author: "Antoine de Saint-Exupéry", Quantity: 5},
	{ID: "3", Title: "The Book Thief", Author: "Markus Zusak", Quantity: 7},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {

	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
