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

var db *sql.DB

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

	// mysql

	var (
		id       int
		user_id  int
		modified string
	)
	connectDb(db)
	if db != nil {
		result, err := db.Query("SELECT id,user_id,modified FROM cart where is_deleted=0;")
		if err != nil {
			log.Fatal(err)
		}
		for result.Next() {
			result.Scan(&id, &user_id, &modified)
			fmt.Println(id, user_id, modified)
		}
	}

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
func (s *cartServiceServer) GetCarts(ctx context.Context, request *pbcart.Empty) (*pbcart.CartListResponse, error) {

	return &pbcart.CartListResponse{}, nil
}

func connectDb(db sql.DB) {
	db, err := sql.Open("mysql", "root:ad#12345@/cart")
	if err != nil {
		log.Fatal("Error while opening the connection")
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Connected successfully")
	}

}
