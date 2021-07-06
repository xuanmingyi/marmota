package biz

import "context"

type Node struct {
	Id   int64
	Name string
	Uuid string
	Desc string
}

type NodeRepo interface {
	CreateNode(ctx context.Context, n *Node) (*Node, error)
	FindNodeByUuid(ctx context.Context, uuid string) (*Node, error)
	UpdateNode(ctx context.Context, n *Node) (*Node, error)
	DeleteNode(ctx context.Context, id int64) error
	ListNode(ctx context.Context) ([]*Node, error)
}

type NodeUseCase struct {
	repo NodeRepo
}

func NewNodeUseCase(repo NodeRepo) *NodeUseCase {
	return &NodeUseCase{repo: repo}
}

func (c *NodeUseCase) CreateNode(ctx context.Context, n *Node) (*Node, error) {
	return c.repo.CreateNode(ctx, n)
}

func (c *NodeUseCase) FindNodeByUuid(ctx context.Context, uuid string) (*Node, error) {
	return c.repo.FindNodeByUuid(ctx, uuid)
}

func (c *NodeUseCase) UpdateNode(ctx context.Context, n *Node) (*Node, error) {
	return c.repo.UpdateNode(ctx, n)
}

func (c *NodeUseCase) DeleteNode(ctx context.Context, id int64) error {
	return c.repo.DeleteNode(ctx, id)
}

func (c *NodeUseCase) ListNode(ctx context.Context) ([]*Node, error) {
	return c.repo.ListNode(ctx)
}
