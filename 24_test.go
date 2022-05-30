package l1_test

import (
	"log"
	"math"
	"testing"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func Test_point(t *testing.T) {
	log.Println(
		Distance(
			NewPoint(4, 10),
			NewPoint(4, 5),
		),
	)
}
