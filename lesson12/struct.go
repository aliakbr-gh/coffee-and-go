package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	// A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

	// While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

	// A struct can be useful for grouping data together to create records.

	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	fmt.Println(reflect.TypeOf(pers1.name)) // string
	fmt.Println(reflect.TypeOf(pers1.age))  // int
	fmt.Println(reflect.TypeOf(pers1.job))  // string
	fmt.Println(reflect.TypeOf(pers1.salary))

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	fmt.Println(pers1)
	fmt.Println(pers2)
	printPerson(pers2)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
