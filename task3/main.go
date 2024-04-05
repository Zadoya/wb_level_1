package main

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их 
// квадратов(22+32+42….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
)

func SquareSum(num int, result *int) {
	*result += num * num
}

func main() {
	var (
		sum int
		wg  sync.WaitGroup
	)

	nums := []int{2, 4, 6, 8, 10}
	for i := range nums {
		wg.Add(1)
		go func(i int) {
			SquareSum(nums[i], &sum)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Сумма кваратов последовательности", nums, "равна", sum)
}
