package main

import (
	"errors"
	"fmt"
	"log"
)

type book struct {
	title  string
	author string
}

func main() {
	log.SetFlags(0)

	book, err := newBook("Book", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(book)
}

func newBook(title string, author string) (book, error) {
	if title == "" {
		return book{}, errors.New("error: field 'title' is empty")
	}

	if author == "" {
		return book{}, errors.New("error: field 'author' is empty")
	}

	return book{title: title, author: author}, nil
}
