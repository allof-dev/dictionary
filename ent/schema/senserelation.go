package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SenseRelation holds the schema definition for the SenseRelation entity.
type SenseRelation struct {
	ent.Schema
}

// Fields of the SenseRelation.
func (SenseRelation) Fields() []ent.Field {
	return []ent.Field{
		field.String("relType").NotEmpty(),
	}
}

// Edges of the SenseRelation.
func (SenseRelation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("from", Sense.Type).Required().Unique(),
		edge.To("to", Sense.Type).Required().Unique(),
	}
}
