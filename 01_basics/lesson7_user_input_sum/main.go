package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("День 7: Работа с пользовательским вводом")
	fmt.Println("Программа вычисляет сумму элементов массива")
	fmt.Println("===========================================")

	var n int
	fmt.Print("\nВведите количество элементов массива: ")
	fmt.Scan(&n)

	// Проверка корректности ввода
	if n <= 0 {
		fmt.Println("Ошибка: количество элементов должно быть больше 0")
		return
	}

	// Создаем слайс нужного размера
	numbers := make([]int, n)

	fmt.Printf("\nВведите %d элементов массива (после каждого числа нажмите Enter):\n", n)

	// Заполняем массив
	for i := 0; i < n; i++ {
		fmt.Printf("Элемент [%d]: ", i)
		fmt.Scan(&numbers[i])
	}

	// Вычисляем сумму
	sum := 0
	for i := 0; i < n; i++ {
		sum += numbers[i]
	}

	// Выводим результаты
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Результаты:")
	fmt.Printf("Введенный массив: %v\n", numbers)
	fmt.Printf("Количество элементов: %d\n", n)
	fmt.Printf("Сумма элементов: %d\n", sum)
	fmt.Printf("Среднее значение: %.2f\n", float64(sum)/float64(n))

	// Дополнительно: находим min и max
	if n > 0 {
		min := numbers[0]
		max := numbers[0]
		for i := 1; i < n; i++ {
			if numbers[i] < min {
				min = numbers[i]
			}
			if numbers[i] > max {
				max = numbers[i]
			}
		}
		fmt.Printf("Минимальный элемент: %d\n", min)
		fmt.Printf("Максимальный элемент: %d\n", max)
	}

	fmt.Println("\n--- Объяснение ---")
	fmt.Println("1. make([]int, n) - создает слайс заданного размера")
	fmt.Println("2. fmt.Scan(&variable) - читает ввод пользователя")
	fmt.Println("3. Цикл for - для обработки каждого элемента")
	fmt.Println("4. += - оператор присваивания с сложением")
}
