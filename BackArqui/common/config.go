package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Server, PostgresDBHost, PostgresDBUser, PostgresDBPwd, Database string
}

/*AppConfig contiene la configuración de la aplicación ...*/
var AppConfig configuration

func initConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("./config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[Error open file config.json]: %s\n", err)
		panic(err)
	}
	AppConfig = configuration{}
	err = json.NewDecoder(file).Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[Error Decode file config.json]: %s\n", err)
		panic(err)
	}
}
