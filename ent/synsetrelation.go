// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/allof-dev/dictionary/ent/synset"
	"github.com/allof-dev/dictionary/ent/synsetrelation"
)

// SynsetRelation is the model entity for the SynsetRelation schema.
type SynsetRelation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RelType holds the value of the "relType" field.
	RelType string `json:"relType,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SynsetRelationQuery when eager-loading is set.
	Edges                SynsetRelationEdges `json:"edges"`
	synset_relation_from *string
	synset_relation_to   *string
	selectValues         sql.SelectValues
}

// SynsetRelationEdges holds the relations/edges for other nodes in the graph.
type SynsetRelationEdges struct {
	// From holds the value of the from edge.
	From *Synset `json:"from,omitempty"`
	// To holds the value of the to edge.
	To *Synset `json:"to,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FromOrErr returns the From value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SynsetRelationEdges) FromOrErr() (*Synset, error) {
	if e.From != nil {
		return e.From, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: synset.Label}
	}
	return nil, &NotLoadedError{edge: "from"}
}

// ToOrErr returns the To value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SynsetRelationEdges) ToOrErr() (*Synset, error) {
	if e.To != nil {
		return e.To, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: synset.Label}
	}
	return nil, &NotLoadedError{edge: "to"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SynsetRelation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case synsetrelation.FieldID:
			values[i] = new(sql.NullInt64)
		case synsetrelation.FieldRelType:
			values[i] = new(sql.NullString)
		case synsetrelation.ForeignKeys[0]: // synset_relation_from
			values[i] = new(sql.NullString)
		case synsetrelation.ForeignKeys[1]: // synset_relation_to
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SynsetRelation fields.
func (sr *SynsetRelation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case synsetrelation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sr.ID = int(value.Int64)
		case synsetrelation.FieldRelType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field relType", values[i])
			} else if value.Valid {
				sr.RelType = value.String
			}
		case synsetrelation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field synset_relation_from", values[i])
			} else if value.Valid {
				sr.synset_relation_from = new(string)
				*sr.synset_relation_from = value.String
			}
		case synsetrelation.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field synset_relation_to", values[i])
			} else if value.Valid {
				sr.synset_relation_to = new(string)
				*sr.synset_relation_to = value.String
			}
		default:
			sr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SynsetRelation.
// This includes values selected through modifiers, order, etc.
func (sr *SynsetRelation) Value(name string) (ent.Value, error) {
	return sr.selectValues.Get(name)
}

// QueryFrom queries the "from" edge of the SynsetRelation entity.
func (sr *SynsetRelation) QueryFrom() *SynsetQuery {
	return NewSynsetRelationClient(sr.config).QueryFrom(sr)
}

// QueryTo queries the "to" edge of the SynsetRelation entity.
func (sr *SynsetRelation) QueryTo() *SynsetQuery {
	return NewSynsetRelationClient(sr.config).QueryTo(sr)
}

// Update returns a builder for updating this SynsetRelation.
// Note that you need to call SynsetRelation.Unwrap() before calling this method if this SynsetRelation
// was returned from a transaction, and the transaction was committed or rolled back.
func (sr *SynsetRelation) Update() *SynsetRelationUpdateOne {
	return NewSynsetRelationClient(sr.config).UpdateOne(sr)
}

// Unwrap unwraps the SynsetRelation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sr *SynsetRelation) Unwrap() *SynsetRelation {
	_tx, ok := sr.config.driver.(*txDriver)
	if !ok {
		panic("ent: SynsetRelation is not a transactional entity")
	}
	sr.config.driver = _tx.drv
	return sr
}

// String implements the fmt.Stringer.
func (sr *SynsetRelation) String() string {
	var builder strings.Builder
	builder.WriteString("SynsetRelation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sr.ID))
	builder.WriteString("relType=")
	builder.WriteString(sr.RelType)
	builder.WriteByte(')')
	return builder.String()
}

// SynsetRelations is a parsable slice of SynsetRelation.
type SynsetRelations []*SynsetRelation
