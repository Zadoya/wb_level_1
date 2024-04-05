package main

// Дана переменная int64. Разработать программу,
// которая устанавливает i-й бит в 1 или 0

import (
	"fmt"
)

func main() {
	var (
		num int64 = 4567
		i   int   = 3
	)
	if i >= 0 && i < 64 {
		fmt.Printf("%064b", uint64(num)^1<<i)
	}
}
