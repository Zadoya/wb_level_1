package main

// К каким негативным последствиям может привести данный фрагмент кода.
// Как это исправить? Приведите корректный пример реализации.

// 	var justString string
// 	func someFunc() {
// 		v := createHugeString(1 << 10)
//  	justString = v[:100]
//	}

//	func main() {
//  	someFunc()
//	}

import (
	"fmt"
	"math/rand"
)
// нет необходимости в объявлении глобальной переменной
// var justString string

func createHugeString(size int) (str string) {
	chars := []rune("ЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЁЯЧСМИТЬБЮйцукенгшщзхъфывапролджэёячсмитьбю")
	
	for i := 0; i < size; i++ {
		str += string(chars[rand.Intn(len(chars))])
	}
	return
}

func someFunc() {
// объявлена функция createHugeString(), но нет её реализации
	v := createHugeString(1 << 10)

// срез происходит по количеству байт, а не символов, если у нас будут символы не из ASCII(т.е. руны),
// то строка будет короче, а так же может быть искажён последний символ
// например, если поставить для русских символов не чётное число байт
	byteString := v[:99]
	runeString := string([]rune(v)[:100])
	
	fmt.Println("по байтам:  ", byteString)
	fmt.Println("по символам:", runeString)
}

func main() {
	someFunc()
}
