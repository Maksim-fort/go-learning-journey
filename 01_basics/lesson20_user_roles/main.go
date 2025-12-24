package main

import "fmt"

// addAdultUsers добавляет роль "adult" для пользователей старше 18 лет
func addAdultUsers(users map[string]int, roles map[string]string) map[string]string {
	for name, age := range users {
		if age >= 18 {
			roles[name] = "adult"
		}
	}
	return roles
}

func main() {
	fmt.Println(" Система ролей пользователей")

	users := map[string]int{"Alice": 25, "Bob": 17, "Charlie": 30}
	roles := make(map[string]string)

	fmt.Println("Пользователи (возраст):", users)
	fmt.Println("Роли (исходно):", roles)

	addAdultUsers(users, roles)
	fmt.Println("Роли (после обработки):", roles)

	fmt.Println("\n--- Пример 1: Все совершеннолетние ---")
	users1 := map[string]int{"John": 20, "Sarah": 25, "Mike": 18}
	roles1 := make(map[string]string)

	fmt.Println("Пользователи:", users1)
	addAdultUsers(users1, roles1)
	fmt.Println("Роли:", roles1)

	fmt.Println("\n--- Пример 2: Есть несовершеннолетние ---")
	users2 := map[string]int{"Anna": 16, "Tom": 21, "Lisa": 15, "Mark": 19}
	roles2 := map[string]string{"Admin": "superuser"} // уже есть роль

	fmt.Println("Пользователи:", users2)
	fmt.Println("Роли (до):", roles2)
	addAdultUsers(users2, roles2)
	fmt.Println("Роли (после):", roles2)

	fmt.Println("\n--- Пример 3: Подробный разбор ---")
	users3 := map[string]int{
		"Алексей": 22,
		"Мария":   17,
		"Иван":    30,
		"Ольга":   16,
		"Петр":    18,
	}
	roles3 := make(map[string]string)

	fmt.Println("Обработка пользователей:")
	for name, age := range users3 {
		if age >= 18 {
			fmt.Printf("  %s: %d лет → роль 'adult'\n", name, age)
			roles3[name] = "adult"
		} else {
			fmt.Printf("  %s: %d лет → без роли (несовершеннолетний)\n", name, age)
		}
	}

	fmt.Println("\nИтоговые роли:", roles3)
}
