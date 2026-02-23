package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name string
	Time time.Duration
}

var Wg sync.WaitGroup

func main() {
	var tasks []Task
	tasks = []Task{
		{Name: "Скачать данные", Time: 500 * time.Millisecond},
		{Name: "Обработать изображение", Time: 300 * time.Millisecond},
		{Name: "Отправить email", Time: 200 * time.Millisecond},
		{Name: "Обновить базу данных", Time: 400 * time.Millisecond},
		{Name: "Сгенерировать отчет", Time: 600 * time.Millisecond},
		{Name: "Проверить права доступа", Time: 150 * time.Millisecond},
		{Name: "Архивировать файлы", Time: 350 * time.Millisecond},
		{Name: "Запустить тесты", Time: 450 * time.Millisecond},
	}

	for _, task := range tasks {
		Wg.Add(1)
		go func(t Task) {
			defer Wg.Done()
			fmt.Printf("Задача: %s\n", t.Name)
			time.Sleep(t.Time)
			fmt.Printf("Задача: %s завершена\n", t.Name)
		}(task)
	}

	Wg.Wait()
	fmt.Println("Все задачи выполнены")
}
