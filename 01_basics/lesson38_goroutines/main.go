package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Вторая горутина окончена")
	}()

	for j := 6; j < 12; j++ {
		fmt.Println(j)
		time.Sleep(400 * time.Millisecond)
	}
	fmt.Println("Основная герутина окончена")
	time.Sleep(2 * time.Second)
	fmt.Println("Программа завершена")
}
