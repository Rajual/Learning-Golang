package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	enDate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Funll time employee"
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Vegeta"
	ftEmployee.age = 30
	ftEmployee.id = 0
	fmt.Println(ftEmployee)

	tEmploye := TemporaryEmployee{}
	getMessage(tEmploye)
	getMessage(ftEmployee)
}
