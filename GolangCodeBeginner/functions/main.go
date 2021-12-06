package main

import (
	"fmt"
	"strings"
)

func main() {
	//Parametros por valor
	helloWork("Julio", "Rodiguez")

	//Paramtros por referencia
	name := "Julio"
	change(&name)
	fmt.Println(name)

	//Funciones con retorno
	fmt.Println(sum(1, 6))

	//Retorno de multiples valores
	fmt.Println(convert("Julio"))

}

func convert(text string) (string, string) {
	return strings.ToLower(text), strings.ToUpper(text)
}

func sum(a, b int) int {
	return a + b
}

func helloWork(name string, lastName string) {
	fmt.Printf("¡Hola %s %s!", name, lastName)
	fmt.Println("")
}

func change(name *string) {
	*name = "Alfonso"
}
