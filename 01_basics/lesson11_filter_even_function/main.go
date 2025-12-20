package main

import (
	"fmt"
	"strings"
)

// filterEven возвращает слайс только с четными числами
func filterEven(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// filterOdd возвращает слайс только с нечетными числами
func filterOdd(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if v%2 != 0 {
			result = append(result, v)
		}
	}
	return result
}

// filterByCondition универсальная функция фильтрации
func filterByCondition(slice []int, condition func(int) bool) []int {
	result := []int{}
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// isPositive проверяет положительное ли число
func isPositive(x int) bool {
	return x > 0
}

// isDivisibleBy проверяет делится ли число на divisor
func isDivisibleBy(divisor int) func(int) bool {
	return func(x int) bool {
		return x%divisor == 0
	}
}

func main() {
	fmt.Println(" День 11: Функции фильтрации слайсов")
	fmt.Println("========================================")

	var n int
	fmt.Print("Введите количество элементов: ")
	fmt.Scan(&n)

	// Проверка корректности ввода
	if n <= 0 {
		fmt.Println("Ошибка: количество элементов должно быть больше 0")
		return
	}

	numbers := make([]int, n)

	fmt.Printf("\nВведите %d элементов (после каждого числа нажмите Enter):\n", n)

	// Заполняем слайс
	for i := 0; i < n; i++ {
		fmt.Printf("Элемент [%d]: ", i)
		fmt.Scan(&numbers[i])
	}

	// Применяем разные фильтры
	evenNumbers := filterEven(numbers)
	oddNumbers := filterOdd(numbers)

	// Используем универсальную функцию с разными условиями
	positiveNumbers := filterByCondition(numbers, isPositive)
	divisibleBy3 := filterByCondition(numbers, isDivisibleBy(3))
	divisibleBy5 := filterByCondition(numbers, isDivisibleBy(5))

	// Выводим результаты
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Результаты фильтрации:")
	fmt.Printf("Исходный слайс: %v\n", numbers)
	fmt.Printf("Всего элементов: %d\n", n)

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("Фильтрация по четности:")
	fmt.Printf("Четные числа: %v\n", evenNumbers)
	fmt.Printf("Количество четных: %d (%.1f%%)\n",
		len(evenNumbers), float64(len(evenNumbers))/float64(n)*100)

	fmt.Printf("\nНечетные числа: %v\n", oddNumbers)
	fmt.Printf("Количество нечетных: %d (%.1f%%)\n",
		len(oddNumbers), float64(len(oddNumbers))/float64(n)*100)

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("Другие фильтры:")
	fmt.Printf("Положительные числа (>0): %v\n", positiveNumbers)
	fmt.Printf("Числа делящиеся на 3: %v\n", divisibleBy3)
	fmt.Printf("Числа делящиеся на 5: %v\n", divisibleBy5)

	// Статистика
	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("Статистика по четным числам:")
	if len(evenNumbers) > 0 {
		sumEven := 0
		minEven := evenNumbers[0]
		maxEven := evenNumbers[0]

		for _, v := range evenNumbers {
			sumEven += v
			if v < minEven {
				minEven = v
			}
			if v > maxEven {
				maxEven = v
			}
		}

		fmt.Printf("Сумма четных: %d\n", sumEven)
		fmt.Printf("Минимальное четное: %d\n", minEven)
		fmt.Printf("Максимальное четное: %d\n", maxEven)
		fmt.Printf("Среднее четных: %.2f\n", float64(sumEven)/float64(len(evenNumbers)))
	}

	fmt.Println("\n--- Объяснение ---")
	fmt.Println("1. filterEven - специализированная функция для четных чисел")
	fmt.Println("2. filterByCondition - универсальная функция с условием")
	fmt.Println("3. Замыкание: isDivisibleBy возвращает функцию-условие")
	fmt.Println("4. v%2 == 0 - проверка четности (остаток от деления на 2)")
}
