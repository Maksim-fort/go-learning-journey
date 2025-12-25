package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) Describe() {
	fmt.Println("Книга:", b.Title)
	fmt.Println("Автор:", b.Author)
	fmt.Println("Год:", b.Year)
}

func main() {
	book := Book{
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Year:   1869,
	}

	book.Describe()

	// Еще одна книга
	book2 := Book{"Преступление и наказание", "Достоевский", 1866}
	book2.Describe()
}
