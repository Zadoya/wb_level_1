package main

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20

import (
	"fmt"
	"math/big"
)

// в целом числа до 2^31 должны помещаться в рамках int64,
// или можно воспользоваться встроенной библиотекой использовать
// неограниченно большие числа

func calculator(aStr, bStr string) {
	a, ok := big.NewInt(0).SetString(aStr, 10)
	if !ok {
		fmt.Println("ошибка в создании первого числа")
		return
	}

	b, ok := big.NewInt(0).SetString(bStr, 10)
	if !ok {
		fmt.Println("ошибка в создании второго числа")
		return
	}

	result := big.NewInt(0)

	fmt.Println("Результат сложения:  ", result.Add(a, b))
	fmt.Println("Результат вычитания: ", result.Sub(a, b))
	fmt.Println("Результат множения:  ", result.Mul(a, b))
	fmt.Println("Результат деления:   ", result.Div(a, b))
}

func main() {
	a := "34545465465475687465756436345634636453654363563456363456"
	b := "876545354353536534635464564566453463453"
	calculator(a, b)
}
