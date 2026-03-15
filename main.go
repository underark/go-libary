package main

import "fmt"

type book struct {
	title  string
	author string
}

func main() {
	book := newBook("Fear and Loathing in Las Vegas", "Hunter S Thompson")
	fmt.Println(book)
}

func newBook(title string, author string) book {
	return book{title: title, author: author}
}
