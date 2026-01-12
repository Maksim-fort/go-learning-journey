package main

import "fmt"

func main() {
	// Создаем указатель через new
	var p *int
	p = new(int)

	// Присваиваем значение через указатель
	*p = 100

	// Выводим значение
	fmt.Println(*p)

	// Еще примеры
	fmt.Println("\n--- Еще примеры ---")

	// Сокращенная запись
	p2 := new(int)
	*p2 = 200
	fmt.Println("*p2 =", *p2)

	// Указатель на строку
	strPtr := new(string)
	*strPtr = "Hello, Go!"
	fmt.Println("*strPtr =", *strPtr)

	// Проверка
	fmt.Println("p == nil?", p == nil)
	fmt.Println("strPtr == nil?", strPtr == nil)
}
