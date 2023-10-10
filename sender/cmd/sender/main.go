package main

import (
	"graph/sender/pkg/generator"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s count\n", os.Args[0])
	}
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Invalid count", err)
	}

	g := generator.New(50, 8*1024, "http://127.0.0.1:8080")
	for i := 0; i < count; i++ {
		if err := g.Send(); err != nil {
			log.Fatal("Error sending", err)
		}
	}
	log.Printf("Sent %d messages!\n", count)
}
