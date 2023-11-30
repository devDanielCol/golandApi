package main

import (
	"apiGoHttp/pkg/api"
)

func main() {
	api.StartRouter()
	defer api.StartServer()
}
