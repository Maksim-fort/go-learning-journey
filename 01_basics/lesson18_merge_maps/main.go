package main

import "fmt"

// mergeMaps объединяет две map, суммируя значения одинаковых ключей
func mergeMaps(a, b map[string]int) map[string]int {
	result := make(map[string]int)

	// Копируем первую map
	for k, v := range a {
		result[k] = v
	}

	// Добавляем вторую map, суммируя значения
	for k, v := range b {
		result[k] = result[k] + v
	}

	return result
}

func main() {
	fmt.Println(" Объединение map с суммированием")

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}

	fmt.Println("Map 1:", map1)
	fmt.Println("Map 2:", map2)

	merged := mergeMaps(map1, map2)
	fmt.Println("Объединенная map:", merged)

}
