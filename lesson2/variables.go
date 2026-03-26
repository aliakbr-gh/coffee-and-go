// In Go, there are different types of variables, for example:

// int- stores integers (whole numbers), such as 123 or -123
// float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
// string - stores text, such as "Hello World". String values are surrounded by double quotes
// bool- stores values with two states: true or false

// var
// Can be used inside and outside of functions
// Can only be used inside functions

// :=
// Variable declaration and value assignment can be done separately
// Variable declaration and value assignment cannot be done separately (must be done in the same line)

package main

import (
	"fmt"
)

var x int
var y int = 2
var z = 3

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)

  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)

  var player1 string
  player1 = "Imran Khan"
  fmt.Println(player1)

  x = 1
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
}