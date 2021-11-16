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
}

/*
func GetMessage(p Person) {
	fmt.Println("%s with age %d", p.name, p.age)
}
*/
func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Vegeta"
	ftEmployee.age = 30
	ftEmployee.id = 0
	fmt.Println(ftEmployee)
}
