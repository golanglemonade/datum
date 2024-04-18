// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/integrationhistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// IntegrationHistoryQuery is the builder for querying IntegrationHistory entities.
type IntegrationHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []integrationhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.IntegrationHistory
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*IntegrationHistory) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IntegrationHistoryQuery builder.
func (ihq *IntegrationHistoryQuery) Where(ps ...predicate.IntegrationHistory) *IntegrationHistoryQuery {
	ihq.predicates = append(ihq.predicates, ps...)
	return ihq
}

// Limit the number of records to be returned by this query.
func (ihq *IntegrationHistoryQuery) Limit(limit int) *IntegrationHistoryQuery {
	ihq.ctx.Limit = &limit
	return ihq
}

// Offset to start from.
func (ihq *IntegrationHistoryQuery) Offset(offset int) *IntegrationHistoryQuery {
	ihq.ctx.Offset = &offset
	return ihq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ihq *IntegrationHistoryQuery) Unique(unique bool) *IntegrationHistoryQuery {
	ihq.ctx.Unique = &unique
	return ihq
}

// Order specifies how the records should be ordered.
func (ihq *IntegrationHistoryQuery) Order(o ...integrationhistory.OrderOption) *IntegrationHistoryQuery {
	ihq.order = append(ihq.order, o...)
	return ihq
}

// First returns the first IntegrationHistory entity from the query.
// Returns a *NotFoundError when no IntegrationHistory was found.
func (ihq *IntegrationHistoryQuery) First(ctx context.Context) (*IntegrationHistory, error) {
	nodes, err := ihq.Limit(1).All(setContextOp(ctx, ihq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{integrationhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ihq *IntegrationHistoryQuery) FirstX(ctx context.Context) *IntegrationHistory {
	node, err := ihq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IntegrationHistory ID from the query.
// Returns a *NotFoundError when no IntegrationHistory ID was found.
func (ihq *IntegrationHistoryQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ihq.Limit(1).IDs(setContextOp(ctx, ihq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{integrationhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ihq *IntegrationHistoryQuery) FirstIDX(ctx context.Context) string {
	id, err := ihq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IntegrationHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IntegrationHistory entity is found.
// Returns a *NotFoundError when no IntegrationHistory entities are found.
func (ihq *IntegrationHistoryQuery) Only(ctx context.Context) (*IntegrationHistory, error) {
	nodes, err := ihq.Limit(2).All(setContextOp(ctx, ihq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{integrationhistory.Label}
	default:
		return nil, &NotSingularError{integrationhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ihq *IntegrationHistoryQuery) OnlyX(ctx context.Context) *IntegrationHistory {
	node, err := ihq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IntegrationHistory ID in the query.
// Returns a *NotSingularError when more than one IntegrationHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ihq *IntegrationHistoryQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ihq.Limit(2).IDs(setContextOp(ctx, ihq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{integrationhistory.Label}
	default:
		err = &NotSingularError{integrationhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ihq *IntegrationHistoryQuery) OnlyIDX(ctx context.Context) string {
	id, err := ihq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IntegrationHistories.
func (ihq *IntegrationHistoryQuery) All(ctx context.Context) ([]*IntegrationHistory, error) {
	ctx = setContextOp(ctx, ihq.ctx, "All")
	if err := ihq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*IntegrationHistory, *IntegrationHistoryQuery]()
	return withInterceptors[[]*IntegrationHistory](ctx, ihq, qr, ihq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ihq *IntegrationHistoryQuery) AllX(ctx context.Context) []*IntegrationHistory {
	nodes, err := ihq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IntegrationHistory IDs.
func (ihq *IntegrationHistoryQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ihq.ctx.Unique == nil && ihq.path != nil {
		ihq.Unique(true)
	}
	ctx = setContextOp(ctx, ihq.ctx, "IDs")
	if err = ihq.Select(integrationhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ihq *IntegrationHistoryQuery) IDsX(ctx context.Context) []string {
	ids, err := ihq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ihq *IntegrationHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ihq.ctx, "Count")
	if err := ihq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ihq, querierCount[*IntegrationHistoryQuery](), ihq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ihq *IntegrationHistoryQuery) CountX(ctx context.Context) int {
	count, err := ihq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ihq *IntegrationHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ihq.ctx, "Exist")
	switch _, err := ihq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ihq *IntegrationHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ihq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IntegrationHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ihq *IntegrationHistoryQuery) Clone() *IntegrationHistoryQuery {
	if ihq == nil {
		return nil
	}
	return &IntegrationHistoryQuery{
		config:     ihq.config,
		ctx:        ihq.ctx.Clone(),
		order:      append([]integrationhistory.OrderOption{}, ihq.order...),
		inters:     append([]Interceptor{}, ihq.inters...),
		predicates: append([]predicate.IntegrationHistory{}, ihq.predicates...),
		// clone intermediate query.
		sql:  ihq.sql.Clone(),
		path: ihq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IntegrationHistory.Query().
//		GroupBy(integrationhistory.FieldHistoryTime).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (ihq *IntegrationHistoryQuery) GroupBy(field string, fields ...string) *IntegrationHistoryGroupBy {
	ihq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IntegrationHistoryGroupBy{build: ihq}
	grbuild.flds = &ihq.ctx.Fields
	grbuild.label = integrationhistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//	}
//
//	client.IntegrationHistory.Query().
//		Select(integrationhistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (ihq *IntegrationHistoryQuery) Select(fields ...string) *IntegrationHistorySelect {
	ihq.ctx.Fields = append(ihq.ctx.Fields, fields...)
	sbuild := &IntegrationHistorySelect{IntegrationHistoryQuery: ihq}
	sbuild.label = integrationhistory.Label
	sbuild.flds, sbuild.scan = &ihq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IntegrationHistorySelect configured with the given aggregations.
func (ihq *IntegrationHistoryQuery) Aggregate(fns ...AggregateFunc) *IntegrationHistorySelect {
	return ihq.Select().Aggregate(fns...)
}

func (ihq *IntegrationHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ihq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ihq); err != nil {
				return err
			}
		}
	}
	for _, f := range ihq.ctx.Fields {
		if !integrationhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if ihq.path != nil {
		prev, err := ihq.path(ctx)
		if err != nil {
			return err
		}
		ihq.sql = prev
	}
	return nil
}

func (ihq *IntegrationHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IntegrationHistory, error) {
	var (
		nodes = []*IntegrationHistory{}
		_spec = ihq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*IntegrationHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &IntegrationHistory{config: ihq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = ihq.schemaConfig.IntegrationHistory
	ctx = internal.NewSchemaConfigContext(ctx, ihq.schemaConfig)
	if len(ihq.modifiers) > 0 {
		_spec.Modifiers = ihq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ihq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range ihq.loadTotal {
		if err := ihq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ihq *IntegrationHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ihq.querySpec()
	_spec.Node.Schema = ihq.schemaConfig.IntegrationHistory
	ctx = internal.NewSchemaConfigContext(ctx, ihq.schemaConfig)
	if len(ihq.modifiers) > 0 {
		_spec.Modifiers = ihq.modifiers
	}
	_spec.Node.Columns = ihq.ctx.Fields
	if len(ihq.ctx.Fields) > 0 {
		_spec.Unique = ihq.ctx.Unique != nil && *ihq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ihq.driver, _spec)
}

func (ihq *IntegrationHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(integrationhistory.Table, integrationhistory.Columns, sqlgraph.NewFieldSpec(integrationhistory.FieldID, field.TypeString))
	_spec.From = ihq.sql
	if unique := ihq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ihq.path != nil {
		_spec.Unique = true
	}
	if fields := ihq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, integrationhistory.FieldID)
		for i := range fields {
			if fields[i] != integrationhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ihq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ihq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ihq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ihq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ihq *IntegrationHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ihq.driver.Dialect())
	t1 := builder.Table(integrationhistory.Table)
	columns := ihq.ctx.Fields
	if len(columns) == 0 {
		columns = integrationhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ihq.sql != nil {
		selector = ihq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ihq.ctx.Unique != nil && *ihq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(ihq.schemaConfig.IntegrationHistory)
	ctx = internal.NewSchemaConfigContext(ctx, ihq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range ihq.predicates {
		p(selector)
	}
	for _, p := range ihq.order {
		p(selector)
	}
	if offset := ihq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ihq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IntegrationHistoryGroupBy is the group-by builder for IntegrationHistory entities.
type IntegrationHistoryGroupBy struct {
	selector
	build *IntegrationHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ihgb *IntegrationHistoryGroupBy) Aggregate(fns ...AggregateFunc) *IntegrationHistoryGroupBy {
	ihgb.fns = append(ihgb.fns, fns...)
	return ihgb
}

// Scan applies the selector query and scans the result into the given value.
func (ihgb *IntegrationHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ihgb.build.ctx, "GroupBy")
	if err := ihgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IntegrationHistoryQuery, *IntegrationHistoryGroupBy](ctx, ihgb.build, ihgb, ihgb.build.inters, v)
}

func (ihgb *IntegrationHistoryGroupBy) sqlScan(ctx context.Context, root *IntegrationHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ihgb.fns))
	for _, fn := range ihgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ihgb.flds)+len(ihgb.fns))
		for _, f := range *ihgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ihgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ihgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IntegrationHistorySelect is the builder for selecting fields of IntegrationHistory entities.
type IntegrationHistorySelect struct {
	*IntegrationHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ihs *IntegrationHistorySelect) Aggregate(fns ...AggregateFunc) *IntegrationHistorySelect {
	ihs.fns = append(ihs.fns, fns...)
	return ihs
}

// Scan applies the selector query and scans the result into the given value.
func (ihs *IntegrationHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ihs.ctx, "Select")
	if err := ihs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IntegrationHistoryQuery, *IntegrationHistorySelect](ctx, ihs.IntegrationHistoryQuery, ihs, ihs.inters, v)
}

func (ihs *IntegrationHistorySelect) sqlScan(ctx context.Context, root *IntegrationHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ihs.fns))
	for _, fn := range ihs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ihs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ihs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}