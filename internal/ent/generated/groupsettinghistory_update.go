// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/groupsettinghistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// GroupSettingHistoryUpdate is the builder for updating GroupSettingHistory entities.
type GroupSettingHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *GroupSettingHistoryMutation
}

// Where appends a list predicates to the GroupSettingHistoryUpdate builder.
func (gshu *GroupSettingHistoryUpdate) Where(ps ...predicate.GroupSettingHistory) *GroupSettingHistoryUpdate {
	gshu.mutation.Where(ps...)
	return gshu
}

// SetUpdatedAt sets the "updated_at" field.
func (gshu *GroupSettingHistoryUpdate) SetUpdatedAt(t time.Time) *GroupSettingHistoryUpdate {
	gshu.mutation.SetUpdatedAt(t)
	return gshu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gshu *GroupSettingHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *GroupSettingHistoryUpdate {
	if t != nil {
		gshu.SetUpdatedAt(*t)
	}
	return gshu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gshu *GroupSettingHistoryUpdate) ClearUpdatedAt() *GroupSettingHistoryUpdate {
	gshu.mutation.ClearUpdatedAt()
	return gshu
}

// SetUpdatedBy sets the "updated_by" field.
func (gshu *GroupSettingHistoryUpdate) SetUpdatedBy(s string) *GroupSettingHistoryUpdate {
	gshu.mutation.SetUpdatedBy(s)
	return gshu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gshu *GroupSettingHistoryUpdate) SetNillableUpdatedBy(s *string) *GroupSettingHistoryUpdate {
	if s != nil {
		gshu.SetUpdatedBy(*s)
	}
	return gshu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (gshu *GroupSettingHistoryUpdate) ClearUpdatedBy() *GroupSettingHistoryUpdate {
	gshu.mutation.ClearUpdatedBy()
	return gshu
}

// SetDeletedAt sets the "deleted_at" field.
func (gshu *GroupSettingHistoryUpdate) SetDeletedAt(t time.Time) *GroupSettingHistoryUpdate {
	gshu.mutation.SetDeletedAt(t)
	return gshu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gshu *GroupSettingHistoryUpdate) SetNillableDeletedAt(t *time.Time) *GroupSettingHistoryUpdate {
	if t != nil {
		gshu.SetDeletedAt(*t)
	}
	return gshu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gshu *GroupSettingHistoryUpdate) ClearDeletedAt() *GroupSettingHistoryUpdate {
	gshu.mutation.ClearDeletedAt()
	return gshu
}

// SetDeletedBy sets the "deleted_by" field.
func (gshu *GroupSettingHistoryUpdate) SetDeletedBy(s string) *GroupSettingHistoryUpdate {
	gshu.mutation.SetDeletedBy(s)
	return gshu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gshu *GroupSettingHistoryUpdate) SetNillableDeletedBy(s *string) *GroupSettingHistoryUpdate {
	if s != nil {
		gshu.SetDeletedBy(*s)
	}
	return gshu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (gshu *GroupSettingHistoryUpdate) ClearDeletedBy() *GroupSettingHistoryUpdate {
	gshu.mutation.ClearDeletedBy()
	return gshu
}

// SetVisibility sets the "visibility" field.
func (gshu *GroupSettingHistoryUpdate) SetVisibility(e enums.Visibility) *GroupSettingHistoryUpdate {
	gshu.mutation.SetVisibility(e)
	return gshu
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (gshu *GroupSettingHistoryUpdate) SetNillableVisibility(e *enums.Visibility) *GroupSettingHistoryUpdate {
	if e != nil {
		gshu.SetVisibility(*e)
	}
	return gshu
}

// SetJoinPolicy sets the "join_policy" field.
func (gshu *GroupSettingHistoryUpdate) SetJoinPolicy(ep enums.JoinPolicy) *GroupSettingHistoryUpdate {
	gshu.mutation.SetJoinPolicy(ep)
	return gshu
}

// SetNillableJoinPolicy sets the "join_policy" field if the given value is not nil.
func (gshu *GroupSettingHistoryUpdate) SetNillableJoinPolicy(ep *enums.JoinPolicy) *GroupSettingHistoryUpdate {
	if ep != nil {
		gshu.SetJoinPolicy(*ep)
	}
	return gshu
}

// SetTags sets the "tags" field.
func (gshu *GroupSettingHistoryUpdate) SetTags(s []string) *GroupSettingHistoryUpdate {
	gshu.mutation.SetTags(s)
	return gshu
}

// AppendTags appends s to the "tags" field.
func (gshu *GroupSettingHistoryUpdate) AppendTags(s []string) *GroupSettingHistoryUpdate {
	gshu.mutation.AppendTags(s)
	return gshu
}

// ClearTags clears the value of the "tags" field.
func (gshu *GroupSettingHistoryUpdate) ClearTags() *GroupSettingHistoryUpdate {
	gshu.mutation.ClearTags()
	return gshu
}

// SetSyncToSlack sets the "sync_to_slack" field.
func (gshu *GroupSettingHistoryUpdate) SetSyncToSlack(b bool) *GroupSettingHistoryUpdate {
	gshu.mutation.SetSyncToSlack(b)
	return gshu
}

// SetNillableSyncToSlack sets the "sync_to_slack" field if the given value is not nil.
func (gshu *GroupSettingHistoryUpdate) SetNillableSyncToSlack(b *bool) *GroupSettingHistoryUpdate {
	if b != nil {
		gshu.SetSyncToSlack(*b)
	}
	return gshu
}

// ClearSyncToSlack clears the value of the "sync_to_slack" field.
func (gshu *GroupSettingHistoryUpdate) ClearSyncToSlack() *GroupSettingHistoryUpdate {
	gshu.mutation.ClearSyncToSlack()
	return gshu
}

// SetSyncToGithub sets the "sync_to_github" field.
func (gshu *GroupSettingHistoryUpdate) SetSyncToGithub(b bool) *GroupSettingHistoryUpdate {
	gshu.mutation.SetSyncToGithub(b)
	return gshu
}

// SetNillableSyncToGithub sets the "sync_to_github" field if the given value is not nil.
func (gshu *GroupSettingHistoryUpdate) SetNillableSyncToGithub(b *bool) *GroupSettingHistoryUpdate {
	if b != nil {
		gshu.SetSyncToGithub(*b)
	}
	return gshu
}

// ClearSyncToGithub clears the value of the "sync_to_github" field.
func (gshu *GroupSettingHistoryUpdate) ClearSyncToGithub() *GroupSettingHistoryUpdate {
	gshu.mutation.ClearSyncToGithub()
	return gshu
}

// SetGroupID sets the "group_id" field.
func (gshu *GroupSettingHistoryUpdate) SetGroupID(s string) *GroupSettingHistoryUpdate {
	gshu.mutation.SetGroupID(s)
	return gshu
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (gshu *GroupSettingHistoryUpdate) SetNillableGroupID(s *string) *GroupSettingHistoryUpdate {
	if s != nil {
		gshu.SetGroupID(*s)
	}
	return gshu
}

// ClearGroupID clears the value of the "group_id" field.
func (gshu *GroupSettingHistoryUpdate) ClearGroupID() *GroupSettingHistoryUpdate {
	gshu.mutation.ClearGroupID()
	return gshu
}

// Mutation returns the GroupSettingHistoryMutation object of the builder.
func (gshu *GroupSettingHistoryUpdate) Mutation() *GroupSettingHistoryMutation {
	return gshu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gshu *GroupSettingHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gshu.sqlSave, gshu.mutation, gshu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gshu *GroupSettingHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := gshu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gshu *GroupSettingHistoryUpdate) Exec(ctx context.Context) error {
	_, err := gshu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gshu *GroupSettingHistoryUpdate) ExecX(ctx context.Context) {
	if err := gshu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gshu *GroupSettingHistoryUpdate) check() error {
	if v, ok := gshu.mutation.Visibility(); ok {
		if err := groupsettinghistory.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "GroupSettingHistory.visibility": %w`, err)}
		}
	}
	if v, ok := gshu.mutation.JoinPolicy(); ok {
		if err := groupsettinghistory.JoinPolicyValidator(v); err != nil {
			return &ValidationError{Name: "join_policy", err: fmt.Errorf(`generated: validator failed for field "GroupSettingHistory.join_policy": %w`, err)}
		}
	}
	return nil
}

func (gshu *GroupSettingHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gshu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupsettinghistory.Table, groupsettinghistory.Columns, sqlgraph.NewFieldSpec(groupsettinghistory.FieldID, field.TypeString))
	if ps := gshu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gshu.mutation.RefCleared() {
		_spec.ClearField(groupsettinghistory.FieldRef, field.TypeString)
	}
	if gshu.mutation.CreatedAtCleared() {
		_spec.ClearField(groupsettinghistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := gshu.mutation.UpdatedAt(); ok {
		_spec.SetField(groupsettinghistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if gshu.mutation.UpdatedAtCleared() {
		_spec.ClearField(groupsettinghistory.FieldUpdatedAt, field.TypeTime)
	}
	if gshu.mutation.CreatedByCleared() {
		_spec.ClearField(groupsettinghistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := gshu.mutation.UpdatedBy(); ok {
		_spec.SetField(groupsettinghistory.FieldUpdatedBy, field.TypeString, value)
	}
	if gshu.mutation.UpdatedByCleared() {
		_spec.ClearField(groupsettinghistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := gshu.mutation.DeletedAt(); ok {
		_spec.SetField(groupsettinghistory.FieldDeletedAt, field.TypeTime, value)
	}
	if gshu.mutation.DeletedAtCleared() {
		_spec.ClearField(groupsettinghistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gshu.mutation.DeletedBy(); ok {
		_spec.SetField(groupsettinghistory.FieldDeletedBy, field.TypeString, value)
	}
	if gshu.mutation.DeletedByCleared() {
		_spec.ClearField(groupsettinghistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := gshu.mutation.Visibility(); ok {
		_spec.SetField(groupsettinghistory.FieldVisibility, field.TypeEnum, value)
	}
	if value, ok := gshu.mutation.JoinPolicy(); ok {
		_spec.SetField(groupsettinghistory.FieldJoinPolicy, field.TypeEnum, value)
	}
	if value, ok := gshu.mutation.Tags(); ok {
		_spec.SetField(groupsettinghistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := gshu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, groupsettinghistory.FieldTags, value)
		})
	}
	if gshu.mutation.TagsCleared() {
		_spec.ClearField(groupsettinghistory.FieldTags, field.TypeJSON)
	}
	if value, ok := gshu.mutation.SyncToSlack(); ok {
		_spec.SetField(groupsettinghistory.FieldSyncToSlack, field.TypeBool, value)
	}
	if gshu.mutation.SyncToSlackCleared() {
		_spec.ClearField(groupsettinghistory.FieldSyncToSlack, field.TypeBool)
	}
	if value, ok := gshu.mutation.SyncToGithub(); ok {
		_spec.SetField(groupsettinghistory.FieldSyncToGithub, field.TypeBool, value)
	}
	if gshu.mutation.SyncToGithubCleared() {
		_spec.ClearField(groupsettinghistory.FieldSyncToGithub, field.TypeBool)
	}
	if value, ok := gshu.mutation.GroupID(); ok {
		_spec.SetField(groupsettinghistory.FieldGroupID, field.TypeString, value)
	}
	if gshu.mutation.GroupIDCleared() {
		_spec.ClearField(groupsettinghistory.FieldGroupID, field.TypeString)
	}
	_spec.Node.Schema = gshu.schemaConfig.GroupSettingHistory
	ctx = internal.NewSchemaConfigContext(ctx, gshu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, gshu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupsettinghistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gshu.mutation.done = true
	return n, nil
}

// GroupSettingHistoryUpdateOne is the builder for updating a single GroupSettingHistory entity.
type GroupSettingHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupSettingHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetUpdatedAt(t time.Time) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetUpdatedAt(t)
	return gshuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gshuo *GroupSettingHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *GroupSettingHistoryUpdateOne {
	if t != nil {
		gshuo.SetUpdatedAt(*t)
	}
	return gshuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gshuo *GroupSettingHistoryUpdateOne) ClearUpdatedAt() *GroupSettingHistoryUpdateOne {
	gshuo.mutation.ClearUpdatedAt()
	return gshuo
}

// SetUpdatedBy sets the "updated_by" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetUpdatedBy(s string) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetUpdatedBy(s)
	return gshuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gshuo *GroupSettingHistoryUpdateOne) SetNillableUpdatedBy(s *string) *GroupSettingHistoryUpdateOne {
	if s != nil {
		gshuo.SetUpdatedBy(*s)
	}
	return gshuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (gshuo *GroupSettingHistoryUpdateOne) ClearUpdatedBy() *GroupSettingHistoryUpdateOne {
	gshuo.mutation.ClearUpdatedBy()
	return gshuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetDeletedAt(t time.Time) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetDeletedAt(t)
	return gshuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gshuo *GroupSettingHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupSettingHistoryUpdateOne {
	if t != nil {
		gshuo.SetDeletedAt(*t)
	}
	return gshuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gshuo *GroupSettingHistoryUpdateOne) ClearDeletedAt() *GroupSettingHistoryUpdateOne {
	gshuo.mutation.ClearDeletedAt()
	return gshuo
}

// SetDeletedBy sets the "deleted_by" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetDeletedBy(s string) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetDeletedBy(s)
	return gshuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gshuo *GroupSettingHistoryUpdateOne) SetNillableDeletedBy(s *string) *GroupSettingHistoryUpdateOne {
	if s != nil {
		gshuo.SetDeletedBy(*s)
	}
	return gshuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (gshuo *GroupSettingHistoryUpdateOne) ClearDeletedBy() *GroupSettingHistoryUpdateOne {
	gshuo.mutation.ClearDeletedBy()
	return gshuo
}

// SetVisibility sets the "visibility" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetVisibility(e enums.Visibility) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetVisibility(e)
	return gshuo
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (gshuo *GroupSettingHistoryUpdateOne) SetNillableVisibility(e *enums.Visibility) *GroupSettingHistoryUpdateOne {
	if e != nil {
		gshuo.SetVisibility(*e)
	}
	return gshuo
}

// SetJoinPolicy sets the "join_policy" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetJoinPolicy(ep enums.JoinPolicy) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetJoinPolicy(ep)
	return gshuo
}

// SetNillableJoinPolicy sets the "join_policy" field if the given value is not nil.
func (gshuo *GroupSettingHistoryUpdateOne) SetNillableJoinPolicy(ep *enums.JoinPolicy) *GroupSettingHistoryUpdateOne {
	if ep != nil {
		gshuo.SetJoinPolicy(*ep)
	}
	return gshuo
}

// SetTags sets the "tags" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetTags(s []string) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetTags(s)
	return gshuo
}

// AppendTags appends s to the "tags" field.
func (gshuo *GroupSettingHistoryUpdateOne) AppendTags(s []string) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.AppendTags(s)
	return gshuo
}

// ClearTags clears the value of the "tags" field.
func (gshuo *GroupSettingHistoryUpdateOne) ClearTags() *GroupSettingHistoryUpdateOne {
	gshuo.mutation.ClearTags()
	return gshuo
}

// SetSyncToSlack sets the "sync_to_slack" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetSyncToSlack(b bool) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetSyncToSlack(b)
	return gshuo
}

// SetNillableSyncToSlack sets the "sync_to_slack" field if the given value is not nil.
func (gshuo *GroupSettingHistoryUpdateOne) SetNillableSyncToSlack(b *bool) *GroupSettingHistoryUpdateOne {
	if b != nil {
		gshuo.SetSyncToSlack(*b)
	}
	return gshuo
}

// ClearSyncToSlack clears the value of the "sync_to_slack" field.
func (gshuo *GroupSettingHistoryUpdateOne) ClearSyncToSlack() *GroupSettingHistoryUpdateOne {
	gshuo.mutation.ClearSyncToSlack()
	return gshuo
}

// SetSyncToGithub sets the "sync_to_github" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetSyncToGithub(b bool) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetSyncToGithub(b)
	return gshuo
}

// SetNillableSyncToGithub sets the "sync_to_github" field if the given value is not nil.
func (gshuo *GroupSettingHistoryUpdateOne) SetNillableSyncToGithub(b *bool) *GroupSettingHistoryUpdateOne {
	if b != nil {
		gshuo.SetSyncToGithub(*b)
	}
	return gshuo
}

// ClearSyncToGithub clears the value of the "sync_to_github" field.
func (gshuo *GroupSettingHistoryUpdateOne) ClearSyncToGithub() *GroupSettingHistoryUpdateOne {
	gshuo.mutation.ClearSyncToGithub()
	return gshuo
}

// SetGroupID sets the "group_id" field.
func (gshuo *GroupSettingHistoryUpdateOne) SetGroupID(s string) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.SetGroupID(s)
	return gshuo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (gshuo *GroupSettingHistoryUpdateOne) SetNillableGroupID(s *string) *GroupSettingHistoryUpdateOne {
	if s != nil {
		gshuo.SetGroupID(*s)
	}
	return gshuo
}

// ClearGroupID clears the value of the "group_id" field.
func (gshuo *GroupSettingHistoryUpdateOne) ClearGroupID() *GroupSettingHistoryUpdateOne {
	gshuo.mutation.ClearGroupID()
	return gshuo
}

// Mutation returns the GroupSettingHistoryMutation object of the builder.
func (gshuo *GroupSettingHistoryUpdateOne) Mutation() *GroupSettingHistoryMutation {
	return gshuo.mutation
}

// Where appends a list predicates to the GroupSettingHistoryUpdate builder.
func (gshuo *GroupSettingHistoryUpdateOne) Where(ps ...predicate.GroupSettingHistory) *GroupSettingHistoryUpdateOne {
	gshuo.mutation.Where(ps...)
	return gshuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gshuo *GroupSettingHistoryUpdateOne) Select(field string, fields ...string) *GroupSettingHistoryUpdateOne {
	gshuo.fields = append([]string{field}, fields...)
	return gshuo
}

// Save executes the query and returns the updated GroupSettingHistory entity.
func (gshuo *GroupSettingHistoryUpdateOne) Save(ctx context.Context) (*GroupSettingHistory, error) {
	return withHooks(ctx, gshuo.sqlSave, gshuo.mutation, gshuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gshuo *GroupSettingHistoryUpdateOne) SaveX(ctx context.Context) *GroupSettingHistory {
	node, err := gshuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gshuo *GroupSettingHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := gshuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gshuo *GroupSettingHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := gshuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gshuo *GroupSettingHistoryUpdateOne) check() error {
	if v, ok := gshuo.mutation.Visibility(); ok {
		if err := groupsettinghistory.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "GroupSettingHistory.visibility": %w`, err)}
		}
	}
	if v, ok := gshuo.mutation.JoinPolicy(); ok {
		if err := groupsettinghistory.JoinPolicyValidator(v); err != nil {
			return &ValidationError{Name: "join_policy", err: fmt.Errorf(`generated: validator failed for field "GroupSettingHistory.join_policy": %w`, err)}
		}
	}
	return nil
}

func (gshuo *GroupSettingHistoryUpdateOne) sqlSave(ctx context.Context) (_node *GroupSettingHistory, err error) {
	if err := gshuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupsettinghistory.Table, groupsettinghistory.Columns, sqlgraph.NewFieldSpec(groupsettinghistory.FieldID, field.TypeString))
	id, ok := gshuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "GroupSettingHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gshuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupsettinghistory.FieldID)
		for _, f := range fields {
			if !groupsettinghistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != groupsettinghistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gshuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gshuo.mutation.RefCleared() {
		_spec.ClearField(groupsettinghistory.FieldRef, field.TypeString)
	}
	if gshuo.mutation.CreatedAtCleared() {
		_spec.ClearField(groupsettinghistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := gshuo.mutation.UpdatedAt(); ok {
		_spec.SetField(groupsettinghistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if gshuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(groupsettinghistory.FieldUpdatedAt, field.TypeTime)
	}
	if gshuo.mutation.CreatedByCleared() {
		_spec.ClearField(groupsettinghistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := gshuo.mutation.UpdatedBy(); ok {
		_spec.SetField(groupsettinghistory.FieldUpdatedBy, field.TypeString, value)
	}
	if gshuo.mutation.UpdatedByCleared() {
		_spec.ClearField(groupsettinghistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := gshuo.mutation.DeletedAt(); ok {
		_spec.SetField(groupsettinghistory.FieldDeletedAt, field.TypeTime, value)
	}
	if gshuo.mutation.DeletedAtCleared() {
		_spec.ClearField(groupsettinghistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gshuo.mutation.DeletedBy(); ok {
		_spec.SetField(groupsettinghistory.FieldDeletedBy, field.TypeString, value)
	}
	if gshuo.mutation.DeletedByCleared() {
		_spec.ClearField(groupsettinghistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := gshuo.mutation.Visibility(); ok {
		_spec.SetField(groupsettinghistory.FieldVisibility, field.TypeEnum, value)
	}
	if value, ok := gshuo.mutation.JoinPolicy(); ok {
		_spec.SetField(groupsettinghistory.FieldJoinPolicy, field.TypeEnum, value)
	}
	if value, ok := gshuo.mutation.Tags(); ok {
		_spec.SetField(groupsettinghistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := gshuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, groupsettinghistory.FieldTags, value)
		})
	}
	if gshuo.mutation.TagsCleared() {
		_spec.ClearField(groupsettinghistory.FieldTags, field.TypeJSON)
	}
	if value, ok := gshuo.mutation.SyncToSlack(); ok {
		_spec.SetField(groupsettinghistory.FieldSyncToSlack, field.TypeBool, value)
	}
	if gshuo.mutation.SyncToSlackCleared() {
		_spec.ClearField(groupsettinghistory.FieldSyncToSlack, field.TypeBool)
	}
	if value, ok := gshuo.mutation.SyncToGithub(); ok {
		_spec.SetField(groupsettinghistory.FieldSyncToGithub, field.TypeBool, value)
	}
	if gshuo.mutation.SyncToGithubCleared() {
		_spec.ClearField(groupsettinghistory.FieldSyncToGithub, field.TypeBool)
	}
	if value, ok := gshuo.mutation.GroupID(); ok {
		_spec.SetField(groupsettinghistory.FieldGroupID, field.TypeString, value)
	}
	if gshuo.mutation.GroupIDCleared() {
		_spec.ClearField(groupsettinghistory.FieldGroupID, field.TypeString)
	}
	_spec.Node.Schema = gshuo.schemaConfig.GroupSettingHistory
	ctx = internal.NewSchemaConfigContext(ctx, gshuo.schemaConfig)
	_node = &GroupSettingHistory{config: gshuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gshuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupsettinghistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gshuo.mutation.done = true
	return _node, nil
}