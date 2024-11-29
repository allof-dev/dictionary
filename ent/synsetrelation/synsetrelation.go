// Code generated by ent, DO NOT EDIT.

package synsetrelation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the synsetrelation type in the database.
	Label = "synset_relation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRelType holds the string denoting the reltype field in the database.
	FieldRelType = "rel_type"
	// EdgeFrom holds the string denoting the from edge name in mutations.
	EdgeFrom = "from"
	// EdgeTo holds the string denoting the to edge name in mutations.
	EdgeTo = "to"
	// Table holds the table name of the synsetrelation in the database.
	Table = "synset_relations"
	// FromTable is the table that holds the from relation/edge.
	FromTable = "synset_relations"
	// FromInverseTable is the table name for the Synset entity.
	// It exists in this package in order to avoid circular dependency with the "synset" package.
	FromInverseTable = "synsets"
	// FromColumn is the table column denoting the from relation/edge.
	FromColumn = "synset_relation_from"
	// ToTable is the table that holds the to relation/edge.
	ToTable = "synset_relations"
	// ToInverseTable is the table name for the Synset entity.
	// It exists in this package in order to avoid circular dependency with the "synset" package.
	ToInverseTable = "synsets"
	// ToColumn is the table column denoting the to relation/edge.
	ToColumn = "synset_relation_to"
)

// Columns holds all SQL columns for synsetrelation fields.
var Columns = []string{
	FieldID,
	FieldRelType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "synset_relations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"synset_relation_from",
	"synset_relation_to",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// RelTypeValidator is a validator for the "relType" field. It is called by the builders before save.
	RelTypeValidator func(string) error
)

// OrderOption defines the ordering options for the SynsetRelation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRelType orders the results by the relType field.
func ByRelType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelType, opts...).ToFunc()
}

// ByFromField orders the results by from field.
func ByFromField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFromStep(), sql.OrderByField(field, opts...))
	}
}

// ByToField orders the results by to field.
func ByToField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToStep(), sql.OrderByField(field, opts...))
	}
}
func newFromStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FromInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, FromTable, FromColumn),
	)
}
func newToStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ToTable, ToColumn),
	)
}