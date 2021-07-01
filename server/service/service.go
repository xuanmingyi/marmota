package service

import (
	"marmota/api"
	"marmota/server/biz"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewNodeService)

type NodeService struct {
	api.UnimplementedNodeServer

	nodeUseCase *biz.NodeUseCase
}

func NewNodeService(nodeUseCase *biz.NodeUseCase) *NodeService {
	return &NodeService{nodeUseCase: nodeUseCase}
}
