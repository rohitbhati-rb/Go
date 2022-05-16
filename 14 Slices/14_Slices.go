package main

import "fmt"

func main() {
	// Creating slice with given values
	slice := [5]int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println("")

	// Creating slice from an array
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("nums =>", nums)

	slice1 := nums[2:7] // from index 2 to index 6
	fmt.Println("slice1 =>", slice1)

	slice2 := nums[:5] // from index 0 to 4
	fmt.Println("slice2 =>", slice2)

	slice3 := nums[6:] // from index 6 till end
	fmt.Println("slice3 =>", slice3)

	slice4 := nums[:] // all elements are included
	fmt.Println("slice4 =>", slice4)

	// We can also create a slice from a slice
	slice5 := slice3[2:4]
	// slice3 => [7, 8, 9, 10] slice5 have elements from index 2 to index 3
	fmt.Println("slice5 =>", slice5)
	fmt.Println("")

	// Creating slice using make() function
	slice6 := make([]int, 10, 50)
	fmt.Println("slice6 =>", slice6)
	fmt.Println("len(slice6) =>", len(slice6))
	fmt.Println("cap(slice6) =>", cap(slice6))
	fmt.Println("")

	// Modifying slice values
	fmt.Println("nums =>", nums)
	fmt.Println("slice1 =>", slice1)
	fmt.Println("After modifying slice 1")
	for i, e := range slice1 {
		slice1[i] = e * 10
	}
	fmt.Println("nums =>", nums)
	fmt.Println("slice1 =>", slice1)

}
