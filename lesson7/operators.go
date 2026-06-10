package main

import (
	"fmt"
)

func main() {
	var a = 15 + 13
	fmt.Println(a)

	var x = 10
	x += 5
	fmt.Println(x)

	var e = 5
	var w = 3
	fmt.Println(e > w)

	var o = 9
	var p = 8

	fmt.Printf("x = %b\n", o)
	fmt.Printf("y = %b\n", p)

	fmt.Printf("x & y is %b\n", o&p)
}
