package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("Type: %v \n", reflect.TypeOf(100))
	fmt.Printf("Type: %v \n", reflect.TypeOf(78.4545))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
	fmt.Printf("Type: %v \n", reflect.TypeOf("hello"))

	var age int = 30
	var name string = "Rohit"
	fmt.Printf("Variable age = %v is of type %v \n", age, reflect.TypeOf(age))
	fmt.Printf("Variable name = %v is of type %v \n", name, reflect.TypeOf(name))
}

// Type: int
// Type: float64
// Type: bool
// Type: string
// Variable age = 30 is of type int
// Variable name = Rohit is of type string
