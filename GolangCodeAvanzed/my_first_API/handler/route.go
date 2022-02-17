package handler

import "net/http"

//RoutePerson
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)
	mux.HandleFunc("/v1/persons/created", h.create)
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", h.delete)
	mux.HandleFunc("/v1/persons/get-all", h.getAll)
}
