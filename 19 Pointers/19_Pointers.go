package main

import "fmt"

func main() {
	num := 10
	fmt.Printf("%T %v \n", &num, &num)       // *int 0xc000014098  & -> Address Operator
	fmt.Printf("%T %v \n", *(&num), *(&num)) // int 10 * -> Dereference Operator
	fmt.Println("")

	// Declaring a Pointer
	var i *int
	var s *string
	fmt.Println(i) // <nil>
	fmt.Println(s) // <nil>
	fmt.Println("")

	// Initializing a Pointer
	str := "Golang"
	ptr_str := &str
	fmt.Println(str)     // Golang
	fmt.Println(ptr_str) // 0xc000052250
	fmt.Println("")

	// Dereferencing a Poionter
	fmt.Println(*ptr_str) // Golang
	*ptr_str = "Go"
	fmt.Println(*ptr_str) // Go
	fmt.Println("")

	// Pass by value
	fmt.Println(str) // Go
	modifyStr(str)
	fmt.Println(str) // Go
	fmt.Println("")

	// Pass by reference
	msg := "Hello"
	fmt.Println(msg) // Hello
	changeMsg(&msg)
	fmt.Println(msg) // Welcome
	fmt.Println("")

}
func modifyStr(s string) {
	s = "C++"
}
func changeMsg(s *string) {
	*s = "Welcome"
}
