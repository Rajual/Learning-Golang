package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//Parametros por valor
	helloWork("Julio", "Rodiguez")

	//Paramtros por referencia
	name := "Julio"
	change(&name)
	fmt.Println(name)

	//Funciones con retorno
	fmt.Println(sum(1, 6))

	//Retorno de multiples valores
	fmt.Println(convert("Julio"))

	//Controlar errores
	content, err := ioutil.ReadFile("./things.txt")
	if err == nil {
		fmt.Println(err)
		return
	}
	//Retorna pues esta el error
	/* else {
		fmt.Println(string(content))
	}*/
	fmt.Println(string(content))
	resultado, err := division(10, 2)
	if err != nil {
		return
	}
	fmt.Println(resultado)

	//Funciones que reciben y retornan funciones
	//Se recibe una funcion como parametro
	nums := []int{1, 4, 70, 5, 67, 90, 2}
	filterNumber := filter(nums, func(i int) bool {
		if i <= 10 {
			return true
		}
		return false
	})
	fmt.Println(filterNumber)
	//Se retorna una funcion
	fmt.Println(hello("Julio")("¿Como estas?"))

	//Funciones variaticas
	fmt.Println(suma(1, 2, 3, 4))

	//Funciones anonimas
	f := func() {
		fmt.Println("Hello Julio!")
	}
	f()
	//Puntero a la funcion
	fmt.Println(f)
}

//Funciones variatica
func suma(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

//Funcioes que retornan una funcion
func hello(name string) func(message string) string {
	return func(m string) string {
		return "Hello " + name + " " + m
	}
}

//Funciones que reciben funciones
func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}

	}
	return result
}

func convert(text string) (string, string) {
	return strings.ToLower(text), strings.ToUpper(text)
}

func sum(a, b int) int {
	return a + b
}

func helloWork(name string, lastName string) {
	fmt.Printf("¡Hola %s %s!", name, lastName)
	fmt.Println("")
}

func change(name *string) {
	*name = "Alfonso"
}

//Parametros nombrado
func division(dividendo, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("No puedes divider entre 0")
		return
	}
	result = dividendo / divisor
	return
}
