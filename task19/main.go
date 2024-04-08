package main

// Разработать программу, которая переворачивает подаваемую на ход строку 
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

import (
	"fmt"
)

func main() {
	str := "главрыба"

	fmt.Printf("%s - %s\n", str, ReverseString(str))
}

func ReverseString(str string) string {
	result := []rune(str)
	for i := 0; i < len(result) / 2; i++ {
		result[i], result[len(result) - 1 - i] = result[len(result) - 1 - i], result[i]
	} 
	return string(result)
}