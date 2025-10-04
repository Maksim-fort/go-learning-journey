package main

import "fmt"

// IsDivision возвращает результат деления и остаток от деления двух чисел
func IsDivision(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	fmt.Println("🎯 День 2: Функции с возвратом нескольких значений")

	quotient, remainder := IsDivision(55, 7)
	fmt.Printf("55 разделить на 7:\n")
	fmt.Printf("Частное: %d, Остаток: %d\n", quotient, remainder)

	fmt.Println("\n--- Другие примеры ---")
	q1, r1 := IsDivision(20, 3)
	fmt.Printf("20 / 3 = %d (остаток %d)\n", q1, r1)

	q2, r2 := IsDivision(100, 25)
	fmt.Printf("100 / 25 = %d (остаток %d)\n", q2, r2)
}
