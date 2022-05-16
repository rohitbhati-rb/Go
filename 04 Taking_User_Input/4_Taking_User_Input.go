package main

import "fmt"

func main() {
	var (
		username string
		age      int
	)
	fmt.Print("Enter Your Name and Age: ")
	fmt.Scanf("%s %d", &username, &age)
	fmt.Println("Hello,", username, age)
}
