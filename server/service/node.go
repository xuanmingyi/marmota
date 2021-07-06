package service

import (
	"context"
	"marmota/api"
	"marmota/server/biz"
)

func (s *NodeService) UpdateNode(ctx context.Context, req *api.UpdateNodeReq) (*api.UpdateNodeReply, error) {
	po, _ := s.nodeUseCase.FindNodeByUuid(ctx, req.Item.Uuid)

	if po == nil {
		// create
		po = &biz.Node{
			Name: req.Item.Name,
			Desc: req.Item.Desc,
			Uuid: req.Item.Uuid,
		}
		s.nodeUseCase.CreateNode(ctx, po)
	} else {
		// update
		po.Name = req.Item.Name
		po.Desc = req.Item.Desc
		s.nodeUseCase.UpdateNode(ctx, po)
	}

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
