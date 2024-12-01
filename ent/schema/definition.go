package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Definition holds the schema definition for the Definition entity.
type Definition struct {
	ent.Schema
}

func (Definition) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("synset"),
	}
}

// Fields of the Definition.
func (Definition) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").NotEmpty(),
	}
}

// Edges of the Definition.
func (Definition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("synset", Synset.Type).
			Ref("definitions").
			Unique(),
	}
}
