package main

import (
	"context"
	"fmt"
	"github.com/love477/cooking_book/proto/menu"
	"google.golang.org/grpc"
)

const ServerAddr = "localhost:3001"


func  main()  {
	ops := grpc.WithInsecure()
	conn, err := grpc.Dial(ServerAddr, ops)

	if err != nil {
		fmt.Println("connect to srpc server error: ", err)
	}
	defer conn.Close()

	client := menu.NewMenuClient(conn)
	fmt.Println("CreateMenu")
	client.CreateMenu(context.Background(), &menu.CreateMenuRequest{})
}
