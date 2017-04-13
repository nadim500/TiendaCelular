package controllers

import "../models"

type (

	/*CategoryResource para mostrar la categoria*/
	CategoryResource struct {
		Data models.Category `json:"data"`
	}

	/*CategoriesResource para el array de categorias*/
	CategoriesResource struct {
		Data []models.Category `json:"data"`
	}

	/*MarcaResource para mostrar la marca*/
	MarcaResource struct {
		Data models.Marca `json:"data"`
	}

	/*MarcasResource para el array de marcas*/
	MarcasResource struct {
		Data []models.Marca `json:"data"`
	}
)
