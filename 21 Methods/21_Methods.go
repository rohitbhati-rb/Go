package main

import "fmt"

type Circle struct {
	radius float64
	area   float64
}

// The below function can only be accessed by the instances of Circle struct
func (c *Circle) calcArea() {
	c.area = 3.14 * c.radius * c.radius
}
func (c Circle) calcArea2() {
	c.area = 3.14 * c.radius * c.radius
}

type Student struct {
	name   string
	grades []int
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}
func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, val := range s.grades {
		sum += val
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func main() {
	c := Circle{radius: 5}
	c.calcArea2()
	fmt.Printf("%+v\n", c) // {radius:5 area:0}
	c.calcArea()
	fmt.Printf("%+v\n\n", c) // {radius:5 area:78.5}

	// method sets
	s := Student{name: "Rohit", grades: []int{89, 94, 90, 88, 91}}
	s.displayName()
	fmt.Println(s.calculatePercentage())
}
