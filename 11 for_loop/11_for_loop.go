package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello for", i, "th time")
	}
	fmt.Println("\nCubes from 0 to 5")
	i := 0
	for i <= 5 {
		fmt.Println(i * i * i)
		i++
	}
	fmt.Println("\nEven numbers upto 9")
	for num := 0; num <= 10; num++ {
		if num == 9 {
			break
		} else if num&1 == 1 {
			continue
		} else {
			fmt.Println(num)
		}
	}
}
