package biz

import "context"

type Node struct {
	Id   int64
	Name string
}

type NodeRepo interface {
	UpdateNode(ctx context.Context, n *Node) (*Node, error)
	ListNode(ctx context.Context) ([]*Node, error)
	DeleteNode(ctx context.Context, id int64) error
}

type NodeUseCase struct {
	repo NodeRepo
}

func NewNodeUseCase(repo NodeRepo) *NodeUseCase {
	return &NodeUseCase{repo: repo}
}

func (c *NodeUseCase) UpdateNode(ctx context.Context, n *Node) (*Node, error) {
	return c.repo.UpdateNode(ctx, n)
}

func (c *NodeUseCase) ListNode(ctx context.Context) ([]*Node, error) {
	return c.repo.ListNode(ctx)
}
