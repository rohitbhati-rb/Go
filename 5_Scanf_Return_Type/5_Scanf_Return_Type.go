package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)
	fmt.Print("Enter your name and age: ")
	count, err := fmt.Scanf("%s %d", &name, &age)
	fmt.Println("count: ", count)
	fmt.Println("err: ", err)
	fmt.Println("name: ", name)
	fmt.Println("age: ", age)
}

// Enter your name and age: Rohit Bhati
// count:  1
// err:  expected integer
// name:  Rohit
// age:  0

// Enter your name and age: Rohit 22
// count:  2
// err:  <nil>
// name:  Rohit
// age:  22
