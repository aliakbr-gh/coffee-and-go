package main

import "fmt"

func main() {
	// for loop is the only loop exist in go

	count := 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("i == 3")
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("i == 2")
			break
		}
		fmt.Println(i)
	}

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

}
