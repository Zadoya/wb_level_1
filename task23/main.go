package main

// Удалить i-ый элемент из слайса.

import (
	"fmt"
)

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} 
	i := 5	// i-ый элемент
	
	array = append(array[:i],array[i + 1:]...)
	
	fmt.Println(array)
}