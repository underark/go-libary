# go-libary-demo

A small library demo using Go's Gin framework and server-side template rendering.

# Features

Adding and viewing books in a personal 'library'. This demo app uses an in-memory database of books to store and render the data.

# Learning points

- Using Gin to set up GET and POST routes for a web app.
- Redirecting to another route e.g. POST "/" to GET "/books" to show the user the list of books after adding a book to the list
- Using HTML templates to define a base with different content blocks implemented on each route.
- Attaching static resources like CSS files to the web app.
