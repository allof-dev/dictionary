package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Sense holds the schema definition for the Sense entity.
type Sense struct {
	ent.Schema
}

// Fields of the Sense.
func (Sense) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty(),
	}
}

// Edges of the Sense.
func (Sense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("synset", Synset.Type).
			Unique(),
		edge.To("lemma", Lemma.Type).
			Unique(),

		edge.From("relFrom", SenseRelation.Type).
			Ref("to"),
		edge.From("relTo", SenseRelation.Type).
			Ref("from"),
	}
}
