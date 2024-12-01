// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/allof-dev/dictionary/ent/definition"
	"github.com/allof-dev/dictionary/ent/synset"
)

// DefinitionCreate is the builder for creating a Definition entity.
type DefinitionCreate struct {
	config
	mutation *DefinitionMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (dc *DefinitionCreate) SetText(s string) *DefinitionCreate {
	dc.mutation.SetText(s)
	return dc
}

// SetSynsetID sets the "synset" edge to the Synset entity by ID.
func (dc *DefinitionCreate) SetSynsetID(id string) *DefinitionCreate {
	dc.mutation.SetSynsetID(id)
	return dc
}

// SetNillableSynsetID sets the "synset" edge to the Synset entity by ID if the given value is not nil.
func (dc *DefinitionCreate) SetNillableSynsetID(id *string) *DefinitionCreate {
	if id != nil {
		dc = dc.SetSynsetID(*id)
	}
	return dc
}

// SetSynset sets the "synset" edge to the Synset entity.
func (dc *DefinitionCreate) SetSynset(s *Synset) *DefinitionCreate {
	return dc.SetSynsetID(s.ID)
}

// Mutation returns the DefinitionMutation object of the builder.
func (dc *DefinitionCreate) Mutation() *DefinitionMutation {
	return dc.mutation
}

// Save creates the Definition in the database.
func (dc *DefinitionCreate) Save(ctx context.Context) (*Definition, error) {
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DefinitionCreate) SaveX(ctx context.Context) *Definition {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DefinitionCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DefinitionCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DefinitionCreate) check() error {
	if _, ok := dc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Definition.text"`)}
	}
	if v, ok := dc.mutation.Text(); ok {
		if err := definition.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Definition.text": %w`, err)}
		}
	}
	return nil
}

func (dc *DefinitionCreate) sqlSave(ctx context.Context) (*Definition, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DefinitionCreate) createSpec() (*Definition, *sqlgraph.CreateSpec) {
	var (
		_node = &Definition{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(definition.Table, sqlgraph.NewFieldSpec(definition.FieldID, field.TypeInt))
	)
	if value, ok := dc.mutation.Text(); ok {
		_spec.SetField(definition.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if nodes := dc.mutation.SynsetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   definition.SynsetTable,
			Columns: []string{definition.SynsetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(synset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.synset_definitions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DefinitionCreateBulk is the builder for creating many Definition entities in bulk.
type DefinitionCreateBulk struct {
	config
	err      error
	builders []*DefinitionCreate
}

// Save creates the Definition entities in the database.
func (dcb *DefinitionCreateBulk) Save(ctx context.Context) ([]*Definition, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Definition, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DefinitionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DefinitionCreateBulk) SaveX(ctx context.Context) []*Definition {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DefinitionCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DefinitionCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
