package main

import (
	"log"
	"os"

	"github.com/fsufitch/r9kd/server"
)

func main() {
	port := os.Getenv("R9KD_PORT")
	if port == "" {
		log.Fatal("R9KD_PORT environment var missing")
	}
	log.Fatal(server.RunServer(port))
}
