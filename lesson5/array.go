package main

import (
	"fmt"
)

func main() {
  var arr1 = [4]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  var arr3 = [...]int{1,2,3}
  arr4 := [...]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
  fmt.Println(arr4)

  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)

  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])

  fmt.Print("\n")

  prices[2] = 50
  fmt.Println(prices)

  // arr1 := [5]int{} //not initialized
  // arr2 := [5]int{1,2} //partially initialized
  // arr3 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Print("\n")

  arrz := [5]int{1:10,3:40} // 1:10 means: assign 10 to array index 1 (second element)

  fmt.Println(arrz)

  fmt.Println("Length of arrz is: ", len(arrz))
}