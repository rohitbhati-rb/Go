package main

import "fmt"

func main() {
	nums := [4]int{10, 20, 30, 40}
	fmt.Println(nums)

	slice := nums[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Appending new values to a slice
	slice = append(slice, 11, 22, 33)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println("")

	// Appending 2 slices to a new slice
	nums2 := [5]int{1, 2, 3, 4, 5}
	slice2 := nums2[:2]
	new_slice := append(slice, slice2...)
	fmt.Println(new_slice)
	for i, e := range new_slice {
		new_slice[i] = e * 10
		// here slice and slice2 will not get affected
	}
	fmt.Println(new_slice)
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println("")

	// Deleting/excluding a vlaue in slice
	// This is done by creating a new slice
	i := 2
	temp1 := nums2[:i]
	temp2 := nums2[i+1:]
	new_slice2 := append(temp1, temp2...)
	fmt.Println(new_slice2)
	fmt.Println("")

	// Copying a slice
	src_slice := []int{11, 22, 33, 44, 55}
	dest_slice := make([]int, 3)
	// capacity is not passed, it will be same as size
	num := copy(dest_slice, src_slice)
	// This value is the minimum of the capacities of both the slices passed to the copy function.
	fmt.Println(dest_slice)
	fmt.Println(num)
}
