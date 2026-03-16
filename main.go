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

var books = []book{
	{"Frankenstein", "Mary Shelley"},
	{"Fear and Loathing in Las Vegas", "Hunter S Thompson"},
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	router.POST("/books", newBookHandler)
	router.GET("/books", getBooksHandler)
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
	c.IndentedJSON(http.StatusCreated, gin.H{
		"title": "My Library",
		"books": books,
	})
}

func getBooksHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "getBooks.tmpl", gin.H{
		"title": "My Library",
		"books": books,
	})
}
