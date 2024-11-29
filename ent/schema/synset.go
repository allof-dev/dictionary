package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Synset holds the schema definition for the Synset entity.
type Synset struct {
	ent.Schema
}

// Fields of the Synset.
func (Synset) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty(),
		field.String("partOfSpeech").NotEmpty(),
	}
}

// Edges of the Synset.
func (Synset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("definitions", Definition.Type),
	}
}
