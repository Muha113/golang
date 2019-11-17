package main

import (
	"fmt"
)

func factorial(n uint) uint {
	if n > 20 {
		panic("uint overflow")
	}
	var result uint = 1
	for i := uint(1); i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorial(5))  //ok
	fmt.Println(factorial(24)) //panic
}
