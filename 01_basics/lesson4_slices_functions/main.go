package main

import "fmt"

// applyToSlice применяет функцию к каждому элементу слайса
// и возвращает новый слайс с результатами
func applyToSlice(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func main() {
	fmt.Println(" День 4: Слайсы и функции высшего порядка")

	numbers := []int{15, 20, 40}

	// Определяем функции для преобразования
	double := func(x int) int { return x * 2 }
	square := func(x int) int { return x * x }
	addTen := func(x int) int { return x + 10 }

	// Применяем разные функции к слайсу
	doubled := applyToSlice(numbers, double)
	squared := applyToSlice(numbers, square)
	addedTen := applyToSlice(numbers, addTen)

	fmt.Printf("Исходный слайс: %v\n", numbers)
	fmt.Printf("Умножение на 2: %v\n", doubled)
	fmt.Printf("Возведение в квадрат: %v\n", squared)
	fmt.Printf("Прибавление 10: %v\n", addedTen)

	fmt.Println("\n--- Объяснение ---")
	fmt.Println("Функции высшего порядка - это функции,")
	fmt.Println("которые принимают или возвращают другие функции.")
	fmt.Println("Здесь applyToSlice применяет переданную функцию")
	fmt.Println("к каждому элементу слайса.")
}
