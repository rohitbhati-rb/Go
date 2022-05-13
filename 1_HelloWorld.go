// The package “main” tells the Go compiler that the package should compile
// as an executable program instead of a shared library. The main function in
// the package “main” will be the entry point of our executable program.
// When you build shared libraries, you will not have any main package
// and main function in the package.
package main

// fmt package implements formatted I/O with functions analogous to C’s printf() and scanf() function.
import "fmt"

// Entry point of executable programs
func main() {
	fmt.Println("Hello, World")
}
