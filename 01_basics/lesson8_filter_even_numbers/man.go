package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("День 8: Фильтрация четных чисел из массива")
	fmt.Println("==============================================")

	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scan(&n)

	// Проверка корректности ввода
	if n <= 0 {
		fmt.Println("Ошибка: количество элементов должно быть больше 0")
		return
	}

	numbers := make([]int, n)

	fmt.Printf("\nВведите %d элементов массива (после каждого числа нажмите Enter):\n", n)

	// Заполняем массив
	for i := 0; i < n; i++ {
		fmt.Printf("Элемент [%d]: ", i)
		fmt.Scan(&numbers[i])
	}

	// Фильтруем четные числа
	var evenNumbers []int
	var oddNumbers []int

	for _, v := range numbers {
		if v%2 == 0 {
			evenNumbers = append(evenNumbers, v)
		} else {
			oddNumbers = append(oddNumbers, v)
		}
	}

	// Выводим результаты
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Результаты анализа массива:")
	fmt.Printf("Исходный массив: %v\n", numbers)
	fmt.Printf("Всего элементов: %d\n", n)

	fmt.Printf("\nЧетные числа: %v\n", evenNumbers)
	fmt.Printf("Количество четных чисел: %d\n", len(evenNumbers))

	fmt.Printf("\nНечетные числа: %v\n", oddNumbers)
	fmt.Printf("Количество нечетных чисел: %d\n", len(oddNumbers))

	// Дополнительная статистика
	if len(evenNumbers) > 0 {
		sumEven := 0
		for _, v := range evenNumbers {
			sumEven += v
		}
		fmt.Printf("\nСумма четных чисел: %d\n", sumEven)
		fmt.Printf("Среднее четных чисел: %.2f\n", float64(sumEven)/float64(len(evenNumbers)))
	}

	if len(oddNumbers) > 0 {
		sumOdd := 0
		for _, v := range oddNumbers {
			sumOdd += v
		}
		fmt.Printf("\nСумма нечетных чисел: %d\n", sumOdd)
		fmt.Printf("Среднее нечетных чисел: %.2f\n", float64(sumOdd)/float64(len(oddNumbers)))
	}

	fmt.Println("\n--- Объяснение ---")
	fmt.Println("1. v%2 == 0 - проверка на четность (остаток от деления на 2 равен 0)")
	fmt.Println("2. append(nums, v) - добавляет элемент в слайс")
	fmt.Println("3. len(slice) - возвращает длину слайса")
	fmt.Println("4. Итерируем for _, v - пропускаем индекс, берем только значение")
}
