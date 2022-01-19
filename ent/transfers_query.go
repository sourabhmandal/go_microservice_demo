// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gomicro/ent/accounts"
	"gomicro/ent/predicate"
	"gomicro/ent/transfers"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransfersQuery is the builder for querying Transfers entities.
type TransfersQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Transfers
	// eager-loading edges.
	withAccountOf *AccountsQuery
	withAccountTo *AccountsQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TransfersQuery builder.
func (tq *TransfersQuery) Where(ps ...predicate.Transfers) *TransfersQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TransfersQuery) Limit(limit int) *TransfersQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TransfersQuery) Offset(offset int) *TransfersQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TransfersQuery) Unique(unique bool) *TransfersQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TransfersQuery) Order(o ...OrderFunc) *TransfersQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryAccountOf chains the current query on the "account_of" edge.
func (tq *TransfersQuery) QueryAccountOf() *AccountsQuery {
	query := &AccountsQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transfers.Table, transfers.FieldID, selector),
			sqlgraph.To(accounts.Table, accounts.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transfers.AccountOfTable, transfers.AccountOfColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccountTo chains the current query on the "account_to" edge.
func (tq *TransfersQuery) QueryAccountTo() *AccountsQuery {
	query := &AccountsQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transfers.Table, transfers.FieldID, selector),
			sqlgraph.To(accounts.Table, accounts.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transfers.AccountToTable, transfers.AccountToColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Transfers entity from the query.
// Returns a *NotFoundError when no Transfers was found.
func (tq *TransfersQuery) First(ctx context.Context) (*Transfers, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transfers.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TransfersQuery) FirstX(ctx context.Context) *Transfers {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Transfers ID from the query.
// Returns a *NotFoundError when no Transfers ID was found.
func (tq *TransfersQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transfers.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TransfersQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Transfers entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Transfers entity is not found.
// Returns a *NotFoundError when no Transfers entities are found.
func (tq *TransfersQuery) Only(ctx context.Context) (*Transfers, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transfers.Label}
	default:
		return nil, &NotSingularError{transfers.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TransfersQuery) OnlyX(ctx context.Context) *Transfers {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Transfers ID in the query.
// Returns a *NotSingularError when exactly one Transfers ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tq *TransfersQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transfers.Label}
	default:
		err = &NotSingularError{transfers.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TransfersQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TransfersSlice.
func (tq *TransfersQuery) All(ctx context.Context) ([]*Transfers, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TransfersQuery) AllX(ctx context.Context) []*Transfers {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Transfers IDs.
func (tq *TransfersQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(transfers.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TransfersQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TransfersQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TransfersQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TransfersQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TransfersQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TransfersQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TransfersQuery) Clone() *TransfersQuery {
	if tq == nil {
		return nil
	}
	return &TransfersQuery{
		config:        tq.config,
		limit:         tq.limit,
		offset:        tq.offset,
		order:         append([]OrderFunc{}, tq.order...),
		predicates:    append([]predicate.Transfers{}, tq.predicates...),
		withAccountOf: tq.withAccountOf.Clone(),
		withAccountTo: tq.withAccountTo.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithAccountOf tells the query-builder to eager-load the nodes that are connected to
// the "account_of" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TransfersQuery) WithAccountOf(opts ...func(*AccountsQuery)) *TransfersQuery {
	query := &AccountsQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withAccountOf = query
	return tq
}

// WithAccountTo tells the query-builder to eager-load the nodes that are connected to
// the "account_to" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TransfersQuery) WithAccountTo(opts ...func(*AccountsQuery)) *TransfersQuery {
	query := &AccountsQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withAccountTo = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FromAccountID int `json:"from_account_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Transfers.Query().
//		GroupBy(transfers.FieldFromAccountID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TransfersQuery) GroupBy(field string, fields ...string) *TransfersGroupBy {
	group := &TransfersGroupBy{config: tq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FromAccountID int `json:"from_account_id,omitempty"`
//	}
//
//	client.Transfers.Query().
//		Select(transfers.FieldFromAccountID).
//		Scan(ctx, &v)
//
func (tq *TransfersQuery) Select(fields ...string) *TransfersSelect {
	tq.fields = append(tq.fields, fields...)
	return &TransfersSelect{TransfersQuery: tq}
}

func (tq *TransfersQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !transfers.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TransfersQuery) sqlAll(ctx context.Context) ([]*Transfers, error) {
	var (
		nodes       = []*Transfers{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withAccountOf != nil,
			tq.withAccountTo != nil,
		}
	)
	if tq.withAccountOf != nil || tq.withAccountTo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, transfers.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Transfers{config: tq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tq.withAccountOf; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Transfers)
		for i := range nodes {
			if nodes[i].accounts_from_account_id == nil {
				continue
			}
			fk := *nodes[i].accounts_from_account_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(accounts.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "accounts_from_account_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.AccountOf = n
			}
		}
	}

	if query := tq.withAccountTo; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Transfers)
		for i := range nodes {
			if nodes[i].accounts_to_account_id == nil {
				continue
			}
			fk := *nodes[i].accounts_to_account_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(accounts.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "accounts_to_account_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.AccountTo = n
			}
		}
	}

	return nodes, nil
}

func (tq *TransfersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TransfersQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tq *TransfersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transfers.Table,
			Columns: transfers.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transfers.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transfers.FieldID)
		for i := range fields {
			if fields[i] != transfers.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TransfersQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(transfers.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = transfers.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TransfersGroupBy is the group-by builder for Transfers entities.
type TransfersGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TransfersGroupBy) Aggregate(fns ...AggregateFunc) *TransfersGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TransfersGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tgb *TransfersGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransfersGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransfersGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tgb *TransfersGroupBy) StringsX(ctx context.Context) []string {
	v, err := tgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransfersGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfers.Label}
	default:
		err = fmt.Errorf("ent: TransfersGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tgb *TransfersGroupBy) StringX(ctx context.Context) string {
	v, err := tgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransfersGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransfersGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tgb *TransfersGroupBy) IntsX(ctx context.Context) []int {
	v, err := tgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransfersGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfers.Label}
	default:
		err = fmt.Errorf("ent: TransfersGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tgb *TransfersGroupBy) IntX(ctx context.Context) int {
	v, err := tgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransfersGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransfersGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tgb *TransfersGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransfersGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfers.Label}
	default:
		err = fmt.Errorf("ent: TransfersGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tgb *TransfersGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransfersGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TransfersGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tgb *TransfersGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TransfersGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfers.Label}
	default:
		err = fmt.Errorf("ent: TransfersGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tgb *TransfersGroupBy) BoolX(ctx context.Context) bool {
	v, err := tgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tgb *TransfersGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tgb.fields {
		if !transfers.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TransfersGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TransfersSelect is the builder for selecting fields of Transfers entities.
type TransfersSelect struct {
	*TransfersQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TransfersSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TransfersQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ts *TransfersSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ts *TransfersSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransfersSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ts *TransfersSelect) StringsX(ctx context.Context) []string {
	v, err := ts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ts *TransfersSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfers.Label}
	default:
		err = fmt.Errorf("ent: TransfersSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ts *TransfersSelect) StringX(ctx context.Context) string {
	v, err := ts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ts *TransfersSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransfersSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ts *TransfersSelect) IntsX(ctx context.Context) []int {
	v, err := ts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ts *TransfersSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfers.Label}
	default:
		err = fmt.Errorf("ent: TransfersSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ts *TransfersSelect) IntX(ctx context.Context) int {
	v, err := ts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ts *TransfersSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransfersSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ts *TransfersSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ts *TransfersSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfers.Label}
	default:
		err = fmt.Errorf("ent: TransfersSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ts *TransfersSelect) Float64X(ctx context.Context) float64 {
	v, err := ts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ts *TransfersSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TransfersSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ts *TransfersSelect) BoolsX(ctx context.Context) []bool {
	v, err := ts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ts *TransfersSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{transfers.Label}
	default:
		err = fmt.Errorf("ent: TransfersSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ts *TransfersSelect) BoolX(ctx context.Context) bool {
	v, err := ts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ts *TransfersSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
