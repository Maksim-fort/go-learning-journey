package main

import "fmt"

// intersect находит общие ключи в двух map и складывает их значения
func intersect(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)

	// Проверяем ключи из первой map
	for key, value1 := range m1 {
		// Если ключ есть во второй map
		if value2, ok := m2[key]; ok {
			// Складываем значения
			result[key] = value1 + value2
		}
	}

	return result
}

func main() {
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}

	result := intersect(map1, map2)
	fmt.Println(result) // map[b:5]
}
