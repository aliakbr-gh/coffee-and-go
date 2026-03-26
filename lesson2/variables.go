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


  var l, m, n, o int = 1, 3, 5, 7 // If use the type keyword, it is only possible to declare one type of variable per line.

  fmt.Println(l)
  fmt.Println(m)
  fmt.Println(n)
  fmt.Println(o)

  var e, f = 6, "Hello"
  g, h := 7, "World!"

  fmt.Println(e)
  fmt.Println(f)
  fmt.Println(g)
  fmt.Println(h)

  var (
    variable1 int
    variable2 int = 1
    variable3 string = "hello"
  )

 fmt.Println(variable1)
 fmt.Println(variable2)
 fmt.Println(variable3)

 // A variable name must start with a letter or an underscore character (_)
 // A variable name cannot start with a digit
 // A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
 // Variable names are case-sensitive (age, Age and AGE are three different variables)
 // There is no limit on the length of the variable name
 // A variable name cannot contain spaces
 // The variable name cannot be any Go keywords
}