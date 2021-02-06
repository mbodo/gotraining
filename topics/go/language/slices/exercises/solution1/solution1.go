// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var numbers = make([]int, 10, 10)

	// Append numbers to the slice.
	for i := 0; i < 10; i++ {
		numbers[i] = i
	}

	// Display each value in the slice.
	fmt.Println("Display each value in the slice:")
	for index, number := range numbers {
		fmt.Printf("index: [%d] | number: [%d] \n", index, number)
	}

	// Declare a slice of strings and populate the slice with names.
	characters := [5]string{"A", "B", "C", "D", "E"}

	// Display each index position and slice value.
	fmt.Println("Display each index position and slice value:")
	for index, character := range characters {
		fmt.Printf("index: [%d] | character: [%v] \n", index, character)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	characterslice := characters[1:3]

	fmt.Printf("characterslice => len: [%d] | cap: [%d] \n", len(characterslice), cap(characterslice))
	// Display each index position and slice values for the new slice.
	fmt.Println("Display each index position and slice values for the new slice:")
	for index, character := range characterslice {
		fmt.Printf("index: [%d] | character: [%v] \n", index, character)
	}
}
