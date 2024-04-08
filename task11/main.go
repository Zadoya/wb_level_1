package main

// Реализовать пересечение двух неупорядоченных множеств

import (
	"fmt"
)

func main() {
	array1 := []int{1, 10, 15, 112, 4, 8, 30, 2, 19, 15}
	array2 := []int{4, 2, 45, 18, 8, 10, 111, 33, 17, 19}
	
	fmt.Println("Пересечение множеств:", Intersection(array1, array2))
}

func Intersection(arr1, arr2 []int) []int{
	result := make([]int, 0, len(arr1))
	myMap  := make(map[int]bool)

	for i := range arr1 {
		myMap[arr1[i]] = true
	}
	for i := range arr2 {
		if _, ok := myMap[arr2[i]]; ok {
			result = append(result, arr2[i])
		}
	}

	return result
}