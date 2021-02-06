// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

import (
	"fmt"
)

// Add imports.

const (
	// Declare a constant named server of kind string and assign a value.
	server = "localhost"

	// Declare a constant named port of type integer and assign a value.
	port int16 = 21
)

func main() {

	// Display the value of both server and port.
	fmt.Printf("Server:[%v], Port:[%v] \n", server, port)

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	const i = 2
	const f = 0.5
	var result = i / f

	unamedResult := 4 / 0.5

	// Display the value of the variable.
	fmt.Printf("Result: [%T %v] \n", result, result)
	fmt.Printf("UnamedResult: [%T %v] \n", unamedResult, unamedResult)

	// Examples
	fmt.Printf("%T %v\n", 'a', 'a')

	const MaxUint = ^uint(0)
	fmt.Printf("%T %x\n", MaxUint, MaxUint)

	var a = 3 * 0.333

	fmt.Printf("%T %v\n", a, a)
}
