package main

import (
	"fmt"
	"os"

	"github.com/poc/cmd/client"
	"github.com/poc/cmd/server"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usgae: go run main.go [client|server]")
		return
	}

	mode := os.Args[1]

	switch mode {
	case "client":
		client.Init()
	case "server":
		server.Init()
	default:
		fmt.Println("Invalid mode. Use 'client' or 'server'.")
	}
}
