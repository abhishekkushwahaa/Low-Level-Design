// Aggregation is a special form of association

package main

import "fmt"

type Book struct {
	title string
}

type Library struct {
	name  string
	books []Book // Aggregation
}

func main() {
	book1 := Book{title: "Go programming"}
	book2 := Book{title: "Js programming"}
	library := Library{name: "Central Library", books: []Book{book1, book2}}

	fmt.Println(library.name, "has these books:")

	for _, book := range library.books {
		fmt.Println(book.title)
	}
}

// Key Takeaways = The library has books, but the books are independent.
