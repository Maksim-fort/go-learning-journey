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

// FilterByYear возвращает книги, изданные после указанного года
func (l Library) FilterByYear(year int) []Book {
	var result []Book
	for _, book := range l.Books {
		if book.Year > year {
			result = append(result, book)
		}
	}
	return result
}

func main() {
	library := Library{
		Books: []Book{
			{"Война и мир", "Лев Толстой", 2005},
			{"Преступление и наказание", "Достоевский", 1985},
			{"1984", "Джордж Оруэлл", 1949},
			{"Гарри Поттер", "Дж. К. Роулинг", 1997},
			{"Мастер и Маргарита", "Михаил Булгаков", 1967},
		},
	}

	fmt.Println("Все книги в библиотеке:")
	for i, book := range library.Books {
		fmt.Printf("%d. %s (%d)\n", i+1, book.Title, book.Year)
	}

	fmt.Println("\nКниги после 1996 года:")
	newBooks := library.FilterByYear(1996)
	for i, book := range newBooks {
		fmt.Printf("%d. %s - %s (%d)\n", i+1, book.Title, book.Author, book.Year)
	}

	fmt.Println("\nКниги после 2000 года:")
	veryNewBooks := library.FilterByYear(2000)
	for i, book := range veryNewBooks {
		fmt.Printf("%d. %s (%d)\n", i+1, book.Title, book.Year)
	}

	// Пустая библиотека
	emptyLibrary := Library{}
	fmt.Println("\nПустая библиотека:", emptyLibrary.FilterByYear(1990))
}
