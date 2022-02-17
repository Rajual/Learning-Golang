package handler

import (
	"net/http"

	"github.com/Rajual/Learning-Golang/tree/main/GolangCodeAvanzed/my_first_API/middlewares"
)

//RoutePerson
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)
	mux.HandleFunc("/v1/persons/created", middlewares.Authetication(h.create))
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", middlewares.Log(h.delete))
	mux.HandleFunc("/v1/persons/get-all", middlewares.Log(h.getAll))
}
