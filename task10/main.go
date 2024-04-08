package main

// Дана последовательность температурных колебаний: 
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. 
// Объединить данные значения в группы с шагом в 10 градусов. 
// Последовательность в подмножноствах не важна.

import (
	"fmt"
)

func main() {
	temperature := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mergeMap := make(map[int][]float64)

	for i := range temperature {
		mergeMap[int(temperature[i] / 10) * 10] = append(mergeMap[int(temperature[i] / 10) * 10], temperature[i])
	}
	fmt.Println(mergeMap)
}