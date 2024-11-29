package schema

import "entgo.io/ent"

// Lexeme holds the schema definition for the Lexeme entity.
type Lexeme struct {
	ent.Schema
}

// Fields of the Lexeme.
func (Lexeme) Fields() []ent.Field {
	return nil
}

// Edges of the Lexeme.
func (Lexeme) Edges() []ent.Edge {
	return nil
}
