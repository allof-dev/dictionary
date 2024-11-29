package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Definition holds the schema definition for the Definition entity.
type Definition struct {
	ent.Schema
}

// Fields of the Definition.
func (Definition) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").NotEmpty(),
	}
}

// Edges of the Definition.
func (Definition) Edges() []ent.Edge {
	return nil
}
