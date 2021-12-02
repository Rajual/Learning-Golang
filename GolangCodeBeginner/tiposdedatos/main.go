package main

import "fmt"

func main() {

	//bool
	var a bool = true
	fmt.Printf("Tipo: %T, Valor: %v", a, a)
	fmt.Println("")

	//string
	var e string = "Julio"
	fmt.Printf("Tipo: %T, Valor: %v", e, e)
	fmt.Println("")

	//uint
	var i uint8 = 200
	fmt.Printf("Tipo: %T, Valor: %v", i, i)
	fmt.Println("")

	//rune
	var o rune = 'a'
	fmt.Printf("Tipo: %T, Valor: %v", o, o)
	fmt.Println("")

	//float
	var u float64 = 24.5
	fmt.Printf("Tipo: %T, Valor: %v", u, u)
	fmt.Println("")

	//Casting
	c := rune(i) + o
	fmt.Printf("Tipo: %T, Valor: %v", c, c)
	fmt.Println("")

	//var sin usar
	_ = 123
	var _ string = "321"
}
