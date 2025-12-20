package main

import "fmt"

func removeAtIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	result := removeAtIndex([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println(result)
}
