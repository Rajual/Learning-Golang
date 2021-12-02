package main

import "fmt"

func main() {
	fruit := ":apple"
	var p *string
	p = &fruit
	fmt.Printf("Tipo: %T, Valor: %s, Direccion: %v", fruit, fruit, &fruit)
	fmt.Println("")
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciador: %s", p, p, *p)
	fmt.Println("")
}
