package main

import "fmt"

func main() {
	var lang string = "Golang"
	if lang == "C++" {
		fmt.Println("This is C++")
	} else { // This else statement must be on the same line where the if block ends
		fmt.Println("This is Golang")
	}
}
