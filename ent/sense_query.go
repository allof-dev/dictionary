// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/allof-dev/dictionary/ent/lemma"
	"github.com/allof-dev/dictionary/ent/predicate"
	"github.com/allof-dev/dictionary/ent/sense"
	"github.com/allof-dev/dictionary/ent/synset"
)

// SenseQuery is the builder for querying Sense entities.
type SenseQuery struct {
	config
	ctx        *QueryContext
	order      []sense.OrderOption
	inters     []Interceptor
	predicates []predicate.Sense
	withSynset *SynsetQuery
	withLemma  *LemmaQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SenseQuery builder.
func (sq *SenseQuery) Where(ps ...predicate.Sense) *SenseQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SenseQuery) Limit(limit int) *SenseQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SenseQuery) Offset(offset int) *SenseQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SenseQuery) Unique(unique bool) *SenseQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SenseQuery) Order(o ...sense.OrderOption) *SenseQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QuerySynset chains the current query on the "synset" edge.
func (sq *SenseQuery) QuerySynset() *SynsetQuery {
	query := (&SynsetClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sense.Table, sense.FieldID, selector),
			sqlgraph.To(synset.Table, synset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, sense.SynsetTable, sense.SynsetColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLemma chains the current query on the "lemma" edge.
func (sq *SenseQuery) QueryLemma() *LemmaQuery {
	query := (&LemmaClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sense.Table, sense.FieldID, selector),
			sqlgraph.To(lemma.Table, lemma.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, sense.LemmaTable, sense.LemmaColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Sense entity from the query.
// Returns a *NotFoundError when no Sense was found.
func (sq *SenseQuery) First(ctx context.Context) (*Sense, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sense.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SenseQuery) FirstX(ctx context.Context) *Sense {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Sense ID from the query.
// Returns a *NotFoundError when no Sense ID was found.
func (sq *SenseQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sense.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SenseQuery) FirstIDX(ctx context.Context) string {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Sense entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Sense entity is found.
// Returns a *NotFoundError when no Sense entities are found.
func (sq *SenseQuery) Only(ctx context.Context) (*Sense, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sense.Label}
	default:
		return nil, &NotSingularError{sense.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SenseQuery) OnlyX(ctx context.Context) *Sense {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Sense ID in the query.
// Returns a *NotSingularError when more than one Sense ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SenseQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sense.Label}
	default:
		err = &NotSingularError{sense.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SenseQuery) OnlyIDX(ctx context.Context) string {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Senses.
func (sq *SenseQuery) All(ctx context.Context) ([]*Sense, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryAll)
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Sense, *SenseQuery]()
	return withInterceptors[[]*Sense](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SenseQuery) AllX(ctx context.Context) []*Sense {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Sense IDs.
func (sq *SenseQuery) IDs(ctx context.Context) (ids []string, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryIDs)
	if err = sq.Select(sense.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SenseQuery) IDsX(ctx context.Context) []string {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SenseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryCount)
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SenseQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SenseQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SenseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryExist)
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SenseQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SenseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SenseQuery) Clone() *SenseQuery {
	if sq == nil {
		return nil
	}
	return &SenseQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]sense.OrderOption{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Sense{}, sq.predicates...),
		withSynset: sq.withSynset.Clone(),
		withLemma:  sq.withLemma.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithSynset tells the query-builder to eager-load the nodes that are connected to
// the "synset" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SenseQuery) WithSynset(opts ...func(*SynsetQuery)) *SenseQuery {
	query := (&SynsetClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withSynset = query
	return sq
}

// WithLemma tells the query-builder to eager-load the nodes that are connected to
// the "lemma" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SenseQuery) WithLemma(opts ...func(*LemmaQuery)) *SenseQuery {
	query := (&LemmaClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withLemma = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (sq *SenseQuery) GroupBy(field string, fields ...string) *SenseGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SenseGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = sense.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (sq *SenseQuery) Select(fields ...string) *SenseSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SenseSelect{SenseQuery: sq}
	sbuild.label = sense.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SenseSelect configured with the given aggregations.
func (sq *SenseQuery) Aggregate(fns ...AggregateFunc) *SenseSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SenseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !sense.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SenseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Sense, error) {
	var (
		nodes       = []*Sense{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [2]bool{
			sq.withSynset != nil,
			sq.withLemma != nil,
		}
	)
	if sq.withSynset != nil || sq.withLemma != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, sense.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Sense).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Sense{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withSynset; query != nil {
		if err := sq.loadSynset(ctx, query, nodes, nil,
			func(n *Sense, e *Synset) { n.Edges.Synset = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withLemma; query != nil {
		if err := sq.loadLemma(ctx, query, nodes, nil,
			func(n *Sense, e *Lemma) { n.Edges.Lemma = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SenseQuery) loadSynset(ctx context.Context, query *SynsetQuery, nodes []*Sense, init func(*Sense), assign func(*Sense, *Synset)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Sense)
	for i := range nodes {
		if nodes[i].sense_synset == nil {
			continue
		}
		fk := *nodes[i].sense_synset
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(synset.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "sense_synset" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *SenseQuery) loadLemma(ctx context.Context, query *LemmaQuery, nodes []*Sense, init func(*Sense), assign func(*Sense, *Lemma)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Sense)
	for i := range nodes {
		if nodes[i].sense_lemma == nil {
			continue
		}
		fk := *nodes[i].sense_lemma
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(lemma.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "sense_lemma" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SenseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SenseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sense.Table, sense.Columns, sqlgraph.NewFieldSpec(sense.FieldID, field.TypeString))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sense.FieldID)
		for i := range fields {
			if fields[i] != sense.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SenseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(sense.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = sense.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SenseGroupBy is the group-by builder for Sense entities.
type SenseGroupBy struct {
	selector
	build *SenseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SenseGroupBy) Aggregate(fns ...AggregateFunc) *SenseGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SenseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, ent.OpQueryGroupBy)
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SenseQuery, *SenseGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SenseGroupBy) sqlScan(ctx context.Context, root *SenseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SenseSelect is the builder for selecting fields of Sense entities.
type SenseSelect struct {
	*SenseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SenseSelect) Aggregate(fns ...AggregateFunc) *SenseSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SenseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, ent.OpQuerySelect)
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SenseQuery, *SenseSelect](ctx, ss.SenseQuery, ss, ss.inters, v)
}

func (ss *SenseSelect) sqlScan(ctx context.Context, root *SenseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
