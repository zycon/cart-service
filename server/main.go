package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	pbcart "github.com/zycon/cart-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

	db, err := sql.Open("mysql", "root:ad#1234@/cart")
	if err != nil {
		log.Fatal("unable to establish mysql connection")
	}
	if db != nil {
		log.Println("Connection established")
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
	return &pbcart.CartResponse{CartId: request.CartId}, nil

}
func (s *cartServiceServer) StatusCheck(ctx context.Context, request *pbcart.StatusCheck) (*pbcart.StatusCheck, error) {
	return &pbcart.StatusCheck{Check: "welcome"}, nil
}
func (s *cartServiceServer) GetCarts(ctx context.Context, request *Empty) (*pbcart.CartListResponse, error) {
	return &pbcart.CartListResponse{}, nil
}
