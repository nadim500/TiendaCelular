package data

import (
	"errors"
	"log"
	"strconv"

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

/*GetCategoryByID nos retorna una categoria por su ID*/
func GetCategoryByID(id string) (models.Category, error) {
	category := models.Category{}
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[Error convert id to int] %s\n", err)
		return category, err
	}
	db := common.GetSession()
	err = db.QueryRow(
		"SELECT id_category, name, description FROM category WHERE id_category=$1",
		i,
	).Scan(&category.IDCategory, &category.Name, &category.Description)
	return category, err
}

/*CreateCategory crea en la base de datos una categoría*/
func CreateCategory(c *models.Category) error {
	var id int
	db := common.GetSession()
	err := db.QueryRow(
		"INSERT INTO category(name,description) VALUES($1,$2) returning id_category;",
		c.Name,
		c.Description,
	).Scan(&id)
	if err != nil {
		log.Printf("[Error insert category]: %s\n", err)
		return err
	}
	c.IDCategory = id
	return err
}

/*UpdateCategory actualiza una categoria en la base de datos*/
func UpdateCategory(c *models.Category, id string) error {
	name := len(c.Name)
	descr := len(c.Description)
	query := "UPDATE category SET "
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[Error convert id to int] %s\n", err)
		return err
	}
	db := common.GetSession()
	switch {
	case name == 0 && descr == 0:
		log.Println("[Sin datos para poder actualizar]")
		err = errors.New("Sin datos para poder actualizar")
		return err
	case name != 0 && descr == 0:
		query += "name = $1 WHERE id_category = $2"
		query += " RETURNING id_category, name, description"
		err = db.QueryRow(query, c.Name, i).Scan(&c.IDCategory, &c.Name, &c.Description)
	case name == 0 && descr != 0:
		query += "description = $1 WHERE id_category = $2"
		query += " RETURNING id_category, name, description"
		err = db.QueryRow(query, c.Description, i).Scan(&c.IDCategory, &c.Name, &c.Description)
	case name != 0 && descr != 0:
		query += "name = $1, description = $2 WHERE id_category = $3"
		query += " RETURNING id_category, name, description"
		err = db.QueryRow(query, c.Name, c.Description, i).Scan(&c.IDCategory, &c.Name, &c.Description)
	}
	return err
}

/*DeleteCategory borra una categoria de la base de datos */
func DeleteCategory(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[Error convert id to int]: %s\n", err)
		return err
	}
	db := common.GetSession()
	stmt, err := db.Prepare("DELETE FROM category WHERE id_category=$1")
	if err != nil {
		log.Printf("[Error en prepare query]: %s\n", err)
		return err
	}
	_, err = stmt.Exec(i)
	if err != nil {
		log.Printf("[Error exec query consult]: %s\n", err)
		return err
	}
	return err
}
