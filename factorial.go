package main

import (
	"errors"
	"fmt"
)

func factorial(n uint) (uint, error) {
	if n > 20 {
		return 0, errors.New("uint overflow")
	}
	var result uint = 1
	for i := uint(1); i <= n; i++ {
		result *= i
	}
	return result, nil
}

func getFactorial(number int) {
	if res, err := factorial(uint(number)); err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(res)
	}
}

func main() {
	getFactorial(10) //ok
	getFactorial(20) //ok
	getFactorial(0)  //ok
	getFactorial(1)  //ok
	getFactorial(-3) //error
	getFactorial(25) //error
}
