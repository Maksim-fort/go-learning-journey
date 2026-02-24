package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Wg sync.WaitGroup

// worker обрабатывает задания из канала work и отправляет результаты в result
func worker(n int, work <-chan int, result chan<- int) {
	defer Wg.Done()
	for wr := range work {
		delay := time.Duration(rand.Intn(500)) * time.Millisecond
		time.Sleep(delay)
		result <- wr * 2
		fmt.Printf("Worker %d обработал число %d → %d\n", n, wr, wr*2)
	}
}

// printResult выводит результаты из канала
func printResult(result <-chan int) {
	for rs := range result {
		fmt.Println("Результат:", rs)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	work := make(chan int)
	result := make(chan int)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Запускаем пул воркеров (3 штуки)
	for i := 1; i <= 3; i++ {
		Wg.Add(1)
		id := i
		go worker(id, work, result)
	}

	// Отправляем задания в канал work
	Wg.Add(1)
	go func() {
		for _, v := range numbers {
			work <- v
		}
		close(work)
		Wg.Done()
	}()

	// Закрываем result когда все воркеры закончат
	go func() {
		Wg.Wait()
		close(result)
	}()

	// Выводим результаты
	printResult(result)

	fmt.Println("🎉 Работа окончена")
}
