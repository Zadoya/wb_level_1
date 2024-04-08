package main

// Поменять местами два числа без создания временной переменной

func main() {
	num1, num2 := 5, 10

	num1, num2 = num2, num1
}