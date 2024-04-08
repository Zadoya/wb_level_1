package main

// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel
// из переменной типа interface{}

import (
	"fmt"
)

func typeDetector(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("Тип переменной - целое число.")
	case string:
		fmt.Println("Тип переменной - строка.")
	case bool:
		fmt.Println("Это логическая переменная.")
	case chan int:
		fmt.Println("Тип переменной - целочисленный канал")
	case chan string:
		fmt.Println("Тип переменной - строковый канал")
	case chan bool:
		fmt.Println("Тип переменной - логический канал")
	default:
		fmt.Println("Это неизвестный для меня тип переменной.")
	}
}

type myStruct struct{}

func main() {
	var (
		chInt  chan int
		chStr  chan string
		chBool chan bool
	)

	typeDetector(42)
	typeDetector("hello")
	typeDetector(true)
	typeDetector(chInt)
	typeDetector(chStr)
	typeDetector(chBool)
	typeDetector(myStruct{})
}
