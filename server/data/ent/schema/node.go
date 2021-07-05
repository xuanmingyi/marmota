package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Node struct {
	ent.Schema
}

func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("uuid"),
		field.String("name"),
		field.String("metadata"),
		field.String("desc"),
		field.Time("create_at"),
		field.Time("updated_at"),
	}
}

func (Node) Edges() []ent.Edge {
	return []ent.Edge{}
}
