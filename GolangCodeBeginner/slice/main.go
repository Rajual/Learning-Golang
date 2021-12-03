package main

import "fmt"

func main() {
	//Slice no poseen datos, son apuntadores a un Array
	set := [7]string{":lion", ":hourse", ":cot", ":cat", ":plane"}
	animals := set[0:3]
	fly := set[3:4]
	fly[0] = "Hola Mundo"

	fmt.Println("Array: ", set)
	fmt.Println("Animales: ", animals)
	fmt.Printf("Tipo: %T", animals)
	fmt.Println("Voladores: ", fly)

	//Los slice tienen:
	// len(): # de elementos
	// cap(): # de elementos donde apunta slice a partir
	//	del primer indice del array donde apunta
	fmt.Println("len(): ", len(animals))
	fmt.Println("cap(): ", cap(animals))

	//Slice de referencia anonima
	newSlice := []string{"Hola"}
	fmt.Println("cap(): ", cap(newSlice))

	//Conastruir array con make()
	makeNewSlice := make([]string, 0, 3)
	makeNewSlice = append(makeNewSlice, "Fruta")
	fmt.Println("cap(): ", cap(makeNewSlice))
}
