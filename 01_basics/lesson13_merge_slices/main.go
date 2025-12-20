package main

import "fmt"

// mergeSlices объединяет два слайса
func mergeSlices(a, b []int) []int {
	return append(a, b...)
}

func main() {
	// Твой пример
	result := mergeSlices([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println("Результат:", result)

	// Еще примеры
	fmt.Println("Еще примеры:")
	fmt.Println(mergeSlices([]int{10, 20}, []int{30, 40}))
	fmt.Println(mergeSlices([]int{7, 8}, []int{9}))
	fmt.Println(mergeSlices([]int{}, []int{1, 2, 3}))
}
