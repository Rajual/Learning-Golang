package main

import "fmt"

func main() {
	division(10, 2)
	division(102, 29)
	division(10, 0)
	division(104, 21)
	division(103, 22)
}

func division(dividendo, divisor int) {
	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("Mala operacion")
	}
}
