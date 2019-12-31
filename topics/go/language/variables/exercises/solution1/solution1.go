// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
    var a string
    var b int
    var c bool

	// Display the value of those variables.
    fmt.Printf("Value of declared variable a: %v\n", a)
    fmt.Printf("Value of declared variable b: %v\n", b)
    fmt.Printf("Value of declared variable c: %v\n", c)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
    e := "12-11-10"
    
    // Redeclate b:

    // topics/go/language/variables/exercises/solution1/solution1.go:31:7: no new variables on left side of :=
    //b := 2
    
    // topics/go/language/variables/exercises/solution1/solution1.go:32:9: b redeclared in this block
    //        previous declaration at topics/go/language/variables/exercises/solution1/solution1.go:20:9
    //var b int = 2


    // topics/go/language/variables/exercises/solution1/solution1.go:42:9: b redeclared in this block
    //        previous declaration at topics/go/language/variables/exercises/solution1/solution1.go:20:9
    //var b = 2

    b = 2

    var f = false

	// Display the value of those variables.
    fmt.Printf("Value of declared end initialized variable e: %v\n", e)
    fmt.Printf("Value of declared end initialized variable b: %v\n", b)
    fmt.Printf("Value of declared end initialized variable f: %v, type: %T\n", f, f)

	// Perform a type conversion.
    var pi float32
    pi = float32(3.14)

	// Display the value of that variable.
    fmt.Printf("Value of declared and later initialized pi: %v, type: %T\n", pi, pi)

}
