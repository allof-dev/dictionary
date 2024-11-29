// Code generated by ent, DO NOT EDIT.

package synsetrelation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/allof-dev/dictionary/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldLTE(FieldID, id))
}

// RelType applies equality check predicate on the "relType" field. It's identical to RelTypeEQ.
func RelType(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldEQ(FieldRelType, v))
}

// RelTypeEQ applies the EQ predicate on the "relType" field.
func RelTypeEQ(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldEQ(FieldRelType, v))
}

// RelTypeNEQ applies the NEQ predicate on the "relType" field.
func RelTypeNEQ(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldNEQ(FieldRelType, v))
}

// RelTypeIn applies the In predicate on the "relType" field.
func RelTypeIn(vs ...string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldIn(FieldRelType, vs...))
}

// RelTypeNotIn applies the NotIn predicate on the "relType" field.
func RelTypeNotIn(vs ...string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldNotIn(FieldRelType, vs...))
}

// RelTypeGT applies the GT predicate on the "relType" field.
func RelTypeGT(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldGT(FieldRelType, v))
}

// RelTypeGTE applies the GTE predicate on the "relType" field.
func RelTypeGTE(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldGTE(FieldRelType, v))
}

// RelTypeLT applies the LT predicate on the "relType" field.
func RelTypeLT(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldLT(FieldRelType, v))
}

// RelTypeLTE applies the LTE predicate on the "relType" field.
func RelTypeLTE(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldLTE(FieldRelType, v))
}

// RelTypeContains applies the Contains predicate on the "relType" field.
func RelTypeContains(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldContains(FieldRelType, v))
}

// RelTypeHasPrefix applies the HasPrefix predicate on the "relType" field.
func RelTypeHasPrefix(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldHasPrefix(FieldRelType, v))
}

// RelTypeHasSuffix applies the HasSuffix predicate on the "relType" field.
func RelTypeHasSuffix(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldHasSuffix(FieldRelType, v))
}

// RelTypeEqualFold applies the EqualFold predicate on the "relType" field.
func RelTypeEqualFold(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldEqualFold(FieldRelType, v))
}

// RelTypeContainsFold applies the ContainsFold predicate on the "relType" field.
func RelTypeContainsFold(v string) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.FieldContainsFold(FieldRelType, v))
}

// HasFrom applies the HasEdge predicate on the "from" edge.
func HasFrom() predicate.SynsetRelation {
	return predicate.SynsetRelation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FromTable, FromColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFromWith applies the HasEdge predicate on the "from" edge with a given conditions (other predicates).
func HasFromWith(preds ...predicate.Synset) predicate.SynsetRelation {
	return predicate.SynsetRelation(func(s *sql.Selector) {
		step := newFromStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTo applies the HasEdge predicate on the "to" edge.
func HasTo() predicate.SynsetRelation {
	return predicate.SynsetRelation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ToTable, ToColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasToWith applies the HasEdge predicate on the "to" edge with a given conditions (other predicates).
func HasToWith(preds ...predicate.Synset) predicate.SynsetRelation {
	return predicate.SynsetRelation(func(s *sql.Selector) {
		step := newToStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SynsetRelation) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SynsetRelation) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SynsetRelation) predicate.SynsetRelation {
	return predicate.SynsetRelation(sql.NotPredicates(p))
}