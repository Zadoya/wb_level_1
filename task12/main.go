package main

// Имеется последовательность строк - (cat, cat, dog, cat, tree) 
// создать для нее собственное множество

import (
	"fmt"
)

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println("Cобственное множество: ", Set(array))
}

func Set(array []string) []string {
	result := make([]string, 0, len(array))
	myMap  := make(map[string]bool)

	for i := range array {
		if _, ok := myMap[array[i]]; !ok {
			myMap[array[i]] = true
			result = append(result, array[i])
		}
	}

	return result
}