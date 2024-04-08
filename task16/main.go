package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{9, 7, 8, 6, 4, 5, 1, 3, 2}
	sort.Ints(array)
	fmt.Println(array, sort.IntsAreSorted(array))
}