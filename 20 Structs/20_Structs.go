package main

import "fmt"

type Student struct {
	name   string
	phone  string
	rollno int
	marks  map[string]int
	grades map[string]string
}

func main() {
	var s Student
	fmt.Printf("%+v\n\n", s) // {name: phone: rollno:0 marks:map[] grades:map[]}

	st := new(Student)
	fmt.Printf("%T %+v\n\n", st, st) // *main.Student &{name: phone: rollno:0 marks:map[] grades:map[]}

	st2 := Student{
		name:   "Rohit",
		phone:  "1234567890",
		rollno: 1,
	}
	fmt.Printf("%T %+v\n\n", st2, st2) // main.Student {name:Rohit phone:1234567890 rollno:1 marks:map[] grades:map[]}

	st3 := Student{"Rohit", "1234567890", 1, make(map[string]int), make(map[string]string)}
	fmt.Printf("%T %+v\n\n", st3, st3)

	st3.marks["Maths"] = 100
	st3.marks["Coding"] = 80
	st3.marks["Golang"] = 50

	fmt.Printf("%T %+v\n\n", st3, st3)
	addGrades(&st3)
	fmt.Printf("%T %+v\n\n", st3, st3)
}
func addGrades(st *Student) {
	st.grades["Maths"] = "A+"
	st.grades["Coding"] = "B"
	st.grades["Golang"] = "D"
}
