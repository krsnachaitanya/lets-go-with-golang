package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers.")

	// When "*" is put in front of a type, it becomes part of the type declaration, so you can say "this variable holds a pointer to a string".
	var ptr *string
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	// "&" returns the memory address of the following variable.
	// "*" returns the value of the following variable

	// "&" is used in front of a variable when you want to get that variable's memory address.
	var myNumberPtr = &myNumber
	fmt.Println("Value of pointer is ", myNumberPtr)

	// "*" gets the actual variable instead of memory address
	fmt.Println("Value of pointer is ", *myNumberPtr)
}
