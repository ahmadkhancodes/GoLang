package main

import (
	"fmt"
)

func main() {
	// Declaration of Variables & Printing them
	var name string = "Ahmad"  // Type Explicitly Defined
	var number int = 10        // Type Explicitly Defined
	var name2 string = "Ahmad" // Type Implicitly Defined
	var number2 int = 10       // Type Implicitly Defined
	number++
	number2++
	fmt.Println("Hello", name, number)
	fmt.Println("Hello", name2, number2)
}
