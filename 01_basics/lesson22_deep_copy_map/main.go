package main

import "fmt"

// deepCopy делает глубокую копию map
func deepCopy(original map[string]interface{}) map[string]interface{} {
	copy := make(map[string]interface{})

	for key, value := range original {
		// Если значение - другая map, копируем ее рекурсивно
		if innerMap, ok := value.(map[string]interface{}); ok {
			copy[key] = deepCopy(innerMap)
		} else {
			copy[key] = value
		}
	}

	return copy
}

func main() {
	original := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{"c": 2},
	}

	copied := deepCopy(original)
	fmt.Println(copied) // map[a:1 b:map[c:2]]
}
