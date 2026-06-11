package main

import "fmt"

func myMessage() {
	fmt.Println("I just got executed!")
}

func name(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "...!")
}

func myFunction(x int, y int) int {
	return x + y
}

// func myFunction(x int, y int) (result int) {
// 	result = x + y
// 	return
// }

func myFunc(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

// recursion
func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	// myMessage()
	// name("Ali", 3)
	// fmt.Println(myFunction(1, 2))
	// _, b := myFunc(5, "Hello")
	// fmt.Println(b, "\n")
	// testcount(1)
	fmt.Println(factorial_recursion(4))
}
