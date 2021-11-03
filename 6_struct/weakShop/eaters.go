package main

import (
	"fmt"
	"os"

	"./shop"
)

func main() {
	var name1 string = "OYAKODON"
	if _, err := shop.Eat(name1); err != nil {
		fmt.Fprintf(os.Stderr, "cannot eat: '%s'\n", err) //更新
	}

	var name2 string = ""
	if _, err := shop.Eat(name2); err != nil {
		fmt.Fprintf(os.Stderr, "cannot eat: '%s'\n", err) //更新
	}
}
