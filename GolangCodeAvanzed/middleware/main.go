package main

import "github.com/Rajual/Learning-Golang/tree/main/GolangCodeAvanzed/middleware/functions"

func execute(name string, f functions.MyFunction) {
	f(name)
}

func main() {
	name := "Julio"
	execute(name, functions.MiddlerwareLog(functions.Saludar))
	execute(name, functions.MiddlerwareLog(functions.Despedirse))
}
