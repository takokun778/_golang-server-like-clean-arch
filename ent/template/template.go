// Code generated by entc, DO NOT EDIT.

package template

import (
	"time"
)

const (
	// Label holds the string label denoting the template type in the database.
	Label = "template"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSample holds the string denoting the sample field in the database.
	FieldSample = "sample"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the template in the database.
	Table = "templates"
)

// Columns holds all SQL columns for template fields.
var Columns = []string{
	FieldID,
	FieldSample,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)
