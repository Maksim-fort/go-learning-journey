package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func createNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

func main() {
	// Создаем два узла
	first := createNode(1)
	second := createNode(2)

	// Связываем их: first -> second
	first.Next = second

	fmt.Println("Первый узел:", first.Value)             // 1
	fmt.Println("Второй узел:", first.Next.Value)        // 2
	fmt.Println("Следующий за вторым:", first.Next.Next) // nil
}
