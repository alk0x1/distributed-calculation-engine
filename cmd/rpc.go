package main

import (
	"flag"
	"fmt"
	"matrix"
	"net"
	"net/rpc"
)

func main() {
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	rpc.Register(new(matrix.MatrixServiceImpl))

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Println("Server is starting on port", *port)
	rpc.Accept(listener)
}
