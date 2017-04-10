package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"./common"
	"./routes"
)

func main() {
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	common.StartUp()
	router := routes.InitRoutes()
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: handlers.CORS(headersOk, methodsOk)(router),
	}
	log.Println("Listening...")
	server.ListenAndServe()
	defer common.GetSession().Close()
}
