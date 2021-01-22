// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

import "fmt"

// Add imports.

// Declare a type named user.
type user struct {
	name string
	age  int
}

// Create a function that changes the value of one of the user fields.
func funcName(u *user) {

	// Use the pointer to change the value that the
	// pointer points to.
	fmt.Printf("%v \n", *u)
	u.name = "User2"
	u.age = 43
}

func main() {

	// Create a variable of type user and initialize each field.
	u := user{
		name: "User1",
		age:  42,
	}

	// Display the value of the variable.
	fmt.Printf("user: [%v] \n", u)

	// Share the variable with the function you declared above.
	funcName(&u)

	// Display the value of the variable.
	fmt.Printf("user after funcName: [%v] \n", u)

	// Struc with variable
	var u3 user
	u3.name = "User3"
	u3.age = 43

	fmt.Printf("user: [%v] \n", u3)
}
