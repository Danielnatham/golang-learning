package main

import "fmt"

func main() {
	age := 35
	name := "Daniel"

	// Print
	fmt.Print("Hello")
	fmt.Print(" World \n")

	// Println
	fmt.Println("My age is:", age, "My name is:", name)

	// Printf
	// %_ where _ is a format specifier 
	
	// %v is for variables
	fmt.Printf("My age is %v my name is %v \n", age, name)

	// %q is for quotes
	fmt.Printf("My age is %q my name is %q \n", age, name)
	
	// %T is for types
	fmt.Printf("My age is %T my name is %T \n", age, name)
	
	// %f is for float
	fmt.Printf("Float: %0.2f \n", 20.2)
}
