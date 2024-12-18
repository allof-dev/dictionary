// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/allof-dev/dictionary/ent/sense"
	"github.com/allof-dev/dictionary/ent/senserelation"
)

// SenseRelation is the model entity for the SenseRelation schema.
type SenseRelation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RelType holds the value of the "relType" field.
	RelType string `json:"relType,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SenseRelationQuery when eager-loading is set.
	Edges               SenseRelationEdges `json:"edges"`
	sense_relation_from *string
	sense_relation_to   *string
	selectValues        sql.SelectValues
}

// SenseRelationEdges holds the relations/edges for other nodes in the graph.
type SenseRelationEdges struct {
	// From holds the value of the from edge.
	From *Sense `json:"from,omitempty"`
	// To holds the value of the to edge.
	To *Sense `json:"to,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FromOrErr returns the From value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SenseRelationEdges) FromOrErr() (*Sense, error) {
	if e.From != nil {
		return e.From, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: sense.Label}
	}
	return nil, &NotLoadedError{edge: "from"}
}

// ToOrErr returns the To value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SenseRelationEdges) ToOrErr() (*Sense, error) {
	if e.To != nil {
		return e.To, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: sense.Label}
	}
	return nil, &NotLoadedError{edge: "to"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SenseRelation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case senserelation.FieldID:
			values[i] = new(sql.NullInt64)
		case senserelation.FieldRelType:
			values[i] = new(sql.NullString)
		case senserelation.ForeignKeys[0]: // sense_relation_from
			values[i] = new(sql.NullString)
		case senserelation.ForeignKeys[1]: // sense_relation_to
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SenseRelation fields.
func (sr *SenseRelation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case senserelation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sr.ID = int(value.Int64)
		case senserelation.FieldRelType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field relType", values[i])
			} else if value.Valid {
				sr.RelType = value.String
			}
		case senserelation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sense_relation_from", values[i])
			} else if value.Valid {
				sr.sense_relation_from = new(string)
				*sr.sense_relation_from = value.String
			}
		case senserelation.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sense_relation_to", values[i])
			} else if value.Valid {
				sr.sense_relation_to = new(string)
				*sr.sense_relation_to = value.String
			}
		default:
			sr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SenseRelation.
// This includes values selected through modifiers, order, etc.
func (sr *SenseRelation) Value(name string) (ent.Value, error) {
	return sr.selectValues.Get(name)
}

// QueryFrom queries the "from" edge of the SenseRelation entity.
func (sr *SenseRelation) QueryFrom() *SenseQuery {
	return NewSenseRelationClient(sr.config).QueryFrom(sr)
}

// QueryTo queries the "to" edge of the SenseRelation entity.
func (sr *SenseRelation) QueryTo() *SenseQuery {
	return NewSenseRelationClient(sr.config).QueryTo(sr)
}

// Update returns a builder for updating this SenseRelation.
// Note that you need to call SenseRelation.Unwrap() before calling this method if this SenseRelation
// was returned from a transaction, and the transaction was committed or rolled back.
func (sr *SenseRelation) Update() *SenseRelationUpdateOne {
	return NewSenseRelationClient(sr.config).UpdateOne(sr)
}

// Unwrap unwraps the SenseRelation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sr *SenseRelation) Unwrap() *SenseRelation {
	_tx, ok := sr.config.driver.(*txDriver)
	if !ok {
		panic("ent: SenseRelation is not a transactional entity")
	}
	sr.config.driver = _tx.drv
	return sr
}

// String implements the fmt.Stringer.
func (sr *SenseRelation) String() string {
	var builder strings.Builder
	builder.WriteString("SenseRelation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sr.ID))
	builder.WriteString("relType=")
	builder.WriteString(sr.RelType)
	builder.WriteByte(')')
	return builder.String()
}

// SenseRelations is a parsable slice of SenseRelation.
type SenseRelations []*SenseRelation
