package main

import "fmt"

// mergeThree объединяет три слайса в один
func mergeThree(a, b, c []int) []int {
	result := make([]int, 0, len(a)+len(b)+len(c))
	result = append(result, a...)
	result = append(result, b...)
	result = append(result, c...)
	return result
}

func main() {
	// Пример из твоего кода
	slice1 := []int{1, 4, 7}
	slice2 := []int{2, 5, 8}
	slice3 := []int{3, 6, 9, 10}

	result := mergeThree(slice1, slice2, slice3)
	fmt.Println(result)

	// Еще примеры
	fmt.Println(mergeThree([]int{1, 2}, []int{3}, []int{4, 5, 6}))
	fmt.Println(mergeThree([]int{}, []int{1}, []int{2, 3}))
	fmt.Println(mergeThree([]int{10}, []int{20}, []int{30}))
}
