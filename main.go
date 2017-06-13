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
	fmt.Println("hello world")
	log.Fatal(server.RunServer(port))
}
