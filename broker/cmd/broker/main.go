package main

import (
	"graph/broker/pkg/handler"
	"graph/socket"
	"log"
)

func main() {
	h := handler.New("127.0.0.1:8082")
	server := socket.NewServer(":8081", h.Handle)
	log.Fatal(server.Start())
}
