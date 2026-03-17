package main

import (
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
	router.POST("/add", newBookHandler)
	router.GET("/books", getBooksHandler)
	router.Run("localhost:8080")
}

func newBook(c *gin.Context) error {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println("error with BindJSON")
		return err
	}

	books = append(books, newBook)
	return nil
}

func newBookHandler(c *gin.Context) {
	if err := newBook(c); err != nil {
		c.Error(err)
	}

	c.Redirect(http.StatusTemporaryRedirect, "/books")
}

func getBooksHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "getBooks.tmpl", gin.H{
		"title": "My Library",
		"books": books,
	})
}
