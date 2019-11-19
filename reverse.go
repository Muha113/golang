package main

import (
	"fmt"
)

func reverse(arr []int64) []int64 {
	if arr == nil || len(arr) == 0 {
		panic("wrong argument")
	}
	newArr := make([]int64, len(arr))
	copy(newArr, arr)
	for i := 0; i < len(newArr)/2; i++ {
		newArr[i], newArr[len(newArr)-i-1] = newArr[len(newArr)-i-1], newArr[i]
	}
	return newArr
}

func main() {
	testSlice := []int64{5, 4, 3, 2, 1}
	fmt.Println("Reversed slice:", reverse(testSlice))
	fmt.Println("Original slice:", testSlice)
	testSlice = []int64{1, 2, 3, 4, 5, 6}
	fmt.Println("Reversed slice:", reverse(testSlice))
	fmt.Println("Original slice:", testSlice)
}
