package controllers

import "../models"

type (

	/*CategoriesResource para el array de categorias*/
	CategoriesResource struct {
		Data []models.Category `json:"data"`
	}
)
