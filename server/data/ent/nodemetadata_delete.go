// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"marmota/server/data/ent/nodemetadata"
	"marmota/server/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NodeMetadataDelete is the builder for deleting a NodeMetadata entity.
type NodeMetadataDelete struct {
	config
	hooks    []Hook
	mutation *NodeMetadataMutation
}

// Where adds a new predicate to the NodeMetadataDelete builder.
func (nmd *NodeMetadataDelete) Where(ps ...predicate.NodeMetadata) *NodeMetadataDelete {
	nmd.mutation.predicates = append(nmd.mutation.predicates, ps...)
	return nmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nmd *NodeMetadataDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nmd.hooks) == 0 {
		affected, err = nmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodeMetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nmd.mutation = mutation
			affected, err = nmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nmd.hooks) - 1; i >= 0; i-- {
			mut = nmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nmd *NodeMetadataDelete) ExecX(ctx context.Context) int {
	n, err := nmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nmd *NodeMetadataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: nodemetadata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nodemetadata.FieldID,
			},
		},
	}
	if ps := nmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, nmd.driver, _spec)
}

// NodeMetadataDeleteOne is the builder for deleting a single NodeMetadata entity.
type NodeMetadataDeleteOne struct {
	nmd *NodeMetadataDelete
}

// Exec executes the deletion query.
func (nmdo *NodeMetadataDeleteOne) Exec(ctx context.Context) error {
	n, err := nmdo.nmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nodemetadata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nmdo *NodeMetadataDeleteOne) ExecX(ctx context.Context) {
	nmdo.nmd.ExecX(ctx)
}
