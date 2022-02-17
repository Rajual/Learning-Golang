package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Rajual/Learning-Golang/tree/main/GolangCodeAvanzed/my_first_API/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La estructura no es valida", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Create(&data)

	if err != nil {
		response := newResponse(Error, "Problema al crear a la persona", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := newResponse(Message, "Nuevo usuario creado", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "Problema al crear a la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "Problema al actualizar", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Persona actualizada", nil)
	responseJSON(w, http.StatusOK, response)

}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}
	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El id de la persona no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "Ocurrio un errror a eliminar", nil)
		responseJSON(w, http.StatusInternalServerError, response)
	}
	response := newResponse(Error, "Ok", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Problema al obtener todas las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, response)
}
