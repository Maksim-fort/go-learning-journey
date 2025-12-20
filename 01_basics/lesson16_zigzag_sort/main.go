package main

import "fmt"

// merge сливает два отсортированных слайса в один отсортированный
func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	// Сливаем пока есть элементы в обоих слайсах
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// zigzagSort размещает элементы в зигзаг-порядке
func zigzagSort(merged []int) []int {
	result := make([]int, len(merged))
	left, right := 0, len(merged)-1

	for i := 0; i < len(merged); i++ {
		if i%2 == 0 {
			// Четные позиции: минимальные из оставшихся
			result[i] = merged[left]
			left++
		} else {
			// Нечетные позиции: максимальные из оставшихся
			result[i] = merged[right]
			right--
		}
	}
	return result
}

// mergeZigzag объединяет две функции
func mergeZigzag(a, b []int) []int {
	merged := merge(a, b)
	return zigzagSort(merged)
}

func main() {
	// Твой пример
	result := mergeZigzag([]int{1, 2, 3, 4, 6}, []int{5, 6, 7, 9, 10})
	fmt.Println("Результат:", result)

	// Еще примеры
	fmt.Println("\nДругие примеры:")
	fmt.Println(mergeZigzag([]int{1, 3, 5}, []int{2, 4, 6}))
	fmt.Println(mergeZigzag([]int{10, 20, 30}, []int{15, 25, 35}))
}
