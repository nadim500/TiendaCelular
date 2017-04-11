package models

import "time"

type (
	/*Category model from category table*/
	Category struct {
		IDCategory  int    `json:"id_category,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}

	/*Marca model from model table*/
	Marca struct {
		IDMarca     int    `json:"id_marca"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       string `json:"image"`
	}

	/*Product model from model product*/
	Product struct {
		IDProduct         int       `json:"id_product"`
		Name              string    `json:"name"`
		FirstDescription  string    `json:"first_description"`
		SecondDescription string    `json:"second_description"`
		Feature           string    `json:"feature"`
		DateCreated       time.Time `json:"datecreated"`
		Images            []string  `json:"images"`
		MarcaID           int       `json:"marca_id"`
		CategoryID        int       `json:"category_id"`
	}

	/*Client model from client table*/
	Client struct {
		IDClient      int    `json:"id_client"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		Phone         string `json:"phone"`
		Direction     string `json:"direction"`
		Document      string `json:"document"`
		EmailVerified bool   `json:"email_verified"`
		Password      string `json:"string"`
		HashPassword  []byte `json:"hash_password"`
	}

	/*Annotation model from annotation table*/
	Annotation struct {
		ClientID    int       `json:"client_id"`
		ProductID   int       `json:"product_id"`
		DateCreated time.Time `json:"datecreated"`
		Annotation  string    `json:"annotation"`
	}

	/*Bill model from bill table*/
	Bill struct {
		IDBill      int       `json:"id_bill"`
		DateCreated time.Time `json:"datecreated"`
		Mount       float32   `json:"mount"`
		ClientID    int       `json:"client_id"`
	}

	/*ProductHasBill from product_has_bill table*/
	ProductHasBill struct {
		ProductID int     `json:"product_id"`
		BillID    int     `json:"bill_id"`
		Mount     float32 `json:"mount"`
		Quantity  int     `json:"quantity"`
	}

	/*Connectivity from connectivity table*/
	Connectivity struct {
		IDConnectivity int    `json:"id_connectivity"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		Image          string `json:"image"`
	}

	/*ProductHasConnectivity from product_has_connectivity*/
	ProductHasConnectivity struct {
		ProductID      int `json:"product_id"`
		ConnectivityID int `json:"connectivity_id"`
	}
)
