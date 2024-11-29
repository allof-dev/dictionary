// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/allof-dev/dictionary/ent/synset"
)

// Synset is the model entity for the Synset schema.
type Synset struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// PartOfSpeech holds the value of the "partOfSpeech" field.
	PartOfSpeech string `json:"partOfSpeech,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SynsetQuery when eager-loading is set.
	Edges        SynsetEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SynsetEdges holds the relations/edges for other nodes in the graph.
type SynsetEdges struct {
	// Definitions holds the value of the definitions edge.
	Definitions []*Definition `json:"definitions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DefinitionsOrErr returns the Definitions value or an error if the edge
// was not loaded in eager-loading.
func (e SynsetEdges) DefinitionsOrErr() ([]*Definition, error) {
	if e.loadedTypes[0] {
		return e.Definitions, nil
	}
	return nil, &NotLoadedError{edge: "definitions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Synset) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case synset.FieldID, synset.FieldPartOfSpeech:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Synset fields.
func (s *Synset) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case synset.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case synset.FieldPartOfSpeech:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field partOfSpeech", values[i])
			} else if value.Valid {
				s.PartOfSpeech = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Synset.
// This includes values selected through modifiers, order, etc.
func (s *Synset) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryDefinitions queries the "definitions" edge of the Synset entity.
func (s *Synset) QueryDefinitions() *DefinitionQuery {
	return NewSynsetClient(s.config).QueryDefinitions(s)
}

// Update returns a builder for updating this Synset.
// Note that you need to call Synset.Unwrap() before calling this method if this Synset
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Synset) Update() *SynsetUpdateOne {
	return NewSynsetClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Synset entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Synset) Unwrap() *Synset {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Synset is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Synset) String() string {
	var builder strings.Builder
	builder.WriteString("Synset(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("partOfSpeech=")
	builder.WriteString(s.PartOfSpeech)
	builder.WriteByte(')')
	return builder.String()
}

// Synsets is a parsable slice of Synset.
type Synsets []*Synset
