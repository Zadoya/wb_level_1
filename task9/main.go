package main

// Разработать конвейер чисел. Даны два канала: 
// в первый пишутся числа (x) из массива, 
// во второй — результат операции x*2, 
// после чего данные из второго канала должны выводиться в stdout.

import (
	"fmt"
)

func Square(in, out chan int) {
	num := <-in
	out <- num * num
}

func main() {
	in 	:= make(chan int) 
	out := make(chan int)
	defer close(in)
	defer close(out)

	nums := []int{2, 4, 6, 8, 10}
	for i := range nums {
		go Square(in, out)
		in <- nums[i]
	}
//	for i := range nums {
//		in <- nums[i]
//	}
	for i := range nums {
		fmt.Println("Квадрат числа", nums[i], "равен", <-out)
	}
	
}