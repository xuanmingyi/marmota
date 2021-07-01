package schema

import "entgo.io/ent"

type NodeMetadata struct {
	ent.Schema
}

func (NodeMetadata) Fields() []ent.Field {
	return []ent.Field{}
}

func (NodeMetadata) Edges() []ent.Edge {
	return []ent.Edge{}
}
