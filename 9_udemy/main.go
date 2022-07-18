package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)

	fmt.Printf("%T %v %f \n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d \n", yy, yy, yy)

	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v %d \n", i, i, i)

	var ss string = "14"
	ii, _ := strconv.Atoi(ss)
	fmt.Printf("%T %v %d \n", ii, ii, ii)

	h := "hello world"
	fmt.Println(string(h[0]))

}
