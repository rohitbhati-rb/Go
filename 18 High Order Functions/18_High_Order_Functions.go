package main

import "fmt"

func calcArea(r float64) float64 {
	return 3.14 * r * r
}
func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}
func calcDiameter(r float64) float64 {
	return 2 * r
}

func getResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println(result)
}
func getFunction(query int) func(r float64) float64 {
	query_to_func := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]
}
func main() {
	var query int = 1
	var radius float64
	fmt.Printf("Radius: ")
	fmt.Scanf("%f\n", &radius)
	fmt.Println("Enter \n 1 - Area \n 2 - Perimeter \n 3 - Diameter")
	fmt.Printf("Enter your choice: ")
	n, err := fmt.Scanf("%v", &query)
	if err != nil || n != 1 {
		fmt.Println(n, err)
	}
	getResult(radius, getFunction(query))
}
