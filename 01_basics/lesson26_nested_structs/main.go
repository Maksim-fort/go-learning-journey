package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

type Library struct {
	Books []Book
}

func main() {
	library := Library{
		Books: []Book{
			{"Война и мир", "Лев Толстой", 1869},
			{"Преступление и наказание", "Достоевский", 1866},
			{"1984", "Джордж Оруэлл", 1949},
		},
	}

	fmt.Println("Библиотека:", library)

	fmt.Println("\nКниги в библиотеке:")
	for i, book := range library.Books {
		fmt.Printf("%d. %s - %s (%d)\n", i+1, book.Title, book.Author, book.Year)
	}
	newBook := Book{"Гарри Поттер", "Дж. К. Роулинг", 1997}
	library.Books = append(library.Books, newBook)

	fmt.Println("\nПосле добавления новой книги:")
	fmt.Println("Всего книг:", len(library.Books))

	smallLibrary := Library{
		Books: []Book{
			{"Маленький принц", "Антуан де Сент-Экзюпери", 1943},
		},
	}

	fmt.Println("\nМаленькая библиотека:", smallLibrary)
}
