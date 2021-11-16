package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func name(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

//Return with names
func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 4, 5))
	fmt.Println(sum(1, 2, 4, 5, 6, 7, 8))
	//
	name("Juio", "Alfonso")
}
