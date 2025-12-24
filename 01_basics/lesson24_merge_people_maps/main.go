package main

import "fmt"

func mergeMap(a, b map[string]int) map[string]int {
	result := make(map[string]int)

	for i, v := range a {
		result[i] = v
	}

	for k, v := range b {
		result[k] = result[k] + v
	}

	return result
}

func main() {
	a := map[string]int{
		"Alice": 24,
		"Brain": 12,
		"Alex":  14,
	}

	b := map[string]int{
		"Vasili": 24,
		"Griha":  13,
		"Alice":  2,
	}

	fmt.Println(mergeMap(a, b))
}
