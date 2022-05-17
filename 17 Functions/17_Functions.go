package main

import "fmt"

func main() {
	// A simple function
	printMessage("Golang")
	println("")

	// Returning Multiple values
	sum, product := solve(12, 10)
	fmt.Println(sum, product)

	// Returning named values
	diff, quotient := solve2(10, 3)
	fmt.Println(diff, quotient)
	fmt.Println("")

	// Variadic Function
	fmt.Println(summation())
	fmt.Println(summation(1, 2))
	fmt.Println(summation(1, 2, 3, 4, 5, 6))

	fmt.Println("")
	printDetails("Sam", "Robin", "David", "Peter")
	fmt.Println("")

	// Recursive Function
	printNums(5)
	fmt.Println("")

	// Anonymous Function
	fun := func(x int, y int) int {
		return x + y
	}
	fmt.Printf("%T \n", fun)
	fmt.Println(fun(12, 18))

	val := func(x int, y int) int {
		return x + y
	}(26, 24)
	fmt.Printf("%T \n", val)
	fmt.Println(val)
	fmt.Println("")
}

// A simple function
func printMessage(s string) {
	fmt.Println("This is", s)
}

// Function returning multiple values
func solve(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

// Function with named return values
func solve2(a int, b int) (diff int, quotient float64) {
	diff = a - b
	quotient = float64(a) / float64(b)
	// return diff, quotient
	return // no need to weite return variables as they are already specified above
}

// Function with variable no of arguments
func summation(nums ...int) int {
	sum := 0
	for _, value := range nums { // _ is used when we don't need index variable from range
		sum += value
	}
	return sum
}

func printDetails(name string, friends ...string) {
	fmt.Println("Hello,", name, "Your Friends are:")
	for _, friend := range friends {
		fmt.Print(friend, ", ")
	}
}

// Recursive Function
func printNums(num int) {
	if num == 0 {
		return
	}
	fmt.Println(num)
	printNums(num - 1)
}
