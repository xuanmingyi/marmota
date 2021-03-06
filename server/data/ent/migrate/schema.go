// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NodesColumns holds the columns for the "nodes" table.
	NodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "uuid", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// NodesTable holds the schema information for the "nodes" table.
	NodesTable = &schema.Table{
		Name:        "nodes",
		Columns:     NodesColumns,
		PrimaryKey:  []*schema.Column{NodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// NodeMetadataColumns holds the columns for the "node_metadata" table.
	NodeMetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// NodeMetadataTable holds the schema information for the "node_metadata" table.
	NodeMetadataTable = &schema.Table{
		Name:        "node_metadata",
		Columns:     NodeMetadataColumns,
		PrimaryKey:  []*schema.Column{NodeMetadataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NodesTable,
		NodeMetadataTable,
	}
)

func init() {
}
