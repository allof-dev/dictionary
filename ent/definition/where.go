// Code generated by ent, DO NOT EDIT.

package definition

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/allof-dev/dictionary/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Definition {
	return predicate.Definition(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Definition {
	return predicate.Definition(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Definition {
	return predicate.Definition(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Definition {
	return predicate.Definition(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Definition {
	return predicate.Definition(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Definition {
	return predicate.Definition(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Definition {
	return predicate.Definition(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Definition {
	return predicate.Definition(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Definition {
	return predicate.Definition(sql.FieldLTE(FieldID, id))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Definition {
	return predicate.Definition(sql.FieldEQ(FieldText, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Definition {
	return predicate.Definition(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Definition {
	return predicate.Definition(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Definition {
	return predicate.Definition(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Definition {
	return predicate.Definition(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Definition {
	return predicate.Definition(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Definition {
	return predicate.Definition(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Definition {
	return predicate.Definition(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Definition {
	return predicate.Definition(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Definition {
	return predicate.Definition(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Definition {
	return predicate.Definition(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Definition {
	return predicate.Definition(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Definition {
	return predicate.Definition(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Definition {
	return predicate.Definition(sql.FieldContainsFold(FieldText, v))
}

// HasSynset applies the HasEdge predicate on the "synset" edge.
func HasSynset() predicate.Definition {
	return predicate.Definition(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SynsetTable, SynsetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSynsetWith applies the HasEdge predicate on the "synset" edge with a given conditions (other predicates).
func HasSynsetWith(preds ...predicate.Synset) predicate.Definition {
	return predicate.Definition(func(s *sql.Selector) {
		step := newSynsetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Definition) predicate.Definition {
	return predicate.Definition(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Definition) predicate.Definition {
	return predicate.Definition(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Definition) predicate.Definition {
	return predicate.Definition(sql.NotPredicates(p))
}
