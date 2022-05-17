package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

// In the below code, struct square has implemented shape interface
// as function signatures are same for both square and shape
type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4 * s.side
}

// In the below code, struct rectangle has implemented shape interface
// as function signatures are same for both rectangle and shape
type rectangle struct {
	length, breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}
func (r rectangle) perimeter() float64 {
	return 2*r.breadth + 2*r.length
}

// Below function can be used for both square and rectangle for printing their area and perimeter.
func printData(s shape) {
	fmt.Printf("%+v\n", s)
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rectangle{length: 5, breadth: 10}
	s := square{side: 10}
	printData(r)
	fmt.Println()
	printData(s)
	// We used only one function for printing data related to two different structs using interfaces
}
