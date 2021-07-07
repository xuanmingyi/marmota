package main

import (
	"flag"
	"fmt"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"

	"marmota/server/biz"
	"marmota/server/conf"
	"marmota/server/data"
	"marmota/server/server"
	"marmota/server/service"

	"net"

	"google.golang.org/grpc"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs.yaml", "config path, eg: -conf config.yaml")
}

func NewApp(bc *conf.Bootstrap) (*grpc.Server, func(), error) {
	fmt.Println(bc.Data.Database.Driver)

	d, cleanup, err := data.NewData("mysql", "root:123456@tcp(127.0.0.1:3306)/newdb1?parseTime=true")
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
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap

	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	l, _ := net.Listen("tcp", ":50051")

	app, cleanup, err := NewApp(&bc)

	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Serve(l); err != nil {
		panic(err)
	}
}
