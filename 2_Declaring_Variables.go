package main

import (
	"fmt"
)

// This is a global variable
var user string = "Rohit"

func main() {
	// These are all local variables
	var message string = "Hello there!!!"
	var i int = 123
	var b bool = true
	var f float64 = 987.123
	var str string
	str = "It's Golang baby"
	fmt.Println(message, i, b, f, str)

	// Different ways of declaring variables
	// Shorthand declaration,

	// works when declaring variables of the same data type
	var x, y, z int8 = 15, 68, 49

	// for different data types, should be on separate lines
	var (
		ss string = "hehe"
		aa int    = 987
		bb bool   = true
	)

	// Short Variable Declaration
	text := "This is a string"
	rank := 11

	fmt.Println(x, y, z, ss, aa, bb, text, rank, user)
	// We cant modify data type of a variable after declaring it
}
