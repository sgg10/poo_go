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
	endDate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Fulltime employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Fulltime employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "name"
	ftEmployee.age = 20
	ftEmployee.id = 1
	ftEmployee.vacations = true
	fmt.Println(ftEmployee)
	getMessage(ftEmployee)
}
