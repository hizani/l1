// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {

	return &Point{x, y}
}

func (point *Point) GetCoords() (x, y float64) {
	return point.x, point.y
}

func CalcDistance(p1, p2 *Point) float64 {
	pow1 := math.Pow(p2.x-p1.x, 2)
	pow2 := math.Pow(p2.y-p1.y, 2)
	return math.Sqrt(pow1 + pow2)
}

func main() {
	p1 := NewPoint(10.23, 20.11)
	p2 := NewPoint(0, 0)

	x1, y1 := p1.GetCoords()
	x2, y2 := p2.GetCoords()
	fmt.Printf("p1: x=%f\ty=%f\n", x1, y1)
	fmt.Printf("p2: x=%f\ty=%f\n", x2, y2)

	fmt.Println("Distance: ", CalcDistance(p1, p2))
}
