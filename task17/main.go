package main

// Реализовать бинарный поиск встроенными методами языка.

import (
	"fmt"
	"sort"
)

func main() {
	searchVal := 5
	array := []int{9, 7, 8, 6, 4, 5, 1, 3, 2}
	
	sort.Ints(array)
	index := sort.SearchInts(array, searchVal)
	
	fmt.Printf("Значеие %d находитя на позиции [%d] в array\n", searchVal, index)
}