package service

import (
	"context"
	"fmt"
	"marmota/api"
)

func (s *NodeService) UpdateNode(ctx context.Context, req *api.UpdateNodeReq) (*api.UpdateNodeReply, error) {
	fmt.Println(req)
	return &api.UpdateNodeReply{
		Ret: 0,
		Msg: "ok",
	}, nil
}

func (s *NodeService) ListNode(ctx context.Context, req *api.ListNodeReq) (*api.ListNodeReply, error) {
	return &api.ListNodeReply{}, nil
}

func (s *NodeService) DeleteNode(ctx context.Context, req *api.DeleteNodeReq) (*api.DeleteNodeReply, error) {
	return &api.DeleteNodeReply{}, nil
}
