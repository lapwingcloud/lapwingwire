// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lapwingcloud/lapwingwire/controller/ent/agent"
	"github.com/lapwingcloud/lapwingwire/controller/ent/predicate"
	"github.com/lapwingcloud/lapwingwire/controller/ent/tag"
)

// AgentUpdate is the builder for updating Agent entities.
type AgentUpdate struct {
	config
	hooks    []Hook
	mutation *AgentMutation
}

// Where appends a list predicates to the AgentUpdate builder.
func (au *AgentUpdate) Where(ps ...predicate.Agent) *AgentUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetHostname sets the "hostname" field.
func (au *AgentUpdate) SetHostname(s string) *AgentUpdate {
	au.mutation.SetHostname(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AgentUpdate) SetCreatedAt(t time.Time) *AgentUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AgentUpdate) SetNillableCreatedAt(t *time.Time) *AgentUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (au *AgentUpdate) AddTagIDs(ids ...int) *AgentUpdate {
	au.mutation.AddTagIDs(ids...)
	return au
}

// AddTags adds the "tags" edges to the Tag entity.
func (au *AgentUpdate) AddTags(t ...*Tag) *AgentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTagIDs(ids...)
}

// Mutation returns the AgentMutation object of the builder.
func (au *AgentUpdate) Mutation() *AgentMutation {
	return au.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (au *AgentUpdate) ClearTags() *AgentUpdate {
	au.mutation.ClearTags()
	return au
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (au *AgentUpdate) RemoveTagIDs(ids ...int) *AgentUpdate {
	au.mutation.RemoveTagIDs(ids...)
	return au
}

// RemoveTags removes "tags" edges to Tag entities.
func (au *AgentUpdate) RemoveTags(t ...*Tag) *AgentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AgentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AgentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AgentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AgentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AgentUpdate) check() error {
	if v, ok := au.mutation.Hostname(); ok {
		if err := agent.HostnameValidator(v); err != nil {
			return &ValidationError{Name: "hostname", err: fmt.Errorf(`ent: validator failed for field "Agent.hostname": %w`, err)}
		}
	}
	return nil
}

func (au *AgentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(agent.Table, agent.Columns, sqlgraph.NewFieldSpec(agent.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Hostname(); ok {
		_spec.SetField(agent.FieldHostname, field.TypeString, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(agent.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   agent.TagsTable,
			Columns: agent.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTagsIDs(); len(nodes) > 0 && !au.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   agent.TagsTable,
			Columns: agent.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   agent.TagsTable,
			Columns: agent.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AgentUpdateOne is the builder for updating a single Agent entity.
type AgentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AgentMutation
}

// SetHostname sets the "hostname" field.
func (auo *AgentUpdateOne) SetHostname(s string) *AgentUpdateOne {
	auo.mutation.SetHostname(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AgentUpdateOne) SetCreatedAt(t time.Time) *AgentUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AgentUpdateOne) SetNillableCreatedAt(t *time.Time) *AgentUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (auo *AgentUpdateOne) AddTagIDs(ids ...int) *AgentUpdateOne {
	auo.mutation.AddTagIDs(ids...)
	return auo
}

// AddTags adds the "tags" edges to the Tag entity.
func (auo *AgentUpdateOne) AddTags(t ...*Tag) *AgentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTagIDs(ids...)
}

// Mutation returns the AgentMutation object of the builder.
func (auo *AgentUpdateOne) Mutation() *AgentMutation {
	return auo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (auo *AgentUpdateOne) ClearTags() *AgentUpdateOne {
	auo.mutation.ClearTags()
	return auo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (auo *AgentUpdateOne) RemoveTagIDs(ids ...int) *AgentUpdateOne {
	auo.mutation.RemoveTagIDs(ids...)
	return auo
}

// RemoveTags removes "tags" edges to Tag entities.
func (auo *AgentUpdateOne) RemoveTags(t ...*Tag) *AgentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTagIDs(ids...)
}

// Where appends a list predicates to the AgentUpdate builder.
func (auo *AgentUpdateOne) Where(ps ...predicate.Agent) *AgentUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AgentUpdateOne) Select(field string, fields ...string) *AgentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Agent entity.
func (auo *AgentUpdateOne) Save(ctx context.Context) (*Agent, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AgentUpdateOne) SaveX(ctx context.Context) *Agent {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AgentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AgentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AgentUpdateOne) check() error {
	if v, ok := auo.mutation.Hostname(); ok {
		if err := agent.HostnameValidator(v); err != nil {
			return &ValidationError{Name: "hostname", err: fmt.Errorf(`ent: validator failed for field "Agent.hostname": %w`, err)}
		}
	}
	return nil
}

func (auo *AgentUpdateOne) sqlSave(ctx context.Context) (_node *Agent, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(agent.Table, agent.Columns, sqlgraph.NewFieldSpec(agent.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Agent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, agent.FieldID)
		for _, f := range fields {
			if !agent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != agent.FieldID {
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
	if value, ok := auo.mutation.Hostname(); ok {
		_spec.SetField(agent.FieldHostname, field.TypeString, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(agent.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   agent.TagsTable,
			Columns: agent.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !auo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   agent.TagsTable,
			Columns: agent.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   agent.TagsTable,
			Columns: agent.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Agent{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
