// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lapwingcloud/lapwingwire/controller/ent/agent"
	"github.com/lapwingcloud/lapwingwire/controller/ent/predicate"
)

// AgentDelete is the builder for deleting a Agent entity.
type AgentDelete struct {
	config
	hooks    []Hook
	mutation *AgentMutation
}

// Where appends a list predicates to the AgentDelete builder.
func (ad *AgentDelete) Where(ps ...predicate.Agent) *AgentDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AgentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AgentDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AgentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(agent.Table, sqlgraph.NewFieldSpec(agent.FieldID, field.TypeInt))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AgentDeleteOne is the builder for deleting a single Agent entity.
type AgentDeleteOne struct {
	ad *AgentDelete
}

// Where appends a list predicates to the AgentDelete builder.
func (ado *AgentDeleteOne) Where(ps ...predicate.Agent) *AgentDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AgentDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{agent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AgentDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
