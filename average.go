package main

import (
	"fmt"
)

func average(arr [6]int) float64 {
	var summ int
	var avg float64
	for i := 0; i < len(arr); i++ {
		summ += arr[i]
	}
	avg = float64(summ) / float64(len(arr))
	return avg
}

func main() {
	fmt.Println(average([6]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(average([6]int{2, 5, -3, 6, 1, 4}))
}
