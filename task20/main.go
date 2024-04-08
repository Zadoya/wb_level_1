package main

// Разработать программу, которая переворачивает слова в строке. 
// Пример: «snow dog sun — sun dog snow».

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Printf("%s - %s\n", str, ReverseWords(str))
}

func ReverseWords(str  string) string{
	words := strings.Split(str, " ")

	for i := 0; i < len(words) / 2; i++ {
		words[i], words[len(words) - 1 - i] = words[len(words) - 1 - i], words[i]
	} 
	
	return strings.Join(words, " ")
}