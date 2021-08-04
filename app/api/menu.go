package api

import (
	"context"
	"fmt"
	"github.com/love477/cooking_book/proto/menu"
)

type MenuServer struct {
	menu.UnimplementedMenuServer
}

func (menuServer * MenuServer) CreateMenu(ctx context.Context, request *menu.CreateMenuRequest) (*menu.CreateMenuResponse, error) {
	fmt.Println("get request from client for create menu")
	return &menu.CreateMenuResponse{}, nil
}