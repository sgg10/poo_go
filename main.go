package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	fmt.Println("Creating a new employee...")
	return &Employee{id: id, name: name, vacation: vacation}
}

func main() {
	e := Employee{
		id:       1,
		name:     "Connor",
		vacation: true,
	}
	fmt.Println(e)

	e2 := new(Employee)
	fmt.Println(*e2)

	e3 := NewEmployee(2, "Ezio", true)
	fmt.Println(*e3)
}
