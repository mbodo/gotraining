// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import (
	"fmt"
)

// Add imports.

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int8
}

func main() {

	// Declare variable of type user and init using a struct literal.
	ul := user{
		name:  "Literal",
		email: "literal@struct.com",
		age:   20,
	}

	// Display the field values.

	fmt.Printf("%+v\n", ul)

	fmt.Println("== User Literal ==")
	fmt.Println("name: ", ul.name)
	fmt.Println("email: ", ul.email)
	fmt.Println("age: ", ul.age)

	// Declare a variable using an anonymous struct.

	ua := struct {
		name  string
		email string
		age   int8
	}{
		age:   30,
		name:  "Anonymous",
		email: "anonymous@struct.com",
	}

	fmt.Printf("%+v\n", ua)

	fmt.Println("== User Anonymous ==")
	fmt.Println("age: ", ua.age)
	fmt.Println("name: ", ua.name)
	fmt.Println("email: ", ua.email)

	// Display the field values.
}
