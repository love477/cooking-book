package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/love477/cooking_book/app/api"
	"google.golang.org/grpc"
	"net"

	"github.com/love477/cooking_book/proto/menu"
)

func RunGrpcServer(ctx context.Context, port string, done func()) {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		fmt.Errorf("start grpc server error: %v", err)
		panic(err)
	}
	//var ops []grpc.ServerOption
	grpcServer := grpc.NewServer()

	menu.RegisterMenuServer(grpcServer, &api.MenuServer{})
	grpcServer.Serve(listener)
}
