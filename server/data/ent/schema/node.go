package schema

import (
	"entgo.io/ent"
)

type Node struct {
	ent.Schema
}

func (Node) Fields() []ent.Field {
	return []ent.Field{}
}

func (Node) Edges() []ent.Edge {
	return []ent.Edge{}
}
