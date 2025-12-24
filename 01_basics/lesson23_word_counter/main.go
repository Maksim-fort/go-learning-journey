package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordMap(s string) map[string]int {
	word := strings.Fields(s)
	result := make(map[string]int)
	for _, v := range word {
		result[v]++
	}
	return result
}

func main() {
	fmt.Println("Введите набор слов")
	words := bufio.NewScanner(os.Stdin)
	words.Scan()
	input := words.Text()
	fmt.Println(wordMap(input))
}
