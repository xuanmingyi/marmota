package main

import (
	"context"
	"marmota/api"
	"time"

	"google.golang.org/grpc"
)

type agent struct {
	Conn *grpc.ClientConn
	Node api.NodeClient
	Conf *conf
}

func NewClient(c *conf) (a *agent, err error) {
	a = new(agent)
	a.Conf = c
	a.Conn, err = grpc.Dial(a.Conf.Server.Uri, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *agent) UpdateNode(req *api.UpdateNodeReq) (*api.UpdateNodeReply, error) {
	if a.Node == nil {
		a.Node = api.NewNodeClient(a.Conn)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	return a.Node.UpdateNode(ctx, req)
}
