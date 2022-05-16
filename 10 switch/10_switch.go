package main

import "fmt"

func main() {
	var lang string = "Golang"
	switch lang {
	case "C", "C++": // equivalent to -> if lang == "C" || lang == "C++"
		fmt.Println("This is C or C++")
	case "Java":
		fmt.Println("This is C++")
	case "Golang":
		fmt.Println("This is Golang")
	default:
		fmt.Println("This is Useless")
	}

	// switch case with fallback
	var year int = 2021
	switch year {
	case 2022:
		fmt.Println(2022)
	case 2021:
		fmt.Println(2021)
		fallthrough
	case 2020:
		fmt.Println("Covid")
		fallthrough
	default:
		fmt.Println("default")
	}
}
