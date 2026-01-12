package main

import "fmt"

func areEqual(i, j *int) bool {
	return *i == *j
}

func main() {
	a, b := 10, 10
	fmt.Println(areEqual(&a, &b)) // true

	c := 20
	fmt.Println(areEqual(&a, &c)) // false
}
