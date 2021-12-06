package main

import "fmt"

func main() {

	//If normal
	emojin := ":cat"

	if emojin == ":cactus" {
		fmt.Println("Es un cactus")
	} else if emojin == ":face" {
		fmt.Println("Es una carita")
	} else {
		fmt.Printf("El emojin es : %v", emojin)
		fmt.Println("")
	}

	//Varible propia del score del if
	if variable := "Gato"; variable == "Cactus" {
		fmt.Println("Es un cactus")
	} else if variable == "Face" {
		fmt.Println("Es una carita")
	} else {
		fmt.Printf("El emojin es : %s", variable)
		fmt.Println("")
	}
}
