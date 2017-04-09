package common

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

/*GetSession exporta la session de la base de datos*/
func GetSession() *sql.DB {
	if db == nil {
		createDBSession()
	}
	return db
}

func createDBSession() {
	var err error
	connection := getURLConnection()
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatalf("[Error open connection db]: %s\n", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("[Error ping to db]: %s\n", err)
		panic(err)
	}
}

func getURLConnection() string {
	return "user=" + AppConfig.PostgresDBUser + " password=" + AppConfig.PostgresDBPwd +
		" dbname=" + AppConfig.Database + " host=" + AppConfig.PostgresDBHost +
		" sslmode=disable"
}
