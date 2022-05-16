package main

import "fmt"

func main() {
	var lang string = "Javascript"
	if lang == "C++" {
		fmt.Println("This is C++")
	} else if lang == "Golang" { // This else-if statement must be on the same line where the if block ends
		fmt.Println("This is Golang")
	} else if lang == "Javascript" { // This else-if statement must be on the same line where the if block ends
		fmt.Println("This is Javascript")
	} else { // This else statement must be on the same line where the if block ends
		fmt.Println("This is Useless")
	}
}
