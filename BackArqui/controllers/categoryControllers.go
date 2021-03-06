package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"../common"
	"../data"
)

/*GetCategories devuelve todas las categorias disponibles*/
func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := data.GetCategories()
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Error en el servidor",
			500,
		)
		return
	}
	j, err := json.Marshal(CategoriesResource{Data: categories})
	if err != nil {
		log.Printf("[Error marshal get categories]: %s\n", err)
		common.DisplayError(
			w,
			err,
			"Error convert to json",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*GetCategoryByID devuelve una categoria por su id*/
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	category, err := data.GetCategoryByID(id)
	if err != nil {
		log.Printf("[Error en get by id category] %s\n", err)
		common.DisplayError(
			w,
			err,
			"No existe la categoria",
			500,
		)
		return
	}
	j, err := json.Marshal(CategoryResource{Data: category})
	if err != nil {
		log.Printf("[Error marshal get category by id]: %s\n", err)
		common.DisplayError(
			w,
			err,
			"Error convert to json",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*CreateCategory devuelve la categoria creada*/
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var dataResource CategoryResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		log.Printf("[Error decode in create category]: %s\n", err)
		common.DisplayError(
			w,
			err,
			"Datos invalidos",
			500,
		)
		return
	}
	category := &dataResource.Data
	err = data.CreateCategory(category)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Error in database",
			500,
		)
		return
	}
	j, err := json.Marshal(CategoryResource{Data: *category})
	if err != nil {
		log.Printf("[Error marshal create category]: %s\n", err)
		common.DisplayError(
			w,
			err,
			"Error convert to json",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

/*UpdateCategory devuelve la categoria actualizada*/
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var dataResource CategoryResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		log.Printf("[Error decode in update category]: %s\n", err)
		common.DisplayError(
			w,
			err,
			"Datos invalidos",
			400,
		)
		return
	}
	category := &dataResource.Data
	err = data.UpdateCategory(category, id)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Error actualizando la category en la base de datos",
			500,
		)
		return
	}
	j, err := json.Marshal(CategoryResource{Data: *category})
	if err != nil {
		log.Printf("[Error marshal actualizar category]: %s\n", err)
		common.DisplayError(
			w,
			err,
			"Error convert to json",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*DeleteCategory devulve OK si se elimino*/
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := data.DeleteCategory(id)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Error al eliminar la categoria",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
