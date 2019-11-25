package tests

import (
	"math"
	"testing"

	"github.com/Muha113/task2/task"
	"github.com/stretchr/testify/assert"
)

func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func TestSquareArea(t *testing.T) {
	squareTest := task.Square{X: 1, Y: 1, Side: 7}
	assert.Equal(t, squareTest.Area(), float64(49), "Wrong Square Area() return value")
	squareTest.Side = 1
	assert.Equal(t, squareTest.Area(), float64(1), "Wrong Square Area() return value")
	squareTest.Side = 254
	assert.Equal(t, squareTest.Area(), float64(64516), "Wrong Square Area() return value")
}

func TestSquarePerimeter(t *testing.T) {
	squareTest := task.Square{X: 1, Y: 1, Side: 7}
	assert.Equal(t, squareTest.Perimeter(), float64(28), "Wrong Square Perimeter() return value")
	squareTest.Side = 1
	assert.Equal(t, squareTest.Perimeter(), float64(4), "Wrong Square Perimeter() return value")
	squareTest.Side = 174
	assert.Equal(t, squareTest.Perimeter(), float64(696), "Wrong Square Perimeter() return value")
}

func TestCircleArea(t *testing.T) {
	circleTest := task.Circle{X: 1, Y: 1, Radius: 4}
	assert.Equal(t, Round(circleTest.Area(), 6), 50.265482, "Wrong Circle Area() return value")
	circleTest.Radius = 17
	assert.Equal(t, Round(circleTest.Area(), 6), 907.920277, "Wrong Circle Area() return value")
	circleTest.Radius = 1
	assert.Equal(t, Round(circleTest.Area(), 6), 3.141593, "Wrong Circle Area() return value")
}

func TestCirclePerimeter(t *testing.T) {
	circleTest := task.Circle{X: 1, Y: 1, Radius: 4}
	assert.Equal(t, Round(circleTest.Perimeter(), 5), 25.13274, "Wrong Circle Perimeter() return value")
	circleTest.Radius = 17
	assert.Equal(t, Round(circleTest.Perimeter(), 5), 106.81415, "Wrong Circle Perimeter() return value")
	circleTest.Radius = 1
	assert.Equal(t, Round(circleTest.Perimeter(), 5), 6.28319, "Wrong Circle Perimeter() return value")
}
