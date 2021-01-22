// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and initialize a variable of type int with the value of 20. Display
// the _address of_ and _value of_ the variable.
//
// Declare and initialize a pointer variable of type int that points to the last
// variable you just created. Display the _address of_ , _value of_ and the
// _value that the pointer points to_.
package main

import (
	"fmt"
)

// Add imports.

func main() {

	// Declare an integer variable with the value of 20.
	a := 20

	// Display the address of and value of the variable.
	fmt.Printf("Address of: %p | Value of: %v \n", &a, a)

	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.
	b := &a

	// Display the address of, value of and the value the pointer
	// points to.
	fmt.Printf("Address of: %p | Value of: %v | Points to: %v \n", &b, b, *b)

	//dereference and assign to pointer
	fmt.Println("*b = 30")
	*b = 30

	fmt.Printf("Address of [b]: %p | Value of [b]: %v | Points to [b]: %v \n", &b, b, *b)

	// a after assign to variable b which holds a pointer address
	fmt.Printf("Address of [a]: %p | Value of [a]: %v \n", &a, a)
}
