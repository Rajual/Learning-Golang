package main

import "github.com/Rajual/Learning-Golang/tree/main/GolangCodeAvanzed/middleware/functions"

func execute(name string, f func(string)) {
	f(name)
}

func main() {
	name := "Julio"
	execute(name, functions.Saludar)
	execute(name, functions.Despedirse)
}
