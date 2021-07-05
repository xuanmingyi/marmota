package service

import (
	"context"
	"marmota/api"
)

func (s *NodeService) UpdateNode(ctx context.Context, req *api.UpdateNodeReq) (*api.UpdateNodeReply, error) {
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
