package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=test host=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	query, err := db.Prepare("SELECT last_name FROM actor")
	if err != nil {
		log.Fatal("error prepare query")
	}

	rows, err := query.Query()
	defer rows.Close()
	var names []string

	for rows.Next() {
		var name string
		err1 := rows.Scan(&name)
		if err1 != nil {
			log.Fatal(err1)
		}
		names = append(names, name)
	}
	log.Print(names)
	log.Print("Conecto")

}
