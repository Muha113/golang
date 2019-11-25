package tests

import (
	"testing"
	"time"

	"github.com/Muha113/task1/task"
	"github.com/stretchr/testify/assert"
)

func TestLess(t *testing.T) {
	const timeLayout = "2006-Jan-02"

	a, _ := time.Parse(timeLayout, "2005-Aug-15")
	b, _ := time.Parse(timeLayout, "2006-Aug-15")
	c, _ := time.Parse(timeLayout, "2007-Aug-15")
	p := task.People{
		{FirstName: "Vlad", LastName: "Petrov", BirthDay: c},
		{FirstName: "Sasha", LastName: "Polyakov", BirthDay: a},
		{FirstName: "Sasha", LastName: "Polynkov", BirthDay: a},
		{FirstName: "Andrej", LastName: "Nehay", BirthDay: b},
	}
	assert.Equal(t, p.Less(0, 1), true, "Test failed")
	assert.Equal(t, p.Less(0, 0), false, "Test failed")
	assert.Equal(t, p.Less(1, 2), true, "Test failed")
}
