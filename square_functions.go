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

func (p Sqare) End() (int, int) {
	x := p.start.x + int(p.a)
	y := p.start.y + int(p.a)
	return x, y
}

func (p Sqare) Perimeter() uint {
	perimeter := p.a * 4
	return perimeter
}

func (p Sqare) Area() uint {
	area := p.a * p.a
	return area
}

func main() {
	s := Sqare{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
