package main

import (
	"fmt"
	"strings"
)

// sumSlice вычисляет сумму всех элементов слайса
func sumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// sumSliceRange вычисляет сумму элементов в заданном диапазоне [start, end)
func sumSliceRange(slice []int, start, end int) int {
	if start < 0 || end > len(slice) || start >= end {
		return 0
	}
	sum := 0
	for i := start; i < end; i++ {
		sum += slice[i]
	}
	return sum
}

// averageSlice вычисляет среднее значение элементов слайса
func averageSlice(slice []int) float64 {
	if len(slice) == 0 {
		return 0
	}
	return float64(sumSlice(slice)) / float64(len(slice))
}

func main() {
	fmt.Println(" День 10: Функции для работы со слайсами - 10 дней Go!")
	fmt.Println("=========================================================")

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

	// Вычисляем с помощью разных функций
	totalSum := sumSlice(numbers)
	average := averageSlice(numbers)

	// Суммы частей слайса
	firstHalfSum := sumSliceRange(numbers, 0, n/2)
	secondHalfSum := sumSliceRange(numbers, n/2, n)

	// Суммы четных и нечетных позиций
	var evenPositions []int
	var oddPositions []int
	for i, v := range numbers {
		if i%2 == 0 {
			evenPositions = append(evenPositions, v)
		} else {
			oddPositions = append(oddPositions, v)
		}
	}
	evenPosSum := sumSlice(evenPositions)
	oddPosSum := sumSlice(oddPositions)

	// Выводим результаты
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Результаты анализа слайса:")
	fmt.Printf("Исходный слайс: %v\n", numbers)
	fmt.Printf("Количество элементов: %d\n", n)

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("Основные расчеты:")
	fmt.Printf("Сумма всех элементов: %d\n", totalSum)
	fmt.Printf("Среднее значение: %.2f\n", average)

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("Суммы по частям:")
	fmt.Printf("Первая половина [0-%d]: %v\n", n/2-1, numbers[:n/2])
	fmt.Printf("Сумма первой половины: %d\n", firstHalfSum)

	fmt.Printf("\nВторая половина [%d-%d]: %v\n", n/2, n-1, numbers[n/2:])
	fmt.Printf("Сумма второй половины: %d\n", secondHalfSum)

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("Суммы по позициям:")
	fmt.Printf("Элементы на четных позициях: %v\n", evenPositions)
	fmt.Printf("Сумма на четных позициях: %d\n", evenPosSum)

	fmt.Printf("\nЭлементы на нечетных позициях: %v\n", oddPositions)
	fmt.Printf("Сумма на нечетных позициях: %d\n", oddPosSum)

	// Юбилейное сообщение!
	fmt.Println("\n" + strings.Repeat("🎉", 25))
	fmt.Println("ПОЗДРАВЛЯЮ! Ты изучаешь Go уже 10 дней подряд!")
	fmt.Println("Твой прогресс впечатляет! Продолжай в том же духе! 🚀")
	fmt.Println(strings.Repeat("🎉", 25))
}
