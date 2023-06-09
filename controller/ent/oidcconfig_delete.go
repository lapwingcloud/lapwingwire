// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lapwingcloud/lapwingwire/controller/ent/oidcconfig"
	"github.com/lapwingcloud/lapwingwire/controller/ent/predicate"
)

// OIDCConfigDelete is the builder for deleting a OIDCConfig entity.
type OIDCConfigDelete struct {
	config
	hooks    []Hook
	mutation *OIDCConfigMutation
}

// Where appends a list predicates to the OIDCConfigDelete builder.
func (ocd *OIDCConfigDelete) Where(ps ...predicate.OIDCConfig) *OIDCConfigDelete {
	ocd.mutation.Where(ps...)
	return ocd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ocd *OIDCConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ocd.sqlExec, ocd.mutation, ocd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ocd *OIDCConfigDelete) ExecX(ctx context.Context) int {
	n, err := ocd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ocd *OIDCConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(oidcconfig.Table, sqlgraph.NewFieldSpec(oidcconfig.FieldID, field.TypeInt))
	if ps := ocd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ocd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ocd.mutation.done = true
	return affected, err
}

// OIDCConfigDeleteOne is the builder for deleting a single OIDCConfig entity.
type OIDCConfigDeleteOne struct {
	ocd *OIDCConfigDelete
}

// Where appends a list predicates to the OIDCConfigDelete builder.
func (ocdo *OIDCConfigDeleteOne) Where(ps ...predicate.OIDCConfig) *OIDCConfigDeleteOne {
	ocdo.ocd.mutation.Where(ps...)
	return ocdo
}

// Exec executes the deletion query.
func (ocdo *OIDCConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := ocdo.ocd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{oidcconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdo *OIDCConfigDeleteOne) ExecX(ctx context.Context) {
	if err := ocdo.Exec(ctx); err != nil {
		panic(err)
	}
}
