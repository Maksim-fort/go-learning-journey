package main

import "fmt"

// findIntersection находит общие элементы в двух слайсах
func findIntersection(a, b []int) []int {
	// Создаем map для элементов первого слайса
	elements := make(map[int]bool)
	for _, num := range a {
		elements[num] = true
	}

	// Используем map для отслеживания уже добавленных элементов
	resultMap := make(map[int]bool)
	var result []int

	// Проверяем элементы второго слайса
	for _, num := range b {
		if elements[num] && !resultMap[num] {
			result = append(result, num)
			resultMap[num] = true
		}
	}
	return result
}

func main() {
	fmt.Println("Поиск пересечения слайсов")

	// Твой пример
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{2, 3, 4, 6, 10}

	fmt.Println("Слайс 1:", slice1)
	fmt.Println("Слайс 2:", slice2)

	result := findIntersection(slice1, slice2)
	fmt.Println("Пересечение:", result)

	// Еще примеры
	fmt.Println("\nДругие примеры:")

	fmt.Println(findIntersection([]int{1, 2, 3}, []int{2, 3, 4}))   // [2 3]
	fmt.Println(findIntersection([]int{10, 20, 30}, []int{20, 30})) // [20 30]
	fmt.Println(findIntersection([]int{1, 2, 3}, []int{4, 5, 6}))   // [] (нет пересечения)

	// Пример с повторяющимися элементами
	fmt.Println(findIntersection(
		[]int{1, 2, 2, 3, 3, 3},
		[]int{2, 2, 3, 4},
	)) // [2 3] (без дубликатов)
}
