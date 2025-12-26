package main

import "fmt"

// Swap меняет местами значения двух переменных через указатели
func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	c := 100
	d := 50

	fmt.Println("До Swap:")
	fmt.Println("c =", c)
	fmt.Println("d =", d)

	Swap(&c, &d)

	fmt.Println("\nПосле Swap:")
	fmt.Println("c =", c)
	fmt.Println("d =", d)

	// Еще примеры
	fmt.Println("\n--- Еще примеры ---")

	x := 10
	y := 20

	fmt.Printf("До: x=%d, y=%d\n", x, y)
	Swap(&x, &y)
	fmt.Printf("После: x=%d, y=%d\n", x, y)
	a := 1
	b := 2

	fmt.Printf("\na=%d, b=%d\n", a, b)
	Swap(&a, &b)
	fmt.Printf("После Swap: a=%d, b=%d\n", a, b)
}
