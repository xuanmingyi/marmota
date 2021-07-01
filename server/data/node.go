package data

import (
	"context"
	"marmota/server/biz"
)

type nodeRepo struct {
	data *Data
}

func NewNodeRepo(data *Data) biz.NodeRepo {
	return &nodeRepo{data: data}
}

func (n *nodeRepo) UpdateNode(ctx context.Context, u *biz.Node) (*biz.Node, error) {
	return nil, nil
}

func (n *nodeRepo) DeleteNode(ctx context.Context, id int64) error {
	return nil
}

func (n *nodeRepo) ListNode(ctx context.Context) ([]*biz.Node, error) {
	return nil, nil
}
