// Code generated by ent, DO NOT EDIT.

package lemma

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/allof-dev/dictionary/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Lemma {
	return predicate.Lemma(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Lemma {
	return predicate.Lemma(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Lemma {
	return predicate.Lemma(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Lemma {
	return predicate.Lemma(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Lemma {
	return predicate.Lemma(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Lemma {
	return predicate.Lemma(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Lemma {
	return predicate.Lemma(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Lemma {
	return predicate.Lemma(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Lemma {
	return predicate.Lemma(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Lemma {
	return predicate.Lemma(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Lemma {
	return predicate.Lemma(sql.FieldContainsFold(FieldID, id))
}

// WrittenForm applies equality check predicate on the "writtenForm" field. It's identical to WrittenFormEQ.
func WrittenForm(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldEQ(FieldWrittenForm, v))
}

// PartOfSpeech applies equality check predicate on the "partOfSpeech" field. It's identical to PartOfSpeechEQ.
func PartOfSpeech(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldEQ(FieldPartOfSpeech, v))
}

// WrittenFormEQ applies the EQ predicate on the "writtenForm" field.
func WrittenFormEQ(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldEQ(FieldWrittenForm, v))
}

// WrittenFormNEQ applies the NEQ predicate on the "writtenForm" field.
func WrittenFormNEQ(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldNEQ(FieldWrittenForm, v))
}

// WrittenFormIn applies the In predicate on the "writtenForm" field.
func WrittenFormIn(vs ...string) predicate.Lemma {
	return predicate.Lemma(sql.FieldIn(FieldWrittenForm, vs...))
}

// WrittenFormNotIn applies the NotIn predicate on the "writtenForm" field.
func WrittenFormNotIn(vs ...string) predicate.Lemma {
	return predicate.Lemma(sql.FieldNotIn(FieldWrittenForm, vs...))
}

// WrittenFormGT applies the GT predicate on the "writtenForm" field.
func WrittenFormGT(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldGT(FieldWrittenForm, v))
}

// WrittenFormGTE applies the GTE predicate on the "writtenForm" field.
func WrittenFormGTE(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldGTE(FieldWrittenForm, v))
}

// WrittenFormLT applies the LT predicate on the "writtenForm" field.
func WrittenFormLT(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldLT(FieldWrittenForm, v))
}

// WrittenFormLTE applies the LTE predicate on the "writtenForm" field.
func WrittenFormLTE(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldLTE(FieldWrittenForm, v))
}

// WrittenFormContains applies the Contains predicate on the "writtenForm" field.
func WrittenFormContains(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldContains(FieldWrittenForm, v))
}

// WrittenFormHasPrefix applies the HasPrefix predicate on the "writtenForm" field.
func WrittenFormHasPrefix(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldHasPrefix(FieldWrittenForm, v))
}

// WrittenFormHasSuffix applies the HasSuffix predicate on the "writtenForm" field.
func WrittenFormHasSuffix(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldHasSuffix(FieldWrittenForm, v))
}

// WrittenFormEqualFold applies the EqualFold predicate on the "writtenForm" field.
func WrittenFormEqualFold(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldEqualFold(FieldWrittenForm, v))
}

// WrittenFormContainsFold applies the ContainsFold predicate on the "writtenForm" field.
func WrittenFormContainsFold(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldContainsFold(FieldWrittenForm, v))
}

// PartOfSpeechEQ applies the EQ predicate on the "partOfSpeech" field.
func PartOfSpeechEQ(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldEQ(FieldPartOfSpeech, v))
}

// PartOfSpeechNEQ applies the NEQ predicate on the "partOfSpeech" field.
func PartOfSpeechNEQ(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldNEQ(FieldPartOfSpeech, v))
}

// PartOfSpeechIn applies the In predicate on the "partOfSpeech" field.
func PartOfSpeechIn(vs ...string) predicate.Lemma {
	return predicate.Lemma(sql.FieldIn(FieldPartOfSpeech, vs...))
}

// PartOfSpeechNotIn applies the NotIn predicate on the "partOfSpeech" field.
func PartOfSpeechNotIn(vs ...string) predicate.Lemma {
	return predicate.Lemma(sql.FieldNotIn(FieldPartOfSpeech, vs...))
}

// PartOfSpeechGT applies the GT predicate on the "partOfSpeech" field.
func PartOfSpeechGT(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldGT(FieldPartOfSpeech, v))
}

// PartOfSpeechGTE applies the GTE predicate on the "partOfSpeech" field.
func PartOfSpeechGTE(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldGTE(FieldPartOfSpeech, v))
}

// PartOfSpeechLT applies the LT predicate on the "partOfSpeech" field.
func PartOfSpeechLT(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldLT(FieldPartOfSpeech, v))
}

// PartOfSpeechLTE applies the LTE predicate on the "partOfSpeech" field.
func PartOfSpeechLTE(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldLTE(FieldPartOfSpeech, v))
}

// PartOfSpeechContains applies the Contains predicate on the "partOfSpeech" field.
func PartOfSpeechContains(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldContains(FieldPartOfSpeech, v))
}

// PartOfSpeechHasPrefix applies the HasPrefix predicate on the "partOfSpeech" field.
func PartOfSpeechHasPrefix(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldHasPrefix(FieldPartOfSpeech, v))
}

// PartOfSpeechHasSuffix applies the HasSuffix predicate on the "partOfSpeech" field.
func PartOfSpeechHasSuffix(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldHasSuffix(FieldPartOfSpeech, v))
}

// PartOfSpeechEqualFold applies the EqualFold predicate on the "partOfSpeech" field.
func PartOfSpeechEqualFold(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldEqualFold(FieldPartOfSpeech, v))
}

// PartOfSpeechContainsFold applies the ContainsFold predicate on the "partOfSpeech" field.
func PartOfSpeechContainsFold(v string) predicate.Lemma {
	return predicate.Lemma(sql.FieldContainsFold(FieldPartOfSpeech, v))
}

// HasSenses applies the HasEdge predicate on the "senses" edge.
func HasSenses() predicate.Lemma {
	return predicate.Lemma(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, SensesTable, SensesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSensesWith applies the HasEdge predicate on the "senses" edge with a given conditions (other predicates).
func HasSensesWith(preds ...predicate.Sense) predicate.Lemma {
	return predicate.Lemma(func(s *sql.Selector) {
		step := newSensesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Lemma) predicate.Lemma {
	return predicate.Lemma(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Lemma) predicate.Lemma {
	return predicate.Lemma(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Lemma) predicate.Lemma {
	return predicate.Lemma(sql.NotPredicates(p))
}
