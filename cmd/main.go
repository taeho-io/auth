package main

import (
	"fmt"

	"github.com/taeho-io/auth/server"
)

func main() {
	fmt.Println("Starting Auth gRPC/HTTP server...")
	err := server.Serve()
	fmt.Println(err)
}
