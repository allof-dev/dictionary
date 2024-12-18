package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Lemma holds the schema definition for the Lemma entity.
type Lemma struct {
	ent.Schema
}

func (Lemma) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("writtenForm"),
	}
}

// Fields of the Lemma.
func (Lemma) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty(),
		field.String("writtenForm").NotEmpty(),
		field.String("partOfSpeech").NotEmpty(),
	}
}

// Edges of the Lemma.
func (Lemma) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("senses", Sense.Type).
			Ref("lemma"),
	}
}
