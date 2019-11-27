package main

import (
	"fmt"

	"github.com/Muha113/StringToInt/task"
)

func main() {
	integer, err := task.MyStrToInt1("123456")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(integer)
	integer, err = task.MyStrToInt2("+123.000")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(integer)
}
