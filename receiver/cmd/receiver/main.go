package main

import (
	"graph/receiver/pkg/server"
	"log"
)

func main() {
	server := server.New(":8080", "127.0.0.1:8081")
	log.Fatal(server.Start())
}
