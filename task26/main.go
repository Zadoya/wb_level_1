package main

// Разработать программу, которая проверяет, что все символы в строке уникальные 
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например: 
// abcd — true
// abCdefAaf — false
// aabcd — false

import (
	"strings"
	"fmt"
)

func checkUnique(str string) bool {
	runeString := []rune(strings.ToLower(str))
	unique := make(map[rune]bool)
	for i := range runeString {
		if _, ok := unique[runeString[i]]; ok {
			return false
		}
		unique[runeString[i]] = true
	}
	return true
}

func main() {
	fmt.Println("abCd -", checkUnique("abCd"))
	fmt.Println("abCdefAaf -", checkUnique("abCdefAaf"))
	fmt.Println("aabcd -", checkUnique("aabcd"))
}