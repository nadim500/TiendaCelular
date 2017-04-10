package controllers

import (
	"encoding/json"
	"net/http"

	"../data"
)

/*GetCategories devuelve todas las categorias disponibles*/
func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := data.GetCategories()
	if err != nil {
		return
	}
	j, err := json.Marshal(CategoriesResource{Data: categories})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
