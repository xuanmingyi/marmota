package data

import (
	"context"
	"fmt"
	"marmota/server/biz"
	"marmota/server/data/ent/node"
	"time"
)

type nodeRepo struct {
	data *Data
}

func NewNodeRepo(data *Data) biz.NodeRepo {
	return &nodeRepo{data: data}
}

func (n *nodeRepo) CreateNode(ctx context.Context, u *biz.Node) (*biz.Node, error) {
	po, err := n.data.db.Node.Create().
		SetName(u.Name).
		SetUUID(u.Uuid).
		SetDesc(u.Desc).
		SetMetadata("").
		SetCreateAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	fmt.Println(err)

	return &biz.Node{
		Desc: po.Desc,
		Uuid: po.UUID,
		Name: po.Name,
	}, err
}

func (n *nodeRepo) FindNodeByUuid(ctx context.Context, uuid string) (*biz.Node, error) {
	po, err := n.data.db.Node.Query().Where(node.UUIDEQ(uuid)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Node{
		Id:   po.ID,
		Uuid: po.UUID,
		Name: po.Name,
		Desc: po.Desc,
	}, nil
}

func (n *nodeRepo) UpdateNode(ctx context.Context, u *biz.Node) (*biz.Node, error) {
	po, err := n.data.db.Node.UpdateOneID(u.Id).
		SetName(u.Name).
		SetDesc(u.Desc).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return &biz.Node{
		Id:   po.ID,
		Uuid: po.UUID,
		Name: po.Name,
		Desc: po.Desc,
	}, nil
}

func (n *nodeRepo) DeleteNode(ctx context.Context, id int64) error {
	return nil
}

func (n *nodeRepo) ListNode(ctx context.Context) ([]*biz.Node, error) {
	return nil, nil
}
