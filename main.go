package main

import (
	"context"
	"fmt"

	"github.com/love477/cooking_book/app/server"
)

const PORT = "3001"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("grpc server listen on port " + PORT)
	go server.RunGrpcServer(ctx, PORT, func() {
		fmt.Println("grpc server listen on port " + PORT)
		cancel()
	})

	<-ctx.Done()
}