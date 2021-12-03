package main

import (
	"fmt"
)

func main() {
	//Usando make()
	animals := make(map[string]string)
	animals["cat"] = "cat"

	//Literal
	fruits := map[string]string{
		"apple": "manzana",
	}
	fmt.Println(animals, fruits)

	delete(fruits, ":apple")

	fmt.Println(animals, fruits)

}
