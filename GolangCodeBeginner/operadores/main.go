package main

import "fmt"

func main() {
	//Operadores aritmeticos (), *, /, %, +, - (En ese orden de prioridad)
	var a = 4 + 2*3
	fmt.Printf("Tipo: %T, Valor: %v", a, a)
	fmt.Println("")

	//Operadores de asignacion: =, +=, -=, *=, /=, %=
	var b = 10
	b *= 2
	fmt.Printf("Tipo: %T, Valor: %v", b, b)
	fmt.Println("")

	//Declaracion post-incremento y post-decremento: ++, --
	//(No son una expresion, sino una declaracion)
	var c = 3
	//c=c++ + 1 => error
	c++
	fmt.Printf("Tipo: %T, Valor: %v", c, c)
	fmt.Println("")
	c--
	//fmt.Printf("Tipo: %T, Valor: %v", c--, c) => Error
	fmt.Printf("Tipo: %T, Valor: %v", c, c)
	fmt.Println("")

	//Operadores de comparacion: >, <, ==, !=,>=,<=
	fmt.Println(4 > 5)

	//Operadores logicos: &&, ||
	var age = 40
	fmt.Println("Edad:", age >= 18 && age <= 60)

	//Unario: !
	fmt.Println(!(6 == 4))

}
