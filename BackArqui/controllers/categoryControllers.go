package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"../common"
	"../data"
)

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
