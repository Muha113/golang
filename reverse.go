package main

import (
	"errors"
	"fmt"
)

func reverse(arr []int64) ([]int64, error) {
	if len(arr) == 0 {
		return nil, errors.New("Invalid argument exception")
	}
	newArr := make([]int64, len(arr))
	copy(newArr, arr)
	for i := 0; i < len(newArr)/2; i++ {
		newArr[i], newArr[len(newArr)-i-1] = newArr[len(newArr)-i-1], newArr[i]
	}
	return newArr, nil
}

func printReversed(arr []int64) {
	resultSlice, err := reverse(arr)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Reversed slice:", resultSlice)
		fmt.Println("Original slice:", arr)
	}
}

func main() {
	testSlice := []int64{5, 4, 3, 2, 1}
	printReversed(testSlice)
	testSlice = []int64{1, 2, 3, 4, 5, 6}
	printReversed(testSlice)
	testSlice = []int64{}
	printReversed(testSlice)
	var nilSlice []int64
	printReversed(nilSlice)
}
