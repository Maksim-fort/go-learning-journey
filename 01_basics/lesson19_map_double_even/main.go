package main

import "fmt"

// doubleEven удваивает значения с четными числами в map
func doubleEven(numbers map[int]int) map[int]int {
	result := make(map[int]int)

	for key, value := range numbers {
		if value%2 == 0 {
			result[key] = value * 2
		} else {
			result[key] = value
		}
	}

	return result
}

func main() {
	fmt.Println("Удвоение четных значений в map")

	nums := map[int]int{1: 3, 2: 4, 3: 5}

	fmt.Println("Исходная map:", nums)

	resultNums := doubleEven(nums)
	fmt.Println("После удвоения четных:", resultNums)

	fmt.Println("\nДругие примеры:")

	evenMap := map[int]int{1: 2, 2: 4, 3: 6}
	fmt.Printf("%v → %v\n", evenMap, doubleEven(evenMap))

	oddMap := map[int]int{1: 1, 2: 3, 3: 5}
	fmt.Printf("%v → %v\n", oddMap, doubleEven(oddMap))

	// Пример 3: Смешанные
	mixedMap := map[int]int{
		10: 15, // нечетное - остается
		20: 24, // четное - удваивается
		30: 50, // четное - удваивается
		40: 33, // нечетное - остается
	}
	fmt.Printf("%v → %v\n", mixedMap, doubleEven(mixedMap))

	fmt.Println("\n--- Подробный разбор ---")
	testMap := map[int]int{1: 10, 2: 15, 3: 20, 4: 25}
	fmt.Println("Исходная map:", testMap)

	fmt.Println("\nПоэлементная обработка:")
	for key, value := range testMap {
		if value%2 == 0 {
			fmt.Printf("  Ключ %d: %d (четное) → %d × 2 = %d\n",
				key, value, value, value*2)
		} else {
			fmt.Printf("  Ключ %d: %d (нечетное) → без изменений\n",
				key, value)
		}
	}

	fmt.Println("\nИтог:", doubleEven(testMap))

}
