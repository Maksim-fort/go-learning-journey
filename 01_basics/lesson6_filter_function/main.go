package main

import "fmt"

// filter возвращает новый слайс с элементами,
// для которых функция predicate возвращает true
func filter(slice []int, predicate func(int) bool) []int {
	result := []int{}
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println(" День 6: Фильтрация слайса с помощью функций")

	nums := []int{15, 1000, 12, 11, 7, 24, 30, 3}

	// Разные функции-фильтры
	even := func(i int) bool { return i%2 == 0 }         // Четные числа
	odd := func(i int) bool { return i%2 != 0 }          // Нечетные числа
	greaterThan20 := func(i int) bool { return i > 20 }  // Больше 20
	divisibleBy5 := func(i int) bool { return i%5 == 0 } // Делятся на 5

	evenNumbers := filter(nums, even)
	oddNumbers := filter(nums, odd)
	greaterThan20Numbers := filter(nums, greaterThan20)
	divisibleBy5Numbers := filter(nums, divisibleBy5)

	// Выводим результаты
	fmt.Printf("Исходный слайс: %v\n\n", nums)

	fmt.Printf("Четные числа: filter(nums, even)\n")
	fmt.Printf("  %v → %v\n\n", nums, evenNumbers)

	fmt.Printf("Нечетные числа: filter(nums, odd)\n")
	fmt.Printf("  %v → %v\n\n", nums, oddNumbers)

	fmt.Printf("Числа больше 20: filter(nums, greaterThan20)\n")
	fmt.Printf("  %v → %v\n\n", nums, greaterThan20Numbers)

	fmt.Printf("Числа делящиеся на 5: filter(nums, divisibleBy5)\n")
	fmt.Printf("  %v → %v\n\n", nums, divisibleBy5Numbers)

	fmt.Println("--- Объяснение ---")
	fmt.Println("Функция filter принимает:")
	fmt.Println("1. Слайс для фильтрации")
	fmt.Println("2. Функцию-предикат (возвращает true/false)")
	fmt.Println("3. Возвращает новый слайс с элементами, где predicate вернул true")
	fmt.Println("\nЭто пример функции высшего порядка!")
}
