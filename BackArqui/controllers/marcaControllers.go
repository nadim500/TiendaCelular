package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"../common"
	"../data"
)

/*GetMarcas devuelve todas las marcas disponibles*/
func GetMarcas(w http.ResponseWriter, r *http.Request) {
	marcas, err := data.GetMarcas()
	if err != nil {
		common.DisplayError(w, err, "Error en la base de datos", 500)
		return
	}
	j, err := json.Marshal(MarcasResource{Data: marcas})
	if err != nil {
		log.Printf("[Error marshal get marcas]: %s\n", err)
		common.DisplayError(w, err, "Error convert to json", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*GetMarcaByID devuelve una marca por su id*/
func GetMarcaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	marca, err := data.GetMarcaByID(id)
	if err != nil {
		common.DisplayError(w, err, "Error en la base de datos", 500)
		return
	}
	j, err := json.Marshal(MarcaResource{Data: marca})
	if err != nil {
		log.Printf("[Error marshal get marca by id]: %s\n", err)
		common.DisplayError(w, err, "Error convert to json", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*CreateMarca devuelve la marca creada*/
func CreateMarca(w http.ResponseWriter, r *http.Request) {
	var dataResource MarcaResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		log.Printf("[Error decode create marca]: %s\n", err)
		common.DisplayError(w, err, "Datos inválidos", 400)
		return
	}
	marca := &dataResource.Data
	err = data.CreateMarca(marca)
	if err != nil {
		common.DisplayError(w, err, "Error en la base de datos", 500)
		return
	}
	j, err := json.Marshal(MarcaResource{Data: *marca})
	if err != nil {
		log.Printf("[Error marshal create category]: %s\n", err)
		common.DisplayError(w, err, "Error convert to json", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

/*UpdateMarca devuelve una marca actualizada*/
func UpdateMarca(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var dataResource MarcaResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		log.Printf("[Error decode in update marca]: %s\n", err)
		common.DisplayError(w, err, "Datos inválidos", 400)
		return
	}
	marca := &dataResource.Data
	err = data.UpdateMarca(marca, id)
	if err != nil {
		common.DisplayError(w, err, "Error la base de datos", 500)
		return
	}
	j, err := json.Marshal(MarcaResource{Data: *marca})
	if err != nil {
		log.Printf("[Error marshal create marca]: %s\n", err)
		common.DisplayError(w, err, "Error convert to json", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*DeleteMarca no devuelve nada si elimina*/
func DeleteMarca(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := data.DeleteMarca(id)
	if err != nil {
		common.DisplayError(w, err, "Error al eliminar la marca", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
