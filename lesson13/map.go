package main

import "fmt"

func main() {
	// Maps are used to store data values in key:value pairs.

	// Each element in a map is a key:value pair.

	// A map is an unordered and changeable collection that does not allow duplicates.

	// The default value of a map is nil.

	// Maps hold references to an underlying hash table.
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	var z = make(map[string]string) // The map is empty now
	z["brand"] = "Z1"
	z["model"] = "Z2"
	z["year"] = "Z3"

	fmt.Printf("a\t%v\n", z)

	var y = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := y["brand"] // Checking for existing key and its value
	val2, ok2 := y["color"] // Checking for non-existing key and its value
	val3, ok3 := y["day"]   // Checking for existing key and its value
	_, ok4 := y["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	// Maps are references to hash tables.

	// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

	var e = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	f := e

	fmt.Println(e)
	fmt.Println(f)

	f["year"] = "1970"
	fmt.Println("After change to f:")

	fmt.Println(e)
	fmt.Println(f)

	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range m {
		fmt.Printf("%v : %v, ", k, v)
	}

	// a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	// var b []string             // defining the order
	// b = append(b, "one", "two", "three", "four")

	// for k, v := range a {        // loop with no order
	//   fmt.Printf("%v : %v, ", k, v)
	// }

	// fmt.Println()

	// for _, element := range b {  // loop with the defined order
	//   fmt.Printf("%v : %v, ", element, a[element])
	// }
}
