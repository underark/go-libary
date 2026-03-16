package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []book{}

func main() {
	router := gin.Default()
	router.POST("/books", newBookHandler)
	router.Run("localhost:8080")
}

func newBook(title string, author string) (book, error) {
	if title == "" {
		return book{}, errors.New("error: field 'title' is empty")
	}

	if author == "" {
		return book{}, errors.New("error: field 'author' is empty")
	}

	return book{Title: title, Author: author}, nil
}

func newBookHandler(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println("error with BindJSON")
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, books)
}
