package main

import (
	"fmt"
)

func main() {
	// Printing string
	fmt.Print("Hello World")

	// Printing variable
	var city string = "Jaipur"
	fmt.Print(city)
	fmt.Print("\n\n")

	// Printing variable and string
	var name string = "Rohit"
	var language string = "Golang"
	fmt.Print(name, " is learning ", language, " programming.\n")

	fmt.Print(name, language, "\n")
	// RohitGolang
	fmt.Println(name, language)
	// Rohit Golang
	fmt.Println(name)
	fmt.Println(language)
	// Rohit
	// Golang

	// The printf statement
	fmt.Print("\n")
	fmt.Printf("Welcome to %v, %v\n", language, name)

	var boolean bool = false
	fmt.Printf("%t is a boolean\n", boolean)

	var num uint = 50
	fmt.Printf("%d is a number\n", num)

	var decimal float32 = 123.4567
	fmt.Printf("%f is a decimal\n", decimal)

	fmt.Printf("%.2f is a decimal upto 2 decimal places\n", decimal)

	var char string = "abc"
	fmt.Printf("%c is a character\n", char[1])

	var str string = "ABCabc"
	fmt.Printf("%s is a string\n", str)

	fmt.Printf("%q is a string in quotes\n", str)

	fmt.Printf("%T is a data type.\n", boolean)
	fmt.Printf("%T is a data type.\n", num)
	fmt.Printf("%T is a data type.\n", decimal)
	fmt.Printf("%T is a data type.\n", char)
	fmt.Printf("%T is a data type.\n", str)
}
