package main

import (
	"fmt"
	"os"
)

func main() {
	a := 3
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(a)
	defer fmt.Println(4)
	a *= 2
	fmt.Println("Nuevo valor de: ", a)

	//Leer archivos
	file, err := os.Create("MyFirtFile.txt")
	if err != nil {
		fmt.Printf("Ocurrio un error al crear el archivo %v", err)
		fmt.Println("")
		return
	}
	defer file.Close()
	_, err = file.Write([]byte("Hola Mundo..."))
	if err != nil {
		file.Close()
		fmt.Printf("Ocurrio un error al escribir %v", err)
		return
	}
}
