package main

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/multitemplate"
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
	router.HTMLRender = createRenderer()
	router.GET("/", indexHandler)
	router.POST("/add", newBookHandler)
	router.GET("/books", getBooksHandler)
	router.Run("localhost:8080")
}

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.tmpl", "templates/pages/index.tmpl")
	r.AddFromFiles("books", "templates/base.tmpl", "templates/pages/getBooks.tmpl")
	return r
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Home",
	})
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
	title := c.PostForm("title")
	author := c.PostForm("author")
	if book, err := newBook(title, author); err != nil {
		c.Error(err)
	} else {
		books = append(books, book)
	}

	c.Redirect(http.StatusFound, "/books")
}

func getBooksHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "books", gin.H{
		"title": "My Library",
		"books": books,
	})
}
