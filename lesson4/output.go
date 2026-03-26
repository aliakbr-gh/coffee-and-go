package main

import (
	"fmt"
)

func main() {
  // Print()
  // Println()
  // Printf()

  // The Print() function prints its arguments with their default format.

  var i,j string = "Hello","World"

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
  fmt.Print(i, ",", j)
  fmt.Print("\n")

  // The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end
  fmt.Println(i,j)

  // The Printf() function first formats its argument based on the given formatting verb and then prints them
  // %v is used to print the value of the arguments
  // %T is used to print the type of the arguments
  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}