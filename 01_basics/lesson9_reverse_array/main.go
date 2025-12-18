package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("День 9: Реверс массива (обратный порядок)")
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

	// Способ 1: Создаем новый слайс в обратном порядке
	var reversed1 []int
	for i := len(numbers) - 1; i >= 0; i-- {
		reversed1 = append(reversed1, numbers[i])
	}

	// Способ 2: Реверс на месте (меняем элементы местами)
	reversed2 := make([]int, len(numbers))
	copy(reversed2, numbers)

	for i, j := 0, len(reversed2)-1; i < j; i, j = i+1, j-1 {
		reversed2[i], reversed2[j] = reversed2[j], reversed2[i]
	}

	// Способ 3: Используем два указателя
	reversed3 := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		reversed3[len(numbers)-1-i] = numbers[i]
	}

	// Выводим результаты
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Результаты реверса массива:")
	fmt.Printf("Исходный массив:    %v\n", numbers)
	fmt.Printf("Длина массива:      %d\n", len(numbers))

	fmt.Printf("\nСпособ 1 (append с конца):\n")
	fmt.Printf("  Реверс:           %v\n", reversed1)

	fmt.Printf("\nСпособ 2 (реверс на месте):\n")
	fmt.Printf("  Реверс:           %v\n", reversed2)

	fmt.Printf("\nСпособ 3 (два указателя):\n")
	fmt.Printf("  Реверс:           %v\n", reversed3)

	// Проверка что все способы дают одинаковый результат
	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("Проверка:")
	sameResult := true
	for i := 0; i < len(numbers); i++ {
		if reversed1[i] != reversed2[i] || reversed2[i] != reversed3[i] {
			sameResult = false
			break
		}
	}

	if sameResult {
		fmt.Println("✓ Все три способа дали одинаковый результат!")
	} else {
		fmt.Println("✗ Результаты различаются!")
	}

	fmt.Println("\n--- Объяснение ---")
	fmt.Println("Способ 1: Идем с конца массива и добавляем в новый слайс")
	fmt.Println("Способ 2: Меняем элементы местами (первый с последним и т.д.)")
	fmt.Println("Способ 3: Заполняем новый массив с противоположных концов")
	fmt.Println("\nlen(numbers) - 1 = индекс последнего элемента")
}
