package main

import (
	"fmt"
	"time"
)

func main() {
	// Запускаем горутину
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина:", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	// Основная программа
	for j := 6; j < 11; j++ {
		fmt.Println("Основная:", j)
		time.Sleep(400 * time.Millisecond)
	}

	// Ждем завершения горутины
	time.Sleep(2 * time.Second)
	fmt.Println("Конец")
}
