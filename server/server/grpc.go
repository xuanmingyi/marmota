package server

import (
	"marmota/api"
	"marmota/server/service"

	"google.golang.org/grpc"
)

func NewGRPCServer(ns *service.NodeService) *grpc.Server {
	var opts = []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)

	api.RegisterNodeServer(srv, ns)

	return srv
}
