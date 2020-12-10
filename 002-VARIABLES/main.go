package main

import "fmt"

var y = 42 // var can be used to decalre variables outside the function body unlinke the shorthand one

// Declare there is a variable with a variable name
// and that identifier is of TYPE int
// assigns the ZERO (default value for every specific data type) value of int to the variable

var z int

func main() {
	x := 42 // shorthand operator declares and assigns the value
	x = 99
	fmt.Println(x)
	fmt.Println(z)
	foo()
}

func foo() {
	fmt.Println("I'm in foo: ", y)
}
