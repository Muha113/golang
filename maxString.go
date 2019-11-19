package main

import (
	"fmt"
	"unicode/utf8"
)

func max(arr []string) string {
	if arr == nil || len(arr) == 0 {
		panic("wrong argument")
	}
	maxStr := arr[0]
	for i := 1; i < len(arr); i++ {
		if utf8.RuneCountInString(arr[i]) > utf8.RuneCountInString(maxStr) {
			maxStr = arr[i]
		}
	}
	return maxStr
}

func main() {
	testStrings := []string{"one", "two", "three"}
	fmt.Println(max(testStrings))
	testStrings = []string{"one", "two"}
	fmt.Println(max(testStrings))
}
