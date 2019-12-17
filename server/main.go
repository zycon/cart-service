package main

import (
	"context"
	"fmt"
	pbcart "github.com/zycon/cart-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	port = 1000
)

type cartServiceServer struct {
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Server Starting ..")
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("unable to listen on the port")
	}
	serverOptions := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOptions...)
	pbcart.RegisterCartServiceServer(grpcServer, &cartServiceServer{})
	reflection.Register(grpcServer)
	if e := grpcServer.Serve(lis); e != nil {
		panic(e)
	}

}

func (s *cartServiceServer) UpdateCart(ctx context.Context, request *pbcart.CartRequest) (*pbcart.CartResponse, error) {
	fmt.Println(request.GetCartId())
	return &pbcart.CartResponse{}, nil

}
func (s *cartServiceServer) StatusCheck(ctx context.Context, request *pbcart.StatusCheck) (*pbcart.StatusCheck, error) {
	return &pbcart.StatusCheck{Check: "welcome"}, nil
}
