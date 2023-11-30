package main

import (
	HttpPort "apiGoHttp/internal/ports/http"
)

func main() {
	defer HttpPort.InitApp()
}
