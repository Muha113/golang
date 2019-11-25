package task

import (
	"fmt"
	"math"
)

type Figure interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	X, Y int
	Side uint
}

type Circle struct {
	X, Y   int
	Radius uint
}

func (s Square) Area() float64 {
	return float64(s.Side * s.Side)
}

func (s Square) Perimeter() float64 {
	return float64(s.Side * 4)
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius*c.Radius)
}

func (c Circle) Perimeter() float64 {
	return math.Pi * float64(2*c.Radius)
}

func Task() {
	var s Figure = Square{1, 1, 5}
	var c Figure = Circle{1, 1, 7}

	fmt.Println(s.Area(), s.Perimeter())
	fmt.Println(c.Area(), c.Perimeter())
}
