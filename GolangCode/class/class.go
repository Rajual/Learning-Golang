package main

import "fmt"

type Emploee struct {
	id       int
	name     string
	vacation bool
}

func (e *Emploee) SetId(id int) {
	e.id = id
}

func (e *Emploee) SetName(name string) {
	e.name = name
}

func (e *Emploee) GetId() int {
	return e.id
}

func (e *Emploee) GetName() string {
	return e.name
}

//4.
func NewEmploee(id int, name string, vacation bool) *Emploee {
	return &Emploee{
		id:       id,
		name:     name,
		vacation: vacation,
	}

}

func main() {
	e := Emploee{}
	e.id = 0
	e.name = "Golang"
	fmt.Println(e)
	e.SetId(5)
	fmt.Println(e)
	e.SetName("Goku")
	fmt.Println(e)

	//Constructor
	//1.
	e1 := Emploee{}
	fmt.Println(e1)
	//2.
	e2 := Emploee{
		id:       1,
		name:     "Vegeta",
		vacation: true,
	}
	fmt.Println(e2)
	//3.
	e3 := new(Emploee)
	fmt.Println(e3)
	//4.
	e4 := NewEmploee(1, "Name", false)
	fmt.Println(*e4)
}
