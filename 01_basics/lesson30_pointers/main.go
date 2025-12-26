package main

import "fmt"

// Double принимает указатель и возвращает удвоенное значение
func Double(n *int) int {
	return *n * 2
}

func main() {
	x := 5

	// &x - получаем адрес переменной x
	result := Double(&x)

	fmt.Println(result) // 10
}
