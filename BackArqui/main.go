package main

import (
	"io"
	"log"
	"net/http"

	"./common"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	common.StartUp()
	http.HandleFunc("/hello", helloServer)
	server := &http.Server{
		Addr: common.AppConfig.Server,
	}
	log.Println("Listening...")
	server.ListenAndServe()
	defer common.GetSession().Close()
}
