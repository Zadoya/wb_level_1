package main

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Distance struct {
	Point
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (d *Distance) MeasureDistance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x - p2.x, 2) + math.Pow(p1.y - p2.y, 2))
}

func main() {
	p1 := NewPoint(-1, 3)
	p2 := NewPoint(6, 2)

	distance := Distance{}
	fmt.Printf("Результат измерения: %.3f\n", distance.MeasureDistance(p1, p2))
}
