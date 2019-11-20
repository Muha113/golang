package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func max(arr []string) (string, error) {
	if len(arr) == 0 {
		return "", errors.New("Invalid argument exception")
	}
	maxStr := arr[0]
	for i := 1; i < len(arr); i++ {
		if utf8.RuneCountInString(arr[i]) > utf8.RuneCountInString(maxStr) {
			maxStr = arr[i]
		}
	}
	return maxStr, nil
}

func printStr(arr []string) {
	resultStr, err := max(arr)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(resultStr)
	}
}

func main() {
	testStrings := []string{"one", "two", "three"}
	printStr(testStrings)
	testStrings = []string{"one", "two"}
	printStr(testStrings)
	testStrings = []string{}
	printStr(testStrings)
}
