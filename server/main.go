package main

import (
	cart "../pb/"
	"log"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

var (
	port = 1000;
)

type CartServiceServer struct {
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Server Starting ..")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10000))
	if err != nil {
		log.Fatal("unable to listen on the port")
	}
	serverOptions := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOptions...)
	srv := &CartServiceServer{}
	pbcart.RegisterCartServiceServer(grpcServer, srv)
}
