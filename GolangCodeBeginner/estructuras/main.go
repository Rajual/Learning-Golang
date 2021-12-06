package main

import "fmt"

func main() {
	type course struct {
		Name     string
		Profesor string
		Country  string
	}

	flutter := course{
		Name:     "Flutter",
		Profesor: "Julio",
		Country:  "Colombia",
	}

	git := course{"Git", "Alfonso", "Colombia"}

	golang := course{Name: "Golang"}

	p := &golang

	p.Name = "GO"

	fmt.Printf("%+v", flutter)
	fmt.Println("")
	fmt.Printf("%+v", git)
	fmt.Println("")
	fmt.Printf("%+v", golang)
	fmt.Println("")
}
