package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Birthday увеличивает возраст на 1 (использует pointer receiver)
func (p *Person) Birthday() {
	p.Age = p.Age + 1
}

// GetName возвращает имя (использует value receiver)
func (p Person) GetName() string {
	return p.Name
}

func main() {
	// Создаем через указатель
	p := &Person{Name: "Alice", Age: 25}

	fmt.Println("До дня рождения:")
	fmt.Println("Имя:", p.GetName())
	fmt.Println("Возраст:", p.Age)

	p.Birthday()

	fmt.Println("\nПосле дня рождения:")
	fmt.Println("Возраст:", p.Age)

	// Еще примеры
	fmt.Println("\n--- Еще примеры ---")

	// Value (не указатель) тоже работает!
	p2 := Person{Name: "Bob", Age: 30}
	fmt.Println("\nДо:", p2.Age)
	p2.Birthday() // Go автоматически преобразует (&p2).Birthday()
	fmt.Println("После:", p2.Age)

	// Несколько дней рождения
	p3 := &Person{Name: "Charlie", Age: 40}
	fmt.Println("\nИзначально:", p3.Age)

	for i := 0; i < 5; i++ {
		p3.Birthday()
	}

	fmt.Println("После 5 дней рождений:", p3.Age)

	// Сравнение value и pointer receivers
	fmt.Println("\n--- Value vs Pointer Receivers ---")

	person := Person{Name: "David", Age: 50}

	// Value receiver не меняет оригинал
	tryChangeValue := func(p Person) {
		p.Age = 999
	}

	tryChangeValue(person)
	fmt.Println("После value receiver:", person.Age) // Осталось 50

	// Pointer receiver меняет оригинал
	tryChangePointer := func(p *Person) {
		p.Age = 999
	}

	tryChangePointer(&person)
	fmt.Println("После pointer receiver:", person.Age) // Стало 999
}
