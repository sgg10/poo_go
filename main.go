package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id        int
	vacations bool
}

type FullTimeEmployee struct {
	Person
	Employee
}

// func NewFulltimeEmployee(id int, name string, age int, vacations bool) *FulltimeEmployee {
// 	fte := FullTimeEmployee{}
// 	fte.id = id
// 	fte.name = name
// 	fte.age = age
// 	fte.vacations = vacations
// 	return &fte
// }

func GetMessage(p Person) {
	fmt.Printf("%s with age: %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "name"
	ftEmployee.age = 20
	ftEmployee.id = 1
	ftEmployee.vacations = true
	fmt.Println(ftEmployee)
	//fmt.Println(ftEmployee.)
}
