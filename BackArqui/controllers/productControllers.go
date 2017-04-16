package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"../common"
	"../data"
)

/*GetProducts devuelve todos los productos disponibles*/
func GetProducts(w http.ResponseWriter, r *http.Request) {
	productos, err := data.GetProducts()
	if err != nil {
		common.DisplayError(w, err, "Error en la base de datos", 500)
		return
	}
	j, err := json.Marshal(ProductosResource{Data: productos})
	if err != nil {
		log.Printf("[Error marshal get productos]: %s\n", err)
		common.DisplayError(w, err, "Error convert to JSON", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*GetProductByID devuelve un product por su id*/
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	product, err := data.GetProductByID(id)
	if err != nil {
		common.DisplayError(w, err, "Error en la base de datos", 500)
		return
	}
	j, err := json.Marshal(ProductResource{Data: product})
	if err != nil {
		log.Printf("[Error marshal get product by id]: %s\n", err)
		common.DisplayError(w, err, "Error convert to json", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*CreateProduct devuelve el producto creado*/
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var dataResource ProductResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		log.Printf("[Error decode create product]: %s\n", err)
		common.DisplayError(w, err, "Datos inválidos", 400)
		return
	}
	product := &dataResource.Data
	err = data.CreateProduct(product)
	if err != nil {
		common.DisplayError(w, err, "Error en la base de datos", 500)
		return
	}
	j, err := json.Marshal(ProductResource{Data: *product})
	if err != nil {
		log.Printf("[Error marshal get product by id]: %s\n", err)
		common.DisplayError(w, err, "Error convert to json", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*UpdateProduct devuelve un producto actualizado*/
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var dataResource ProductResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		log.Printf("[Error decode in update product]: %s\n", err)
		common.DisplayError(w, err, "Datos inválidos", 400)
		return
	}
	product := &dataResource.Data
	err = data.UpdateProduct(product, id)
	if err != nil {
		common.DisplayError(w, err, "Error la base de datos", 500)
		return
	}
	j, err := json.Marshal(ProductResource{Data: *product})
	if err != nil {
		log.Printf("[Error marshal create marca]: %s\n", err)
		common.DisplayError(w, err, "Error convert to json", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

/*DeleteProduct no devuelve nada si elimina*/
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := data.DeleteProduct(id)
	if err != nil {
		common.DisplayError(w, err, "Error al eliminar la producto", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
