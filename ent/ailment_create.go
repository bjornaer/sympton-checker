// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bjornaer/sympton-checker/ent/ailment"
	"github.com/bjornaer/sympton-checker/ent/schema"
)

// AilmentCreate is the builder for creating a Ailment entity.
type AilmentCreate struct {
	config
	mutation *AilmentMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AilmentCreate) SetName(s string) *AilmentCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ac *AilmentCreate) SetNillableName(s *string) *AilmentCreate {
	if s != nil {
		ac.SetName(*s)
	}
	return ac
}

// SetSymptoms sets the "symptoms" field.
func (ac *AilmentCreate) SetSymptoms(md map[string]schema.SymptomDetails) *AilmentCreate {
	ac.mutation.SetSymptoms(md)
	return ac
}

// SetHpos sets the "hpos" field.
func (ac *AilmentCreate) SetHpos(s []string) *AilmentCreate {
	ac.mutation.SetHpos(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AilmentCreate) SetID(i int) *AilmentCreate {
	ac.mutation.SetID(i)
	return ac
}

// Mutation returns the AilmentMutation object of the builder.
func (ac *AilmentCreate) Mutation() *AilmentMutation {
	return ac.mutation
}

// Save creates the Ailment in the database.
func (ac *AilmentCreate) Save(ctx context.Context) (*Ailment, error) {
	var (
		err  error
		node *Ailment
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AilmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AilmentCreate) SaveX(ctx context.Context) *Ailment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AilmentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AilmentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AilmentCreate) defaults() {
	if _, ok := ac.mutation.Name(); !ok {
		v := ailment.DefaultName
		ac.mutation.SetName(v)
	}
	if _, ok := ac.mutation.Symptoms(); !ok {
		v := ailment.DefaultSymptoms
		ac.mutation.SetSymptoms(v)
	}
	if _, ok := ac.mutation.Hpos(); !ok {
		v := ailment.DefaultHpos
		ac.mutation.SetHpos(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AilmentCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Ailment.name"`)}
	}
	if _, ok := ac.mutation.Symptoms(); !ok {
		return &ValidationError{Name: "symptoms", err: errors.New(`ent: missing required field "Ailment.symptoms"`)}
	}
	if _, ok := ac.mutation.Hpos(); !ok {
		return &ValidationError{Name: "hpos", err: errors.New(`ent: missing required field "Ailment.hpos"`)}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := ailment.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Ailment.id": %w`, err)}
		}
	}
	return nil
}

func (ac *AilmentCreate) sqlSave(ctx context.Context) (*Ailment, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (ac *AilmentCreate) createSpec() (*Ailment, *sqlgraph.CreateSpec) {
	var (
		_node = &Ailment{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ailment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ailment.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ailment.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Symptoms(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: ailment.FieldSymptoms,
		})
		_node.Symptoms = value
	}
	if value, ok := ac.mutation.Hpos(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: ailment.FieldHpos,
		})
		_node.Hpos = value
	}
	return _node, _spec
}

// AilmentCreateBulk is the builder for creating many Ailment entities in bulk.
type AilmentCreateBulk struct {
	config
	builders []*AilmentCreate
}

// Save creates the Ailment entities in the database.
func (acb *AilmentCreateBulk) Save(ctx context.Context) ([]*Ailment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Ailment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AilmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AilmentCreateBulk) SaveX(ctx context.Context) []*Ailment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AilmentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AilmentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
