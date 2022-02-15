package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Registando la ruta y el handler
	http.HandleFunc("/saludar", saludar)
	//Sube un servidor en el puerto 8080
	http.ListenAndServe(":8080", nil)
}

func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo")
}
