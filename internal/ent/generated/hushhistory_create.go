// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/hushhistory"
	"github.com/datumforge/enthistory"
)

// HushHistoryCreate is the builder for creating a HushHistory entity.
type HushHistoryCreate struct {
	config
	mutation *HushHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (hhc *HushHistoryCreate) SetHistoryTime(t time.Time) *HushHistoryCreate {
	hhc.mutation.SetHistoryTime(t)
	return hhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableHistoryTime(t *time.Time) *HushHistoryCreate {
	if t != nil {
		hhc.SetHistoryTime(*t)
	}
	return hhc
}

// SetOperation sets the "operation" field.
func (hhc *HushHistoryCreate) SetOperation(et enthistory.OpType) *HushHistoryCreate {
	hhc.mutation.SetOperation(et)
	return hhc
}

// SetRef sets the "ref" field.
func (hhc *HushHistoryCreate) SetRef(s string) *HushHistoryCreate {
	hhc.mutation.SetRef(s)
	return hhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableRef(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetRef(*s)
	}
	return hhc
}

// SetCreatedAt sets the "created_at" field.
func (hhc *HushHistoryCreate) SetCreatedAt(t time.Time) *HushHistoryCreate {
	hhc.mutation.SetCreatedAt(t)
	return hhc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableCreatedAt(t *time.Time) *HushHistoryCreate {
	if t != nil {
		hhc.SetCreatedAt(*t)
	}
	return hhc
}

// SetUpdatedAt sets the "updated_at" field.
func (hhc *HushHistoryCreate) SetUpdatedAt(t time.Time) *HushHistoryCreate {
	hhc.mutation.SetUpdatedAt(t)
	return hhc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableUpdatedAt(t *time.Time) *HushHistoryCreate {
	if t != nil {
		hhc.SetUpdatedAt(*t)
	}
	return hhc
}

// SetCreatedBy sets the "created_by" field.
func (hhc *HushHistoryCreate) SetCreatedBy(s string) *HushHistoryCreate {
	hhc.mutation.SetCreatedBy(s)
	return hhc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableCreatedBy(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetCreatedBy(*s)
	}
	return hhc
}

// SetUpdatedBy sets the "updated_by" field.
func (hhc *HushHistoryCreate) SetUpdatedBy(s string) *HushHistoryCreate {
	hhc.mutation.SetUpdatedBy(s)
	return hhc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableUpdatedBy(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetUpdatedBy(*s)
	}
	return hhc
}

// SetMappingID sets the "mapping_id" field.
func (hhc *HushHistoryCreate) SetMappingID(s string) *HushHistoryCreate {
	hhc.mutation.SetMappingID(s)
	return hhc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableMappingID(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetMappingID(*s)
	}
	return hhc
}

// SetDeletedAt sets the "deleted_at" field.
func (hhc *HushHistoryCreate) SetDeletedAt(t time.Time) *HushHistoryCreate {
	hhc.mutation.SetDeletedAt(t)
	return hhc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableDeletedAt(t *time.Time) *HushHistoryCreate {
	if t != nil {
		hhc.SetDeletedAt(*t)
	}
	return hhc
}

// SetDeletedBy sets the "deleted_by" field.
func (hhc *HushHistoryCreate) SetDeletedBy(s string) *HushHistoryCreate {
	hhc.mutation.SetDeletedBy(s)
	return hhc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableDeletedBy(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetDeletedBy(*s)
	}
	return hhc
}

// SetName sets the "name" field.
func (hhc *HushHistoryCreate) SetName(s string) *HushHistoryCreate {
	hhc.mutation.SetName(s)
	return hhc
}

// SetDescription sets the "description" field.
func (hhc *HushHistoryCreate) SetDescription(s string) *HushHistoryCreate {
	hhc.mutation.SetDescription(s)
	return hhc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableDescription(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetDescription(*s)
	}
	return hhc
}

// SetKind sets the "kind" field.
func (hhc *HushHistoryCreate) SetKind(s string) *HushHistoryCreate {
	hhc.mutation.SetKind(s)
	return hhc
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableKind(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetKind(*s)
	}
	return hhc
}

// SetSecretName sets the "secret_name" field.
func (hhc *HushHistoryCreate) SetSecretName(s string) *HushHistoryCreate {
	hhc.mutation.SetSecretName(s)
	return hhc
}

// SetNillableSecretName sets the "secret_name" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableSecretName(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetSecretName(*s)
	}
	return hhc
}

// SetSecretValue sets the "secret_value" field.
func (hhc *HushHistoryCreate) SetSecretValue(s string) *HushHistoryCreate {
	hhc.mutation.SetSecretValue(s)
	return hhc
}

// SetNillableSecretValue sets the "secret_value" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableSecretValue(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetSecretValue(*s)
	}
	return hhc
}

// SetID sets the "id" field.
func (hhc *HushHistoryCreate) SetID(s string) *HushHistoryCreate {
	hhc.mutation.SetID(s)
	return hhc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (hhc *HushHistoryCreate) SetNillableID(s *string) *HushHistoryCreate {
	if s != nil {
		hhc.SetID(*s)
	}
	return hhc
}

// Mutation returns the HushHistoryMutation object of the builder.
func (hhc *HushHistoryCreate) Mutation() *HushHistoryMutation {
	return hhc.mutation
}

// Save creates the HushHistory in the database.
func (hhc *HushHistoryCreate) Save(ctx context.Context) (*HushHistory, error) {
	hhc.defaults()
	return withHooks(ctx, hhc.sqlSave, hhc.mutation, hhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hhc *HushHistoryCreate) SaveX(ctx context.Context) *HushHistory {
	v, err := hhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hhc *HushHistoryCreate) Exec(ctx context.Context) error {
	_, err := hhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hhc *HushHistoryCreate) ExecX(ctx context.Context) {
	if err := hhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hhc *HushHistoryCreate) defaults() {
	if _, ok := hhc.mutation.HistoryTime(); !ok {
		v := hushhistory.DefaultHistoryTime()
		hhc.mutation.SetHistoryTime(v)
	}
	if _, ok := hhc.mutation.CreatedAt(); !ok {
		v := hushhistory.DefaultCreatedAt()
		hhc.mutation.SetCreatedAt(v)
	}
	if _, ok := hhc.mutation.UpdatedAt(); !ok {
		v := hushhistory.DefaultUpdatedAt()
		hhc.mutation.SetUpdatedAt(v)
	}
	if _, ok := hhc.mutation.MappingID(); !ok {
		v := hushhistory.DefaultMappingID()
		hhc.mutation.SetMappingID(v)
	}
	if _, ok := hhc.mutation.ID(); !ok {
		v := hushhistory.DefaultID()
		hhc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hhc *HushHistoryCreate) check() error {
	if _, ok := hhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "HushHistory.history_time"`)}
	}
	if _, ok := hhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "HushHistory.operation"`)}
	}
	if v, ok := hhc.mutation.Operation(); ok {
		if err := hushhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "HushHistory.operation": %w`, err)}
		}
	}
	if _, ok := hhc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "HushHistory.mapping_id"`)}
	}
	if _, ok := hhc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "HushHistory.name"`)}
	}
	return nil
}

func (hhc *HushHistoryCreate) sqlSave(ctx context.Context) (*HushHistory, error) {
	if err := hhc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hhc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected HushHistory.ID type: %T", _spec.ID.Value)
		}
	}
	hhc.mutation.id = &_node.ID
	hhc.mutation.done = true
	return _node, nil
}

func (hhc *HushHistoryCreate) createSpec() (*HushHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &HushHistory{config: hhc.config}
		_spec = sqlgraph.NewCreateSpec(hushhistory.Table, sqlgraph.NewFieldSpec(hushhistory.FieldID, field.TypeString))
	)
	_spec.Schema = hhc.schemaConfig.HushHistory
	if id, ok := hhc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hhc.mutation.HistoryTime(); ok {
		_spec.SetField(hushhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := hhc.mutation.Operation(); ok {
		_spec.SetField(hushhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := hhc.mutation.Ref(); ok {
		_spec.SetField(hushhistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := hhc.mutation.CreatedAt(); ok {
		_spec.SetField(hushhistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hhc.mutation.UpdatedAt(); ok {
		_spec.SetField(hushhistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := hhc.mutation.CreatedBy(); ok {
		_spec.SetField(hushhistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := hhc.mutation.UpdatedBy(); ok {
		_spec.SetField(hushhistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := hhc.mutation.MappingID(); ok {
		_spec.SetField(hushhistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := hhc.mutation.DeletedAt(); ok {
		_spec.SetField(hushhistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := hhc.mutation.DeletedBy(); ok {
		_spec.SetField(hushhistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := hhc.mutation.Name(); ok {
		_spec.SetField(hushhistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := hhc.mutation.Description(); ok {
		_spec.SetField(hushhistory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := hhc.mutation.Kind(); ok {
		_spec.SetField(hushhistory.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if value, ok := hhc.mutation.SecretName(); ok {
		_spec.SetField(hushhistory.FieldSecretName, field.TypeString, value)
		_node.SecretName = value
	}
	if value, ok := hhc.mutation.SecretValue(); ok {
		_spec.SetField(hushhistory.FieldSecretValue, field.TypeString, value)
		_node.SecretValue = value
	}
	return _node, _spec
}

// HushHistoryCreateBulk is the builder for creating many HushHistory entities in bulk.
type HushHistoryCreateBulk struct {
	config
	err      error
	builders []*HushHistoryCreate
}

// Save creates the HushHistory entities in the database.
func (hhcb *HushHistoryCreateBulk) Save(ctx context.Context) ([]*HushHistory, error) {
	if hhcb.err != nil {
		return nil, hhcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hhcb.builders))
	nodes := make([]*HushHistory, len(hhcb.builders))
	mutators := make([]Mutator, len(hhcb.builders))
	for i := range hhcb.builders {
		func(i int, root context.Context) {
			builder := hhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HushHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, hhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hhcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, hhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hhcb *HushHistoryCreateBulk) SaveX(ctx context.Context) []*HushHistory {
	v, err := hhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hhcb *HushHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := hhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hhcb *HushHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := hhcb.Exec(ctx); err != nil {
		panic(err)
	}
}
