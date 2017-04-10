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
)
