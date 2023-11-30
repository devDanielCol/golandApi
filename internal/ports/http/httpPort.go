package HttpPort

import (
	"log"
	"net/http"
	"time"

	httpAdapter "apiGoHttp/internal/adapters/http"
)

var httpAdapters *httpAdapter.Adapters

func startRouter() {
	httpAdapters = &httpAdapter.Adapters{}
	http.HandleFunc("/login", httpAdapters.Login)
}

func startServer() {
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

func InitApp() {
	startRouter()
	defer startServer()
}
