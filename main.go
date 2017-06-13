package main

import (
	"log"

	"github.com/fsufitch/r9kd/server"
)

func main() {
	log.Fatal(server.RunServer(8080))
}
