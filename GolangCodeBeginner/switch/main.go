package main

import "fmt"

func main() {
	emojin := "cat"
	switch {
	case emojin == "cat" || emojin == "dog":
		fmt.Println("Es un animal")
	case emojin == "banana" || emojin == "apple":
		fmt.Println("Es una fruta")
	default:
		fmt.Println("No se animal ni fruta.")
	}
}
