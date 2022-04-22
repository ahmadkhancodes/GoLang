package main

import (
	"fmt"
)

func main() {
	// Declaration of Variables & Printing them
	var name string = "Ahmad" // Type Explicitly Defined
	var number int = 10       // Type Explicitly Defined
	var name2 = "Ahmad"       // Type Implicitly Defined
	var number2 = 10          // Type Implicitly Defined
	name3 := "Ahmad"          // More Professional
	number3 := 10             // More Professional
	number++
	number2++
	number3++
	fmt.Println("Hello", name, number)
	fmt.Println("Hello", name2, number2)
	fmt.Println("Hello", name3, number3)
}
