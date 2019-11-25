package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	fmt.Println(emoji.Sprint(":clap:"), "Hello world", emoji.Sprint(":beer: :grin:"))
}
