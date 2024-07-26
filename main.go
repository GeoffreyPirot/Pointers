package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Println("Before increment:", x)
	// Call the Increment function from the pointers package
	// pointers.Increment(...)
	fmt.Println("After increment:", x)

	a, b := 1, 2
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	// Call the Swap function from the pointers package
	// pointers.Swap(....)
	fmt.Printf("After swap: a = %d, b = %d\n", a, b)

	// initialize a Person struct Name: "Electron", Age: 3
	// person := Person{...}
	fmt.Println("Before setAge:" /*person*/)
	// Call the setAge function from the pointers package and set the age to 30
	// pointers.SetAge(...)

	fmt.Println("After setAge:" /*person*/)

}
