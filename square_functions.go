package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Sqare struct {
	start Point
	a     uint
}

func (p Sqare) End() Point {
	var point Point
	point.x = p.start.x + int(p.a)
	point.y = p.start.y + int(p.a)
	return point
}

func (p Sqare) Perimeter() uint {
	return p.a * 4
}

func (p Sqare) Area() uint {
	return p.a * p.a
}

func main() {
	s := Sqare{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
