package data

import (
	"errors"
	"log"
	"strconv"

	"../common"
	"../models"
)

/*GetMarcas consulta todas las marcas en la base de datos*/
func GetMarcas() ([]models.Marca, error) {
	marcas := make([]models.Marca, 0)
	db := common.GetSession()
	stmt, err := db.Prepare("SELECT * FROM marca")
	if err != nil {
		log.Printf("[Error en prepare query]: %s\n", err)
		return marcas, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("[Error en consult query]: %s\n", err)
		return marcas, err
	}
	defer rows.Close()
	result := models.Marca{}
	for rows.Next() {
		err = rows.Scan(&result.IDMarca, &result.Name, &result.Description, &result.Image)
		marcas = append(marcas, result)
	}
	if err != nil {
		log.Printf("[Error scan get marcas]: %s\n", err)
		return marcas, err
	}
	return marcas, err
}

/*GetMarcaByID retorna una marca por su ID*/
func GetMarcaByID(id string) (models.Marca, error) {
	result := models.Marca{}
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[Error convert id to int] %s\n", err)
		return result, err
	}
	db := common.GetSession()
	err = db.QueryRow(
		"SELECT * FROM marca WHERE id_marca = $1",
		i,
	).Scan(&result.IDMarca, &result.Name, &result.Description, &result.Image)
	if err != nil {
		log.Printf("[Error scan get marca by id]: %s\n", err)
	}
	return result, err
}

/*CreateMarca crea una marca*/
func CreateMarca(m *models.Marca) error {
	var id int
	db := common.GetSession()
	err := db.QueryRow(
		"INSERT INTO marca(name,description,image) VALUES($1,$2,$3) returning id_marca",
		m.Name,
		m.Description,
		m.Image,
	).Scan(&id)
	if err != nil {
		log.Printf("[Error insert marca]: %s\n", err)
		return err
	}
	m.IDMarca = id
	return err
}

/*UpdateMarca actualiza una categoria en la base de datos*/
func UpdateMarca(m *models.Marca, id string) error {
	name := len(m.Name)
	descr := len(m.Description)
	image := len(m.Image)
	query := "UPDATE marca SET "
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[Error convert id to int] %s\n", err)
		return err
	}
	db := common.GetSession()
	switch {
	case name == 0:
		switch {
		case descr != 0:
			switch {
			case image != 0:
				query += "description = $1, image = $2 WHERE id_marca = $3"
				query += " RETURNING id_marca, name, description, image"
				err = db.QueryRow(query, m.Description, m.Image, i).Scan(&m.IDMarca, &m.Name, &m.Description, &m.Image)
			case image == 0:
				query += "description = $1 WHERE id_marca = $2"
				query += " RETURNING id_marca, name, description, image"
				err = db.QueryRow(query, m.Description, i).Scan(&m.IDMarca, &m.Name, &m.Description, &m.Image)
			}
		case descr == 0:
			switch {
			case image != 0:
				query += "image = $1 WHERE id_marca = $2"
				query += " RETURNING id_marca, name, description, image"
				err = db.QueryRow(query, m.Image, i).Scan(&m.IDMarca, &m.Name, &m.Description, &m.Image)
			case image == 0:
				log.Println("[Sin datos para poder actualizar]")
				err = errors.New("Sin datos para poder actualizar")
			}
		}
	case name != 0:
		switch {
		case descr != 0:
			switch {
			case image != 0:
				query += "name = $1, description = $2, image = $3 WHERE id_marca = $4"
				query += " RETURNING id_marca, name, description, image"
				err = db.QueryRow(query, m.Name, m.Description, m.Image, i).Scan(&m.IDMarca, &m.Name, &m.Description, &m.Image)
			case image == 0:
				query += "name = $1, description = $2 WHERE id_marca = $3"
				query += " RETURNING id_marca, name, description, image"
				err = db.QueryRow(query, m.Name, m.Description, i).Scan(&m.IDMarca, &m.Name, &m.Description, &m.Image)
			}
		case descr == 0:
			switch {
			case image != 0:
				query += "name = $1, image = $2 WHERE id_marca = $3"
				query += " RETURNING id_marca, name, description, image"
				err = db.QueryRow(query, m.Name, m.Image, i).Scan(&m.IDMarca, &m.Name, &m.Description, &m.Image)
			case image == 0:
				query += "name = $1 WHERE id_marca = $2"
				query += " RETURNING id_marca, name, description, image"
				err = db.QueryRow(query, m.Name, i).Scan(&m.IDMarca, &m.Name, &m.Description, &m.Image)
			}
		}
	}
	return err
}

/*DeleteMarca borra una categoria de la base de datos */
func DeleteMarca(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[Error convert id to int]: %s\n", err)
		return err
	}
	db := common.GetSession()
	stmt, err := db.Prepare("DELETE FROM marca WHERE id_marca=$1")
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
