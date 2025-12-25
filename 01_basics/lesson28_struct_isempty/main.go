package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

// IsEmpty проверяет, является ли книга пустой
func (b Book) IsEmpty() bool {
	return b.Title == "" && b.Author == "" && b.Year == 0
}

func main() {
	// Твой пример
	book := Book{
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Year:   1869,
	}

	fmt.Println("Книга:", book)
	fmt.Println("Пустая?", book.IsEmpty())

	// Еще примеры
	fmt.Println("\n--- Примеры ---")

	// Полная книга
	fullBook := Book{"Преступление и наказание", "Достоевский", 1866}
	fmt.Printf("%v → Пустая? %v\n", fullBook, fullBook.IsEmpty())

	// Частично заполненная
	partialBook := Book{Title: "1984", Author: "Оруэлл"} // Year = 0
	fmt.Printf("%v → Пустая? %v\n", partialBook, partialBook.IsEmpty())

	// Полностью пустая
	emptyBook := Book{}
	fmt.Printf("%v → Пустая? %v\n", emptyBook, emptyBook.IsEmpty())

	// Только автор
	authorOnly := Book{Author: "Толстой"}
	fmt.Printf("%v → Пустая? %v\n", authorOnly, authorOnly.IsEmpty())

	// Только год
	yearOnly := Book{Year: 2024}
	fmt.Printf("%v → Пустая? %v\n", yearOnly, yearOnly.IsEmpty())

	// Использование в условии
	fmt.Println("\n--- Использование в условии ---")

	books := []Book{
		{"Война и мир", "Толстой", 1869},
		{},
		{"1984", "Оруэлл", 1949},
		Book{},
	}

	fmt.Println("Проверка книг в массиве:")
	for i, b := range books {
		if b.IsEmpty() {
			fmt.Printf("Книга %d: ПУСТАЯ\n", i+1)
		} else {
			fmt.Printf("Книга %d: %s\n", i+1, b.Title)
		}
	}
}
