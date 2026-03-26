// The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

package main

import (
	"fmt"
)

const PI = 3.14
const A int = 1

func main() {
  fmt.Println(PI)
  fmt.Println(A)


  const (
	X int = 1
	Y = 3.14
	Z = "Hi!"
  )

  fmt.Println(X)
  fmt.Println(Y)
  fmt.Println(Z)

}
