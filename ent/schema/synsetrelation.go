package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SynsetRelation holds the schema definition for the SynsetRelation entity.
type SynsetRelation struct {
	ent.Schema
}

// Fields of the SynsetRelation.
func (SynsetRelation) Fields() []ent.Field {
	return []ent.Field{
		field.String("relType").NotEmpty(),
	}
}

// Edges of the SynsetRelation.
func (SynsetRelation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("from", Synset.Type).Required().Unique(),
		edge.To("to", Synset.Type).Required().Unique(),
	}
}
