package main

import (
	"marmota/server/biz"
	"marmota/server/data"
	"marmota/server/server"
	"marmota/server/service"

	"net"

	"google.golang.org/grpc"
)

func NewApp() (*grpc.Server, func(), error) {
	d, cleanup, err := data.NewData()
	if err != nil {
		panic(err)
	}

	nodeRepo := data.NewNodeRepo(d)
	nodeUseCase := biz.NewNodeUseCase(nodeRepo)
	nodeService := service.NewNodeService(nodeUseCase)
	grpcServer := server.NewGRPCServer(nodeService)
	return grpcServer, cleanup, nil
}

func main() {
	l, _ := net.Listen("tcp", ":50051")

	app, cleanup, err := NewApp()

	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Serve(l); err != nil {
		panic(err)
	}
}
