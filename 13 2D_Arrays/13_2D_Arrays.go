package main

import "fmt"

func main() {
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println("")
	}

	fmt.Println("")
	for i, ele := range matrix {
		fmt.Println(i, ele)
	}

	fmt.Println("")
	for i, ele := range matrix {
		for j, num := range ele {
			fmt.Print("[", i, ",", j, "]", "->", num, " ")
		}
		fmt.Println("")
	}
}
