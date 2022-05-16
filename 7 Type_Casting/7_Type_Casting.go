package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Integer to Float
	var i int = 23
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)

	// Float to Integer
	f = 123.456
	i = int(f)
	fmt.Printf("%d\n", i)

	// Integer to String
	i = 456
	var str string = strconv.Itoa(i)
	fmt.Printf("%q\n", str)

	// String to Integer
	str = "789"
	i, err := strconv.Atoi(str)
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("%v %T \n", err, err)
}

// 23.00
// 123
