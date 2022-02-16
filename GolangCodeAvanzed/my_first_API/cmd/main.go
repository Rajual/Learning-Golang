package main

import (
	"log"
	"net/http"

	"github.com/Rajual/Learning-Golang/tree/main/GolangCodeAvanzed/my_first_API/handler"
	"github.com/Rajual/Learning-Golang/tree/main/GolangCodeAvanzed/my_first_API/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()
	handler.RoutePerson(mux, &store)

	log.Println("Servidor inciiado")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error: %v", err)
	}

}
