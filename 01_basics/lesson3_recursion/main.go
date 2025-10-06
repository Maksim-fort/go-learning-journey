package main

import "fmt"

// recursiveFactorial вычисляет факториал числа с помощью рекурсии
// Факториал n! = n * (n-1) * (n-2) * ... * 1
func recursiveFactorial(n int) int {
	// Базовый случай рекурсии
	if n == 0 {
		return 1
	}
	// Рекурсивный вызов
	return n * recursiveFactorial(n-1)
}

func main() {
	fmt.Println("День 3: Рекурсия в Go")
	fmt.Println("Вычисление факториалов:")

	// Тестируем на разных числах
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7}

	for _, num := range numbers {
		result := recursiveFactorial(num)
		fmt.Printf("%d! = %d\n", num, result)
	}

	fmt.Println("Рекурсия - это когда функция вызывает саму себя.")
	fmt.Println("Для 6! вычисление идет:")
	fmt.Println("6! = 6 * 5!")
	fmt.Println("5! = 5 * 4!")
	fmt.Println("... и так до 0! = 1")
}
