package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsufitch/r9kd/server"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Port argument not specified")
	}
	port := os.Args[1]
	fmt.Printf("Starting r9kd server on port %s\n", port)
	log.Fatal(server.RunServer(port))
}
