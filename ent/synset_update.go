// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/allof-dev/dictionary/ent/definition"
	"github.com/allof-dev/dictionary/ent/predicate"
	"github.com/allof-dev/dictionary/ent/synset"
	"github.com/allof-dev/dictionary/ent/synsetrelation"
)

// SynsetUpdate is the builder for updating Synset entities.
type SynsetUpdate struct {
	config
	hooks    []Hook
	mutation *SynsetMutation
}

// Where appends a list predicates to the SynsetUpdate builder.
func (su *SynsetUpdate) Where(ps ...predicate.Synset) *SynsetUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetPartOfSpeech sets the "partOfSpeech" field.
func (su *SynsetUpdate) SetPartOfSpeech(s string) *SynsetUpdate {
	su.mutation.SetPartOfSpeech(s)
	return su
}

// SetNillablePartOfSpeech sets the "partOfSpeech" field if the given value is not nil.
func (su *SynsetUpdate) SetNillablePartOfSpeech(s *string) *SynsetUpdate {
	if s != nil {
		su.SetPartOfSpeech(*s)
	}
	return su
}

// AddDefinitionIDs adds the "definitions" edge to the Definition entity by IDs.
func (su *SynsetUpdate) AddDefinitionIDs(ids ...int) *SynsetUpdate {
	su.mutation.AddDefinitionIDs(ids...)
	return su
}

// AddDefinitions adds the "definitions" edges to the Definition entity.
func (su *SynsetUpdate) AddDefinitions(d ...*Definition) *SynsetUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.AddDefinitionIDs(ids...)
}

// AddRelFromIDs adds the "relFrom" edge to the SynsetRelation entity by IDs.
func (su *SynsetUpdate) AddRelFromIDs(ids ...int) *SynsetUpdate {
	su.mutation.AddRelFromIDs(ids...)
	return su
}

// AddRelFrom adds the "relFrom" edges to the SynsetRelation entity.
func (su *SynsetUpdate) AddRelFrom(s ...*SynsetRelation) *SynsetUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddRelFromIDs(ids...)
}

// AddRelToIDs adds the "relTo" edge to the SynsetRelation entity by IDs.
func (su *SynsetUpdate) AddRelToIDs(ids ...int) *SynsetUpdate {
	su.mutation.AddRelToIDs(ids...)
	return su
}

// AddRelTo adds the "relTo" edges to the SynsetRelation entity.
func (su *SynsetUpdate) AddRelTo(s ...*SynsetRelation) *SynsetUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddRelToIDs(ids...)
}

// Mutation returns the SynsetMutation object of the builder.
func (su *SynsetUpdate) Mutation() *SynsetMutation {
	return su.mutation
}

// ClearDefinitions clears all "definitions" edges to the Definition entity.
func (su *SynsetUpdate) ClearDefinitions() *SynsetUpdate {
	su.mutation.ClearDefinitions()
	return su
}

// RemoveDefinitionIDs removes the "definitions" edge to Definition entities by IDs.
func (su *SynsetUpdate) RemoveDefinitionIDs(ids ...int) *SynsetUpdate {
	su.mutation.RemoveDefinitionIDs(ids...)
	return su
}

// RemoveDefinitions removes "definitions" edges to Definition entities.
func (su *SynsetUpdate) RemoveDefinitions(d ...*Definition) *SynsetUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return su.RemoveDefinitionIDs(ids...)
}

// ClearRelFrom clears all "relFrom" edges to the SynsetRelation entity.
func (su *SynsetUpdate) ClearRelFrom() *SynsetUpdate {
	su.mutation.ClearRelFrom()
	return su
}

// RemoveRelFromIDs removes the "relFrom" edge to SynsetRelation entities by IDs.
func (su *SynsetUpdate) RemoveRelFromIDs(ids ...int) *SynsetUpdate {
	su.mutation.RemoveRelFromIDs(ids...)
	return su
}

// RemoveRelFrom removes "relFrom" edges to SynsetRelation entities.
func (su *SynsetUpdate) RemoveRelFrom(s ...*SynsetRelation) *SynsetUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveRelFromIDs(ids...)
}

// ClearRelTo clears all "relTo" edges to the SynsetRelation entity.
func (su *SynsetUpdate) ClearRelTo() *SynsetUpdate {
	su.mutation.ClearRelTo()
	return su
}

// RemoveRelToIDs removes the "relTo" edge to SynsetRelation entities by IDs.
func (su *SynsetUpdate) RemoveRelToIDs(ids ...int) *SynsetUpdate {
	su.mutation.RemoveRelToIDs(ids...)
	return su
}

// RemoveRelTo removes "relTo" edges to SynsetRelation entities.
func (su *SynsetUpdate) RemoveRelTo(s ...*SynsetRelation) *SynsetUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveRelToIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SynsetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SynsetUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SynsetUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SynsetUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SynsetUpdate) check() error {
	if v, ok := su.mutation.PartOfSpeech(); ok {
		if err := synset.PartOfSpeechValidator(v); err != nil {
			return &ValidationError{Name: "partOfSpeech", err: fmt.Errorf(`ent: validator failed for field "Synset.partOfSpeech": %w`, err)}
		}
	}
	return nil
}

func (su *SynsetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(synset.Table, synset.Columns, sqlgraph.NewFieldSpec(synset.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.PartOfSpeech(); ok {
		_spec.SetField(synset.FieldPartOfSpeech, field.TypeString, value)
	}
	if su.mutation.DefinitionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   synset.DefinitionsTable,
			Columns: []string{synset.DefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(definition.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedDefinitionsIDs(); len(nodes) > 0 && !su.mutation.DefinitionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   synset.DefinitionsTable,
			Columns: []string{synset.DefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(definition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.DefinitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   synset.DefinitionsTable,
			Columns: []string{synset.DefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(definition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.RelFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelFromTable,
			Columns: []string{synset.RelFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedRelFromIDs(); len(nodes) > 0 && !su.mutation.RelFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelFromTable,
			Columns: []string{synset.RelFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RelFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelFromTable,
			Columns: []string{synset.RelFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.RelToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelToTable,
			Columns: []string{synset.RelToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedRelToIDs(); len(nodes) > 0 && !su.mutation.RelToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelToTable,
			Columns: []string{synset.RelToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RelToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelToTable,
			Columns: []string{synset.RelToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{synset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SynsetUpdateOne is the builder for updating a single Synset entity.
type SynsetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SynsetMutation
}

// SetPartOfSpeech sets the "partOfSpeech" field.
func (suo *SynsetUpdateOne) SetPartOfSpeech(s string) *SynsetUpdateOne {
	suo.mutation.SetPartOfSpeech(s)
	return suo
}

// SetNillablePartOfSpeech sets the "partOfSpeech" field if the given value is not nil.
func (suo *SynsetUpdateOne) SetNillablePartOfSpeech(s *string) *SynsetUpdateOne {
	if s != nil {
		suo.SetPartOfSpeech(*s)
	}
	return suo
}

// AddDefinitionIDs adds the "definitions" edge to the Definition entity by IDs.
func (suo *SynsetUpdateOne) AddDefinitionIDs(ids ...int) *SynsetUpdateOne {
	suo.mutation.AddDefinitionIDs(ids...)
	return suo
}

// AddDefinitions adds the "definitions" edges to the Definition entity.
func (suo *SynsetUpdateOne) AddDefinitions(d ...*Definition) *SynsetUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.AddDefinitionIDs(ids...)
}

// AddRelFromIDs adds the "relFrom" edge to the SynsetRelation entity by IDs.
func (suo *SynsetUpdateOne) AddRelFromIDs(ids ...int) *SynsetUpdateOne {
	suo.mutation.AddRelFromIDs(ids...)
	return suo
}

// AddRelFrom adds the "relFrom" edges to the SynsetRelation entity.
func (suo *SynsetUpdateOne) AddRelFrom(s ...*SynsetRelation) *SynsetUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddRelFromIDs(ids...)
}

// AddRelToIDs adds the "relTo" edge to the SynsetRelation entity by IDs.
func (suo *SynsetUpdateOne) AddRelToIDs(ids ...int) *SynsetUpdateOne {
	suo.mutation.AddRelToIDs(ids...)
	return suo
}

// AddRelTo adds the "relTo" edges to the SynsetRelation entity.
func (suo *SynsetUpdateOne) AddRelTo(s ...*SynsetRelation) *SynsetUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddRelToIDs(ids...)
}

// Mutation returns the SynsetMutation object of the builder.
func (suo *SynsetUpdateOne) Mutation() *SynsetMutation {
	return suo.mutation
}

// ClearDefinitions clears all "definitions" edges to the Definition entity.
func (suo *SynsetUpdateOne) ClearDefinitions() *SynsetUpdateOne {
	suo.mutation.ClearDefinitions()
	return suo
}

// RemoveDefinitionIDs removes the "definitions" edge to Definition entities by IDs.
func (suo *SynsetUpdateOne) RemoveDefinitionIDs(ids ...int) *SynsetUpdateOne {
	suo.mutation.RemoveDefinitionIDs(ids...)
	return suo
}

// RemoveDefinitions removes "definitions" edges to Definition entities.
func (suo *SynsetUpdateOne) RemoveDefinitions(d ...*Definition) *SynsetUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return suo.RemoveDefinitionIDs(ids...)
}

// ClearRelFrom clears all "relFrom" edges to the SynsetRelation entity.
func (suo *SynsetUpdateOne) ClearRelFrom() *SynsetUpdateOne {
	suo.mutation.ClearRelFrom()
	return suo
}

// RemoveRelFromIDs removes the "relFrom" edge to SynsetRelation entities by IDs.
func (suo *SynsetUpdateOne) RemoveRelFromIDs(ids ...int) *SynsetUpdateOne {
	suo.mutation.RemoveRelFromIDs(ids...)
	return suo
}

// RemoveRelFrom removes "relFrom" edges to SynsetRelation entities.
func (suo *SynsetUpdateOne) RemoveRelFrom(s ...*SynsetRelation) *SynsetUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveRelFromIDs(ids...)
}

// ClearRelTo clears all "relTo" edges to the SynsetRelation entity.
func (suo *SynsetUpdateOne) ClearRelTo() *SynsetUpdateOne {
	suo.mutation.ClearRelTo()
	return suo
}

// RemoveRelToIDs removes the "relTo" edge to SynsetRelation entities by IDs.
func (suo *SynsetUpdateOne) RemoveRelToIDs(ids ...int) *SynsetUpdateOne {
	suo.mutation.RemoveRelToIDs(ids...)
	return suo
}

// RemoveRelTo removes "relTo" edges to SynsetRelation entities.
func (suo *SynsetUpdateOne) RemoveRelTo(s ...*SynsetRelation) *SynsetUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveRelToIDs(ids...)
}

// Where appends a list predicates to the SynsetUpdate builder.
func (suo *SynsetUpdateOne) Where(ps ...predicate.Synset) *SynsetUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SynsetUpdateOne) Select(field string, fields ...string) *SynsetUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Synset entity.
func (suo *SynsetUpdateOne) Save(ctx context.Context) (*Synset, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SynsetUpdateOne) SaveX(ctx context.Context) *Synset {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SynsetUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SynsetUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SynsetUpdateOne) check() error {
	if v, ok := suo.mutation.PartOfSpeech(); ok {
		if err := synset.PartOfSpeechValidator(v); err != nil {
			return &ValidationError{Name: "partOfSpeech", err: fmt.Errorf(`ent: validator failed for field "Synset.partOfSpeech": %w`, err)}
		}
	}
	return nil
}

func (suo *SynsetUpdateOne) sqlSave(ctx context.Context) (_node *Synset, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(synset.Table, synset.Columns, sqlgraph.NewFieldSpec(synset.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Synset.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, synset.FieldID)
		for _, f := range fields {
			if !synset.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != synset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.PartOfSpeech(); ok {
		_spec.SetField(synset.FieldPartOfSpeech, field.TypeString, value)
	}
	if suo.mutation.DefinitionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   synset.DefinitionsTable,
			Columns: []string{synset.DefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(definition.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedDefinitionsIDs(); len(nodes) > 0 && !suo.mutation.DefinitionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   synset.DefinitionsTable,
			Columns: []string{synset.DefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(definition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.DefinitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   synset.DefinitionsTable,
			Columns: []string{synset.DefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(definition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.RelFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelFromTable,
			Columns: []string{synset.RelFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedRelFromIDs(); len(nodes) > 0 && !suo.mutation.RelFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelFromTable,
			Columns: []string{synset.RelFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RelFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelFromTable,
			Columns: []string{synset.RelFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.RelToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelToTable,
			Columns: []string{synset.RelToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedRelToIDs(); len(nodes) > 0 && !suo.mutation.RelToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelToTable,
			Columns: []string{synset.RelToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RelToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   synset.RelToTable,
			Columns: []string{synset.RelToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synsetrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Synset{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{synset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
