package main

import "fmt"

// reverseSlice переворачивает слайс
func reverseSlice(slice []int) []int {
	// Меняем первый с последним, второй с предпоследним и т.д.
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func main() {
	fmt.Println("Реверс слайса")
	fmt.Println("================")

	var n int
	fmt.Print("Сколько чисел? ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Нужно хотя бы 1 число!")
		return
	}

	// Создаем слайс
	numbers := make([]int, n)

	fmt.Printf("Введите %d чисел (каждое с новой строки):\n", n)

	// Читаем числа
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Printf("\nБыло: %v\n", numbers)

	// Переворачиваем
	reversed := reverseSlice(numbers)

	fmt.Printf("Стало: %v\n", reversed)

	// Простое объяснение
	fmt.Println("\nКак это работает:")
	fmt.Println("1. i идет с начала, j с конца")
	fmt.Println("2. Меняем местами numbers[i] и numbers[j]")
	fmt.Println("3. i увеличиваем, j уменьшаем")
	fmt.Println("4. Повторяем пока i < j")
}
