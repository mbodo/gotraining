// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

// Add imports.

func main() {

	// Declare an array of 5 strings set to its zero value.
	var a1 [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	a2 := [5]string{"A", "B", "C", "D", "E"}

	// Assign the populated array to the array of zero values.
	a1 = a2

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for i, v := range a1 {
		fmt.Printf("Address: [%p] | Value: [%v] \n", &a1[i], v)
	}

	a1[0] = "CHANGE"

	fmt.Println("a1: ", a1[0])
	fmt.Println("a2: ", a2[0])

	fmt.Printf("a1[0]: %p \n", &a1[0])
	fmt.Printf("a2[0]: %p \n", &a2[0])
}
