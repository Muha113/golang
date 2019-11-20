package main

import (
	"fmt"
	"sort"
)

func printSorted(dict map[int]string) {
	keys := make([]int, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Sort(sort.IntSlice(keys))
	for _, a := range keys {
		fmt.Print(dict[a], " ")
	}
	fmt.Print("\n")
}

func main() {
	testMap := map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(testMap)
	testMap = map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(testMap)
	testMap = map[int]string{}
	printSorted(testMap)
	var nilMap map[int]string
	printSorted(nilMap)
}
