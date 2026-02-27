package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var Wg sync.WaitGroup

func longOperation(ctx context.Context, id int) {
	defer Wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Операция %d завершена (причина: %v)\n", id, ctx.Err())
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("Операция %d работает...\n", id)
		}
	}
}

func main() {
	// Создаем контекст с таймаутом 3 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Запускаем горутины
	for i := 1; i <= 3; i++ {
		Wg.Add(1)
		id := i // Копируем i для горутины
		go longOperation(ctx, id)
	}

	Wg.Wait()
	fmt.Println("🏁 Программа завершена")
}
