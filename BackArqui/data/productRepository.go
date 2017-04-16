package data

import (
	"log"
	"strconv"

	"github.com/fatih/structs"
	"github.com/lib/pq"

	"../common"
	"../models"
)

/*GetProducts consulta todos los productos en la base de datos*/
func GetProducts() ([]models.Product, error) {
	productos := make([]models.Product, 0)
	db := common.GetSession()
	stmt, err := db.Prepare("SELECT * FROM product")
	if err != nil {
		log.Printf("[Error en prepare query]: %s\n", err)
		return productos, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("[Error en consult query]: %s\n", err)
		return productos, err
	}
	defer rows.Close()
	result := models.Product{}
	for rows.Next() {
		err = rows.Scan(
			&result.IDProduct,
			&result.Name,
			&result.FirstDescription,
			&result.SecondDescription,
			&result.Feature,
			&result.DateCreated,
			pq.Array(&result.Images),
			&result.Stock,
			&result.MarcaID,
			&result.CategoryID,
		)
		productos = append(productos, result)
	}
	if err != nil {
		log.Printf("[Error scan get productos]: %s\n", err)
		return productos, err
	}
	return productos, err
}

/*GetProductByID nos retorna un producto por su ID*/
func GetProductByID(id string) (models.Product, error) {
	product := models.Product{}
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[Error convert id to int]: %s\n", err)
		return product, err
	}
	db := common.GetSession()
	err = db.QueryRow("SELECT * FROM product WHERE id_product = $1", i).Scan(
		&product.IDProduct,
		&product.Name,
		&product.FirstDescription,
		&product.SecondDescription,
		&product.Feature,
		&product.DateCreated,
		pq.Array(&product.Images),
		&product.Stock,
		&product.MarcaID,
		&product.CategoryID,
	)
	return product, err
}

/*CreateProduct crea en la base de datos un producto*/
func CreateProduct(p *models.Product) error {
	var id int
	db := common.GetSession()
	err := db.QueryRow(
		"INSERT INTO product(name,first_description,second_description,feature,images,stock,marca_id,category_id) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id_product",
		p.Name,
		p.FirstDescription,
		p.SecondDescription,
		p.Feature,
		pq.Array(p.Images),
		p.Stock,
		p.MarcaID,
		p.CategoryID,
	).Scan(&id)
	if err != nil {
		log.Printf("[Error insert product]: %s\n", err)
		return err
	}
	p.IDProduct = id
	return err
}

/*UpdateProduct actualiza un producto en la base de datos*/
func UpdateProduct(p *models.Product, id string) error {
	log.Println(p)
	v := structs.HasZero(p)
	log.Println(v)

	/*i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[Error convert id to int] %s\n", err)
		return err
	}*/
	/*query := "UPDATE product SET "
	name := len(p.Name)
	first := len(p.FirstDescription)
	second := len(p.SecondDescription)
	feature := len(p.Feature)
	images := len(p.Images)
	db := common.GetSession()*/
	return nil
}

/*DeleteProduct borrar un producto de la base de datos*/
func DeleteProduct(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[Error convert id to int]: %s\n", err)
		return err
	}
	db := common.GetSession()
	stmt, err := db.Prepare("DELETE FROM product WHERE id_product = $1")
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
