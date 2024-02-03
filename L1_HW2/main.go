package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "hello world"
	sentence = strings.ToLower(sentence) // Приведем все буквы к нижнему регистру для учета регистра

	letterCount := make(map[rune]int)
	totalLetters := 0

	// Подсчет количества букв
	for _, char := range sentence {
		if 'a' <= char && char <= 'z' {
			letterCount[char]++
			totalLetters++
		}
	}

	// Вывод результатов
	for letter, count := range letterCount {
		percentage := float64(count) / float64(totalLetters) * 100
		fmt.Printf("%c - %d %.2f%%\n", letter, count, percentage)
	}
}
