package main

import (
	"github.com/gorilla/handlers"
	"go_mongo/config"
	"go_mongo/route"
	"log"
	"net/http"
)

func main() {
	config.Load()
	router := route.InitRoute()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	AppPath := config.Global.SERVER_HOST + ":" + config.Global.SERVER_PORT

	log.Println("Server running: ", AppPath)
	err := http.ListenAndServe(AppPath, handlers.CORS(originsOk, headersOk, methodsOk)(router))
	if err != nil {
		log.Println("Can not start server, failed: ", err)
	}
}
