// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lapwingcloud/lapwingwire/controller/ent/oidcconfig"
	"github.com/lapwingcloud/lapwingwire/controller/ent/predicate"
)

// OIDCConfigQuery is the builder for querying OIDCConfig entities.
type OIDCConfigQuery struct {
	config
	ctx        *QueryContext
	order      []oidcconfig.OrderOption
	inters     []Interceptor
	predicates []predicate.OIDCConfig
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OIDCConfigQuery builder.
func (ocq *OIDCConfigQuery) Where(ps ...predicate.OIDCConfig) *OIDCConfigQuery {
	ocq.predicates = append(ocq.predicates, ps...)
	return ocq
}

// Limit the number of records to be returned by this query.
func (ocq *OIDCConfigQuery) Limit(limit int) *OIDCConfigQuery {
	ocq.ctx.Limit = &limit
	return ocq
}

// Offset to start from.
func (ocq *OIDCConfigQuery) Offset(offset int) *OIDCConfigQuery {
	ocq.ctx.Offset = &offset
	return ocq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ocq *OIDCConfigQuery) Unique(unique bool) *OIDCConfigQuery {
	ocq.ctx.Unique = &unique
	return ocq
}

// Order specifies how the records should be ordered.
func (ocq *OIDCConfigQuery) Order(o ...oidcconfig.OrderOption) *OIDCConfigQuery {
	ocq.order = append(ocq.order, o...)
	return ocq
}

// First returns the first OIDCConfig entity from the query.
// Returns a *NotFoundError when no OIDCConfig was found.
func (ocq *OIDCConfigQuery) First(ctx context.Context) (*OIDCConfig, error) {
	nodes, err := ocq.Limit(1).All(setContextOp(ctx, ocq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oidcconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ocq *OIDCConfigQuery) FirstX(ctx context.Context) *OIDCConfig {
	node, err := ocq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OIDCConfig ID from the query.
// Returns a *NotFoundError when no OIDCConfig ID was found.
func (ocq *OIDCConfigQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocq.Limit(1).IDs(setContextOp(ctx, ocq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oidcconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ocq *OIDCConfigQuery) FirstIDX(ctx context.Context) int {
	id, err := ocq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OIDCConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OIDCConfig entity is found.
// Returns a *NotFoundError when no OIDCConfig entities are found.
func (ocq *OIDCConfigQuery) Only(ctx context.Context) (*OIDCConfig, error) {
	nodes, err := ocq.Limit(2).All(setContextOp(ctx, ocq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oidcconfig.Label}
	default:
		return nil, &NotSingularError{oidcconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ocq *OIDCConfigQuery) OnlyX(ctx context.Context) *OIDCConfig {
	node, err := ocq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OIDCConfig ID in the query.
// Returns a *NotSingularError when more than one OIDCConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (ocq *OIDCConfigQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ocq.Limit(2).IDs(setContextOp(ctx, ocq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oidcconfig.Label}
	default:
		err = &NotSingularError{oidcconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ocq *OIDCConfigQuery) OnlyIDX(ctx context.Context) int {
	id, err := ocq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OIDCConfigs.
func (ocq *OIDCConfigQuery) All(ctx context.Context) ([]*OIDCConfig, error) {
	ctx = setContextOp(ctx, ocq.ctx, "All")
	if err := ocq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OIDCConfig, *OIDCConfigQuery]()
	return withInterceptors[[]*OIDCConfig](ctx, ocq, qr, ocq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ocq *OIDCConfigQuery) AllX(ctx context.Context) []*OIDCConfig {
	nodes, err := ocq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OIDCConfig IDs.
func (ocq *OIDCConfigQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ocq.ctx.Unique == nil && ocq.path != nil {
		ocq.Unique(true)
	}
	ctx = setContextOp(ctx, ocq.ctx, "IDs")
	if err = ocq.Select(oidcconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ocq *OIDCConfigQuery) IDsX(ctx context.Context) []int {
	ids, err := ocq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ocq *OIDCConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ocq.ctx, "Count")
	if err := ocq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ocq, querierCount[*OIDCConfigQuery](), ocq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ocq *OIDCConfigQuery) CountX(ctx context.Context) int {
	count, err := ocq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ocq *OIDCConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ocq.ctx, "Exist")
	switch _, err := ocq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ocq *OIDCConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := ocq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OIDCConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ocq *OIDCConfigQuery) Clone() *OIDCConfigQuery {
	if ocq == nil {
		return nil
	}
	return &OIDCConfigQuery{
		config:     ocq.config,
		ctx:        ocq.ctx.Clone(),
		order:      append([]oidcconfig.OrderOption{}, ocq.order...),
		inters:     append([]Interceptor{}, ocq.inters...),
		predicates: append([]predicate.OIDCConfig{}, ocq.predicates...),
		// clone intermediate query.
		sql:  ocq.sql.Clone(),
		path: ocq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProviderKey string `json:"provider_key,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OIDCConfig.Query().
//		GroupBy(oidcconfig.FieldProviderKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ocq *OIDCConfigQuery) GroupBy(field string, fields ...string) *OIDCConfigGroupBy {
	ocq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OIDCConfigGroupBy{build: ocq}
	grbuild.flds = &ocq.ctx.Fields
	grbuild.label = oidcconfig.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProviderKey string `json:"provider_key,omitempty"`
//	}
//
//	client.OIDCConfig.Query().
//		Select(oidcconfig.FieldProviderKey).
//		Scan(ctx, &v)
func (ocq *OIDCConfigQuery) Select(fields ...string) *OIDCConfigSelect {
	ocq.ctx.Fields = append(ocq.ctx.Fields, fields...)
	sbuild := &OIDCConfigSelect{OIDCConfigQuery: ocq}
	sbuild.label = oidcconfig.Label
	sbuild.flds, sbuild.scan = &ocq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OIDCConfigSelect configured with the given aggregations.
func (ocq *OIDCConfigQuery) Aggregate(fns ...AggregateFunc) *OIDCConfigSelect {
	return ocq.Select().Aggregate(fns...)
}

func (ocq *OIDCConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ocq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ocq); err != nil {
				return err
			}
		}
	}
	for _, f := range ocq.ctx.Fields {
		if !oidcconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ocq.path != nil {
		prev, err := ocq.path(ctx)
		if err != nil {
			return err
		}
		ocq.sql = prev
	}
	return nil
}

func (ocq *OIDCConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OIDCConfig, error) {
	var (
		nodes = []*OIDCConfig{}
		_spec = ocq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OIDCConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OIDCConfig{config: ocq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ocq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ocq *OIDCConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ocq.querySpec()
	_spec.Node.Columns = ocq.ctx.Fields
	if len(ocq.ctx.Fields) > 0 {
		_spec.Unique = ocq.ctx.Unique != nil && *ocq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ocq.driver, _spec)
}

func (ocq *OIDCConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(oidcconfig.Table, oidcconfig.Columns, sqlgraph.NewFieldSpec(oidcconfig.FieldID, field.TypeInt))
	_spec.From = ocq.sql
	if unique := ocq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ocq.path != nil {
		_spec.Unique = true
	}
	if fields := ocq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oidcconfig.FieldID)
		for i := range fields {
			if fields[i] != oidcconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ocq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ocq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ocq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ocq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ocq *OIDCConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ocq.driver.Dialect())
	t1 := builder.Table(oidcconfig.Table)
	columns := ocq.ctx.Fields
	if len(columns) == 0 {
		columns = oidcconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ocq.sql != nil {
		selector = ocq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ocq.ctx.Unique != nil && *ocq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ocq.predicates {
		p(selector)
	}
	for _, p := range ocq.order {
		p(selector)
	}
	if offset := ocq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ocq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OIDCConfigGroupBy is the group-by builder for OIDCConfig entities.
type OIDCConfigGroupBy struct {
	selector
	build *OIDCConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ocgb *OIDCConfigGroupBy) Aggregate(fns ...AggregateFunc) *OIDCConfigGroupBy {
	ocgb.fns = append(ocgb.fns, fns...)
	return ocgb
}

// Scan applies the selector query and scans the result into the given value.
func (ocgb *OIDCConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ocgb.build.ctx, "GroupBy")
	if err := ocgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OIDCConfigQuery, *OIDCConfigGroupBy](ctx, ocgb.build, ocgb, ocgb.build.inters, v)
}

func (ocgb *OIDCConfigGroupBy) sqlScan(ctx context.Context, root *OIDCConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ocgb.fns))
	for _, fn := range ocgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ocgb.flds)+len(ocgb.fns))
		for _, f := range *ocgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ocgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ocgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OIDCConfigSelect is the builder for selecting fields of OIDCConfig entities.
type OIDCConfigSelect struct {
	*OIDCConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ocs *OIDCConfigSelect) Aggregate(fns ...AggregateFunc) *OIDCConfigSelect {
	ocs.fns = append(ocs.fns, fns...)
	return ocs
}

// Scan applies the selector query and scans the result into the given value.
func (ocs *OIDCConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ocs.ctx, "Select")
	if err := ocs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OIDCConfigQuery, *OIDCConfigSelect](ctx, ocs.OIDCConfigQuery, ocs, ocs.inters, v)
}

func (ocs *OIDCConfigSelect) sqlScan(ctx context.Context, root *OIDCConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ocs.fns))
	for _, fn := range ocs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ocs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ocs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}