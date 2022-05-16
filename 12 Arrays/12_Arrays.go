package main

import "fmt"

func main() {
	var words [5]string = [5]string{"hello", "there", "this", "is", "Golang"}
	fmt.Println(len(words))
	fmt.Println(words)

	odds := [5]int{1, 3, 5, 7, 9}
	fmt.Println(len(odds))
	fmt.Println(odds)

	evens := [...]int{2, 4, 6, 8, 10}
	fmt.Println("len(evens) = ", len(evens))
	fmt.Println("evens[] = ", evens)
	fmt.Println("evens[0] = ", evens[0])

	fmt.Println("\nLooping through an Array")
	for i := 0; i < len(evens); i++ {
		fmt.Println(i, evens[i])
	}

	fmt.Println("\nLooping through an Array using range keyword")
	for index, element := range evens {
		fmt.Println(index, "->", element)
	}
}
