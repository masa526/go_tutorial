package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(s)
}
