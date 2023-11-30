package api

import (
	"log"
	"net/http"
	"time"

	"apiGoHttp/pkg/adapters"
)

var httpAdapters *adapters.Adapters

func StartRouter() {
	httpAdapters = &adapters.Adapters{}
	http.HandleFunc("/login", httpAdapters.Login)
}

func StartServer() {
	server := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server Initialized")
	log.Fatal(server.ListenAndServe())
}
