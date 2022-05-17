package main

import "fmt"

func main() {
	printMessage("Golang")
}
func printMessage(s string) {
	fmt.Println("This is", s)
}
