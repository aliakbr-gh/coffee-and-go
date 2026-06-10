package main

import "fmt"

func main() {
	// array without size defination basically
	myslice := []string{"a", "b", "c"}
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])
	fmt.Println(prices)

	myslice2 := []int{1, 2, 3}
	fmt.Println("myslice2", myslice2)
	myslice2 = append(myslice2, 4, 5)

	fmt.Println("myslice2 after append", myslice2)

	myslice3 := append(myslice1, myslice2...)
	fmt.Println("myslice3", myslice3)

	arr := [5]int{1, 2, 3, 4, 5}

	s := arr[:]

	fmt.Printf("%T\n", arr) // [5]int
	fmt.Printf("%T\n", s)   // []int

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
