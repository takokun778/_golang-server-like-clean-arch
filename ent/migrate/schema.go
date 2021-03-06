// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TemplatesColumns holds the columns for the "templates" table.
	TemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "sample", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TemplatesTable holds the schema information for the "templates" table.
	TemplatesTable = &schema.Table{
		Name:       "templates",
		Columns:    TemplatesColumns,
		PrimaryKey: []*schema.Column{TemplatesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TemplatesTable,
	}
)

func init() {
}
