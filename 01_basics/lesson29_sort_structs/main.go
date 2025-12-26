package main

import (
	"fmt"
	"sort"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

type Library struct {
	Books []Book
}

func (l *Library) SortByYear() {
	sort.Slice(l.Books, func(i, j int) bool {
		return l.Books[i].Year < l.Books[j].Year
	})
}

func main() {
	library := Library{
		Books: []Book{
			{"Война и мир", "Толстой", 1869},
			{"1984", "Оруэлл", 1949},
			{"Преступление и наказание", "Достоевский", 1866},
		},
	}

	fmt.Println("До сортировки:")
	for _, book := range library.Books {
		fmt.Println(book.Title, book.Year)
	}

	library.SortByYear()

	fmt.Println("\nПосле сортировки:")
	for _, book := range library.Books {
		fmt.Println(book.Title, book.Year)
	}
}
