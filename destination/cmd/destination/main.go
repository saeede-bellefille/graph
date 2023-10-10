package main

import (
	"graph/destination/pkg/processor"
	"graph/socket"
	"log"
)

func main() {
	p := processor.New()
	server := socket.NewServer(":8082", p.Process)
	log.Panic(server.Start())
}
