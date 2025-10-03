package main

import "fmt"

func main() {
	numbers := [5]int{2, 3, 5, 10, 15}
	max := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[max] {
			max = i
		}
	}
	fmt.Println("индекс  числа", max)
}
