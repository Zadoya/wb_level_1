package main

// Дана структура Human (с произвольным набором полей и методов). 
// Реализовать встраивание методов в структуре Action 
// от родительской структуры Human (аналог наследования).

import "fmt"

// В Go наследование реализуется через композиию.

type Human struct {
	Name string
	Age  int
}

func (h *Human) Eat() {
	fmt.Printf("%s eating.\n", h.Name)
}

func (h *Human) Walk() {
	fmt.Printf("%s walking.\n", h.Name)
}


type Action struct {
	Human
} 

func main() {
	act := Action{}

	act.Name = "Ivan"

	act.Eat()
	act.Walk()
}