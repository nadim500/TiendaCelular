package data

import (
	"log"

	"../common"
	"../models"
)

/*GetCategories es la conexión hacia la base de datos*/
func GetCategories() ([]models.Category, error) {
	categories := make([]models.Category, 0)
	db := common.GetSession()
	stmt, err := db.Prepare("SELECT * FROM category")
	if err != nil {
		log.Printf("[Error en prepare query]: %s\n", err)
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("[Error en consult query]: %s\n", err)
		return nil, err
	}
	defer rows.Close()
	result := models.Category{}
	for rows.Next() {
		err = rows.Scan(&result.IDCategory, &result.Name, &result.Description)
		categories = append(categories, result)
	}
	if err != nil {
		log.Printf("[Error Scan get categories]: %s\n", err)
		return nil, err
	}
	return categories, err
}
