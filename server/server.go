package main

import (
	protopkg "./proto"
	service "./service"
	"fmt"
	grpc "google.golang.org/grpc"
	"log"
	"net"
)

func checkForErr(msg string, err error) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func main() {
	li, err := net.Listen("tcp", ":9090")
	checkForErr("Erro creating server: %v", err)
	s := &service.Server{}
	server := grpc.NewServer()
	protopkg.RegisterServiceServer(server, s)
	fmt.Println("Servindo em porta 9090")
	server.Serve(li)
}
