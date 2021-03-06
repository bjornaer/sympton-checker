// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bjornaer/sympton-checker/ent/ailment"
	"github.com/bjornaer/sympton-checker/ent/predicate"
	"github.com/bjornaer/sympton-checker/ent/schema"
)

// AilmentUpdate is the builder for updating Ailment entities.
type AilmentUpdate struct {
	config
	hooks    []Hook
	mutation *AilmentMutation
}

// Where appends a list predicates to the AilmentUpdate builder.
func (au *AilmentUpdate) Where(ps ...predicate.Ailment) *AilmentUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AilmentUpdate) SetName(s string) *AilmentUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AilmentUpdate) SetNillableName(s *string) *AilmentUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetSymptoms sets the "symptoms" field.
func (au *AilmentUpdate) SetSymptoms(md map[string]schema.SymptomDetails) *AilmentUpdate {
	au.mutation.SetSymptoms(md)
	return au
}

// SetHpos sets the "hpos" field.
func (au *AilmentUpdate) SetHpos(s []string) *AilmentUpdate {
	au.mutation.SetHpos(s)
	return au
}

// SetExpert sets the "expert" field.
func (au *AilmentUpdate) SetExpert(s string) *AilmentUpdate {
	au.mutation.SetExpert(s)
	return au
}

// SetNillableExpert sets the "expert" field if the given value is not nil.
func (au *AilmentUpdate) SetNillableExpert(s *string) *AilmentUpdate {
	if s != nil {
		au.SetExpert(*s)
	}
	return au
}

// Mutation returns the AilmentMutation object of the builder.
func (au *AilmentUpdate) Mutation() *AilmentMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AilmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AilmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AilmentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AilmentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AilmentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AilmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ailment.Table,
			Columns: ailment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ailment.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ailment.FieldName,
		})
	}
	if value, ok := au.mutation.Symptoms(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: ailment.FieldSymptoms,
		})
	}
	if value, ok := au.mutation.Hpos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: ailment.FieldHpos,
		})
	}
	if value, ok := au.mutation.Expert(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ailment.FieldExpert,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ailment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AilmentUpdateOne is the builder for updating a single Ailment entity.
type AilmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AilmentMutation
}

// SetName sets the "name" field.
func (auo *AilmentUpdateOne) SetName(s string) *AilmentUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AilmentUpdateOne) SetNillableName(s *string) *AilmentUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetSymptoms sets the "symptoms" field.
func (auo *AilmentUpdateOne) SetSymptoms(md map[string]schema.SymptomDetails) *AilmentUpdateOne {
	auo.mutation.SetSymptoms(md)
	return auo
}

// SetHpos sets the "hpos" field.
func (auo *AilmentUpdateOne) SetHpos(s []string) *AilmentUpdateOne {
	auo.mutation.SetHpos(s)
	return auo
}

// SetExpert sets the "expert" field.
func (auo *AilmentUpdateOne) SetExpert(s string) *AilmentUpdateOne {
	auo.mutation.SetExpert(s)
	return auo
}

// SetNillableExpert sets the "expert" field if the given value is not nil.
func (auo *AilmentUpdateOne) SetNillableExpert(s *string) *AilmentUpdateOne {
	if s != nil {
		auo.SetExpert(*s)
	}
	return auo
}

// Mutation returns the AilmentMutation object of the builder.
func (auo *AilmentUpdateOne) Mutation() *AilmentMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AilmentUpdateOne) Select(field string, fields ...string) *AilmentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Ailment entity.
func (auo *AilmentUpdateOne) Save(ctx context.Context) (*Ailment, error) {
	var (
		err  error
		node *Ailment
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AilmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AilmentUpdateOne) SaveX(ctx context.Context) *Ailment {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AilmentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AilmentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AilmentUpdateOne) sqlSave(ctx context.Context) (_node *Ailment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ailment.Table,
			Columns: ailment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ailment.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ailment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ailment.FieldID)
		for _, f := range fields {
			if !ailment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ailment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ailment.FieldName,
		})
	}
	if value, ok := auo.mutation.Symptoms(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: ailment.FieldSymptoms,
		})
	}
	if value, ok := auo.mutation.Hpos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: ailment.FieldHpos,
		})
	}
	if value, ok := auo.mutation.Expert(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ailment.FieldExpert,
		})
	}
	_node = &Ailment{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ailment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
