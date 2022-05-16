package main

import "fmt"

func main() {
	nums := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println("nums[\"two\"] = ", nums["two"])

	// Using make() function
	languages := make(map[string]string)
	fmt.Println("languages =>", languages)

	// Adding, Updating and Deleting the map
	languages["cpp"] = "C++"
	languages["c"] = "C"
	languages["js"] = "Javascript"
	languages["ts"] = "Typescript"
	languages["go"] = "Golang"
	fmt.Println("languages =>", languages)
	languages["cpp"] = "CPP"
	fmt.Println("languages =>", languages)
	delete(languages, "c") // Deletes the key value pair if it is present in the map
	fmt.Println("languages =>", languages)

	// Accessing a key value pair
	value, found := languages["go"]
	fmt.Println(value, found) // value is Golang and found is true
	value, found = languages["java"]
	fmt.Println(value, found) // value is a zero value of string and found is false

	// Looping through a map
	for key, value := range languages {
		fmt.Println(key, "=>", value)
	}
	// Truncating / Clearing the map
	languages = make(map[string]string)
	fmt.Println("languages =>", languages)
}
