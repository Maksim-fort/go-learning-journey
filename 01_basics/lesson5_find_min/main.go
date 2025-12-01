package main

import "fmt"

// findMin находит минимальный элемент в слайсе
func findMin(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	fmt.Println(" День 5: Поиск минимального элемента в слайсе")
	testCases := []struct {
		name  string
		slice []int
	}{
		{"Положительные числа", []int{15, 11, 12, 8, 20}},
		{"Отрицательные числа", []int{-5, -10, -3, -1}},
		{"Смешанные числа", []int{10, -5, 0, 25, -15}},
		{"Один элемент", []int{42}},
		{"Пустой слайс", []int{}},
		{"Повторяющиеся значения", []int{5, 5, 5, 2, 5}},
	}

	for _, tc := range testCases {
		result := findMin(tc.slice)
		fmt.Printf("\nСлайс: %v", tc.slice)
		if len(tc.slice) > 0 {
			fmt.Printf(" → Минимальное: %d", result)
		} else {
			fmt.Printf(" → Пустой слайс, результат: %d", result)
		}
	}

	fmt.Println("\n\n--- Объяснение ---")
	fmt.Println("1. Проверяем не пустой ли слайс")
	fmt.Println("2. Берем первый элемент как начальный минимум")
	fmt.Println("3. Сравниваем с остальными элементами")
	fmt.Println("4. Возвращаем найденный минимум")
}
