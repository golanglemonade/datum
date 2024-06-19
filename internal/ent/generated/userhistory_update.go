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
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/userhistory"
	"github.com/datumforge/datum/pkg/enums"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// UserHistoryUpdate is the builder for updating UserHistory entities.
type UserHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *UserHistoryMutation
}

// Where appends a list predicates to the UserHistoryUpdate builder.
func (uhu *UserHistoryUpdate) Where(ps ...predicate.UserHistory) *UserHistoryUpdate {
	uhu.mutation.Where(ps...)
	return uhu
}

// SetUpdatedAt sets the "updated_at" field.
func (uhu *UserHistoryUpdate) SetUpdatedAt(t time.Time) *UserHistoryUpdate {
	uhu.mutation.SetUpdatedAt(t)
	return uhu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *UserHistoryUpdate {
	if t != nil {
		uhu.SetUpdatedAt(*t)
	}
	return uhu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (uhu *UserHistoryUpdate) ClearUpdatedAt() *UserHistoryUpdate {
	uhu.mutation.ClearUpdatedAt()
	return uhu
}

// SetUpdatedBy sets the "updated_by" field.
func (uhu *UserHistoryUpdate) SetUpdatedBy(s string) *UserHistoryUpdate {
	uhu.mutation.SetUpdatedBy(s)
	return uhu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableUpdatedBy(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetUpdatedBy(*s)
	}
	return uhu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uhu *UserHistoryUpdate) ClearUpdatedBy() *UserHistoryUpdate {
	uhu.mutation.ClearUpdatedBy()
	return uhu
}

// SetDeletedAt sets the "deleted_at" field.
func (uhu *UserHistoryUpdate) SetDeletedAt(t time.Time) *UserHistoryUpdate {
	uhu.mutation.SetDeletedAt(t)
	return uhu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableDeletedAt(t *time.Time) *UserHistoryUpdate {
	if t != nil {
		uhu.SetDeletedAt(*t)
	}
	return uhu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uhu *UserHistoryUpdate) ClearDeletedAt() *UserHistoryUpdate {
	uhu.mutation.ClearDeletedAt()
	return uhu
}

// SetDeletedBy sets the "deleted_by" field.
func (uhu *UserHistoryUpdate) SetDeletedBy(s string) *UserHistoryUpdate {
	uhu.mutation.SetDeletedBy(s)
	return uhu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableDeletedBy(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetDeletedBy(*s)
	}
	return uhu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (uhu *UserHistoryUpdate) ClearDeletedBy() *UserHistoryUpdate {
	uhu.mutation.ClearDeletedBy()
	return uhu
}

// SetTags sets the "tags" field.
func (uhu *UserHistoryUpdate) SetTags(s []string) *UserHistoryUpdate {
	uhu.mutation.SetTags(s)
	return uhu
}

// AppendTags appends s to the "tags" field.
func (uhu *UserHistoryUpdate) AppendTags(s []string) *UserHistoryUpdate {
	uhu.mutation.AppendTags(s)
	return uhu
}

// ClearTags clears the value of the "tags" field.
func (uhu *UserHistoryUpdate) ClearTags() *UserHistoryUpdate {
	uhu.mutation.ClearTags()
	return uhu
}

// SetEmail sets the "email" field.
func (uhu *UserHistoryUpdate) SetEmail(s string) *UserHistoryUpdate {
	uhu.mutation.SetEmail(s)
	return uhu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableEmail(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetEmail(*s)
	}
	return uhu
}

// SetFirstName sets the "first_name" field.
func (uhu *UserHistoryUpdate) SetFirstName(s string) *UserHistoryUpdate {
	uhu.mutation.SetFirstName(s)
	return uhu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableFirstName(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetFirstName(*s)
	}
	return uhu
}

// ClearFirstName clears the value of the "first_name" field.
func (uhu *UserHistoryUpdate) ClearFirstName() *UserHistoryUpdate {
	uhu.mutation.ClearFirstName()
	return uhu
}

// SetLastName sets the "last_name" field.
func (uhu *UserHistoryUpdate) SetLastName(s string) *UserHistoryUpdate {
	uhu.mutation.SetLastName(s)
	return uhu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableLastName(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetLastName(*s)
	}
	return uhu
}

// ClearLastName clears the value of the "last_name" field.
func (uhu *UserHistoryUpdate) ClearLastName() *UserHistoryUpdate {
	uhu.mutation.ClearLastName()
	return uhu
}

// SetDisplayName sets the "display_name" field.
func (uhu *UserHistoryUpdate) SetDisplayName(s string) *UserHistoryUpdate {
	uhu.mutation.SetDisplayName(s)
	return uhu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableDisplayName(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetDisplayName(*s)
	}
	return uhu
}

// SetAvatarRemoteURL sets the "avatar_remote_url" field.
func (uhu *UserHistoryUpdate) SetAvatarRemoteURL(s string) *UserHistoryUpdate {
	uhu.mutation.SetAvatarRemoteURL(s)
	return uhu
}

// SetNillableAvatarRemoteURL sets the "avatar_remote_url" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableAvatarRemoteURL(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetAvatarRemoteURL(*s)
	}
	return uhu
}

// ClearAvatarRemoteURL clears the value of the "avatar_remote_url" field.
func (uhu *UserHistoryUpdate) ClearAvatarRemoteURL() *UserHistoryUpdate {
	uhu.mutation.ClearAvatarRemoteURL()
	return uhu
}

// SetAvatarLocalFile sets the "avatar_local_file" field.
func (uhu *UserHistoryUpdate) SetAvatarLocalFile(s string) *UserHistoryUpdate {
	uhu.mutation.SetAvatarLocalFile(s)
	return uhu
}

// SetNillableAvatarLocalFile sets the "avatar_local_file" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableAvatarLocalFile(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetAvatarLocalFile(*s)
	}
	return uhu
}

// ClearAvatarLocalFile clears the value of the "avatar_local_file" field.
func (uhu *UserHistoryUpdate) ClearAvatarLocalFile() *UserHistoryUpdate {
	uhu.mutation.ClearAvatarLocalFile()
	return uhu
}

// SetAvatarUpdatedAt sets the "avatar_updated_at" field.
func (uhu *UserHistoryUpdate) SetAvatarUpdatedAt(t time.Time) *UserHistoryUpdate {
	uhu.mutation.SetAvatarUpdatedAt(t)
	return uhu
}

// SetNillableAvatarUpdatedAt sets the "avatar_updated_at" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableAvatarUpdatedAt(t *time.Time) *UserHistoryUpdate {
	if t != nil {
		uhu.SetAvatarUpdatedAt(*t)
	}
	return uhu
}

// ClearAvatarUpdatedAt clears the value of the "avatar_updated_at" field.
func (uhu *UserHistoryUpdate) ClearAvatarUpdatedAt() *UserHistoryUpdate {
	uhu.mutation.ClearAvatarUpdatedAt()
	return uhu
}

// SetLastSeen sets the "last_seen" field.
func (uhu *UserHistoryUpdate) SetLastSeen(t time.Time) *UserHistoryUpdate {
	uhu.mutation.SetLastSeen(t)
	return uhu
}

// SetNillableLastSeen sets the "last_seen" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableLastSeen(t *time.Time) *UserHistoryUpdate {
	if t != nil {
		uhu.SetLastSeen(*t)
	}
	return uhu
}

// ClearLastSeen clears the value of the "last_seen" field.
func (uhu *UserHistoryUpdate) ClearLastSeen() *UserHistoryUpdate {
	uhu.mutation.ClearLastSeen()
	return uhu
}

// SetPassword sets the "password" field.
func (uhu *UserHistoryUpdate) SetPassword(s string) *UserHistoryUpdate {
	uhu.mutation.SetPassword(s)
	return uhu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillablePassword(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetPassword(*s)
	}
	return uhu
}

// ClearPassword clears the value of the "password" field.
func (uhu *UserHistoryUpdate) ClearPassword() *UserHistoryUpdate {
	uhu.mutation.ClearPassword()
	return uhu
}

// SetSub sets the "sub" field.
func (uhu *UserHistoryUpdate) SetSub(s string) *UserHistoryUpdate {
	uhu.mutation.SetSub(s)
	return uhu
}

// SetNillableSub sets the "sub" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableSub(s *string) *UserHistoryUpdate {
	if s != nil {
		uhu.SetSub(*s)
	}
	return uhu
}

// ClearSub clears the value of the "sub" field.
func (uhu *UserHistoryUpdate) ClearSub() *UserHistoryUpdate {
	uhu.mutation.ClearSub()
	return uhu
}

// SetAuthProvider sets the "auth_provider" field.
func (uhu *UserHistoryUpdate) SetAuthProvider(ep enums.AuthProvider) *UserHistoryUpdate {
	uhu.mutation.SetAuthProvider(ep)
	return uhu
}

// SetNillableAuthProvider sets the "auth_provider" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableAuthProvider(ep *enums.AuthProvider) *UserHistoryUpdate {
	if ep != nil {
		uhu.SetAuthProvider(*ep)
	}
	return uhu
}

// SetRole sets the "role" field.
func (uhu *UserHistoryUpdate) SetRole(e enums.Role) *UserHistoryUpdate {
	uhu.mutation.SetRole(e)
	return uhu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uhu *UserHistoryUpdate) SetNillableRole(e *enums.Role) *UserHistoryUpdate {
	if e != nil {
		uhu.SetRole(*e)
	}
	return uhu
}

// ClearRole clears the value of the "role" field.
func (uhu *UserHistoryUpdate) ClearRole() *UserHistoryUpdate {
	uhu.mutation.ClearRole()
	return uhu
}

// Mutation returns the UserHistoryMutation object of the builder.
func (uhu *UserHistoryUpdate) Mutation() *UserHistoryMutation {
	return uhu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uhu *UserHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uhu.sqlSave, uhu.mutation, uhu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uhu *UserHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := uhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uhu *UserHistoryUpdate) Exec(ctx context.Context) error {
	_, err := uhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uhu *UserHistoryUpdate) ExecX(ctx context.Context) {
	if err := uhu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uhu *UserHistoryUpdate) check() error {
	if v, ok := uhu.mutation.AuthProvider(); ok {
		if err := userhistory.AuthProviderValidator(v); err != nil {
			return &ValidationError{Name: "auth_provider", err: fmt.Errorf(`generated: validator failed for field "UserHistory.auth_provider": %w`, err)}
		}
	}
	if v, ok := uhu.mutation.Role(); ok {
		if err := userhistory.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "UserHistory.role": %w`, err)}
		}
	}
	return nil
}

func (uhu *UserHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uhu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userhistory.Table, userhistory.Columns, sqlgraph.NewFieldSpec(userhistory.FieldID, field.TypeString))
	if ps := uhu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uhu.mutation.RefCleared() {
		_spec.ClearField(userhistory.FieldRef, field.TypeString)
	}
	if uhu.mutation.CreatedAtCleared() {
		_spec.ClearField(userhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := uhu.mutation.UpdatedAt(); ok {
		_spec.SetField(userhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if uhu.mutation.UpdatedAtCleared() {
		_spec.ClearField(userhistory.FieldUpdatedAt, field.TypeTime)
	}
	if uhu.mutation.CreatedByCleared() {
		_spec.ClearField(userhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := uhu.mutation.UpdatedBy(); ok {
		_spec.SetField(userhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if uhu.mutation.UpdatedByCleared() {
		_spec.ClearField(userhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := uhu.mutation.DeletedAt(); ok {
		_spec.SetField(userhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if uhu.mutation.DeletedAtCleared() {
		_spec.ClearField(userhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uhu.mutation.DeletedBy(); ok {
		_spec.SetField(userhistory.FieldDeletedBy, field.TypeString, value)
	}
	if uhu.mutation.DeletedByCleared() {
		_spec.ClearField(userhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := uhu.mutation.Tags(); ok {
		_spec.SetField(userhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := uhu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, userhistory.FieldTags, value)
		})
	}
	if uhu.mutation.TagsCleared() {
		_spec.ClearField(userhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := uhu.mutation.Email(); ok {
		_spec.SetField(userhistory.FieldEmail, field.TypeString, value)
	}
	if value, ok := uhu.mutation.FirstName(); ok {
		_spec.SetField(userhistory.FieldFirstName, field.TypeString, value)
	}
	if uhu.mutation.FirstNameCleared() {
		_spec.ClearField(userhistory.FieldFirstName, field.TypeString)
	}
	if value, ok := uhu.mutation.LastName(); ok {
		_spec.SetField(userhistory.FieldLastName, field.TypeString, value)
	}
	if uhu.mutation.LastNameCleared() {
		_spec.ClearField(userhistory.FieldLastName, field.TypeString)
	}
	if value, ok := uhu.mutation.DisplayName(); ok {
		_spec.SetField(userhistory.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uhu.mutation.AvatarRemoteURL(); ok {
		_spec.SetField(userhistory.FieldAvatarRemoteURL, field.TypeString, value)
	}
	if uhu.mutation.AvatarRemoteURLCleared() {
		_spec.ClearField(userhistory.FieldAvatarRemoteURL, field.TypeString)
	}
	if value, ok := uhu.mutation.AvatarLocalFile(); ok {
		_spec.SetField(userhistory.FieldAvatarLocalFile, field.TypeString, value)
	}
	if uhu.mutation.AvatarLocalFileCleared() {
		_spec.ClearField(userhistory.FieldAvatarLocalFile, field.TypeString)
	}
	if value, ok := uhu.mutation.AvatarUpdatedAt(); ok {
		_spec.SetField(userhistory.FieldAvatarUpdatedAt, field.TypeTime, value)
	}
	if uhu.mutation.AvatarUpdatedAtCleared() {
		_spec.ClearField(userhistory.FieldAvatarUpdatedAt, field.TypeTime)
	}
	if value, ok := uhu.mutation.LastSeen(); ok {
		_spec.SetField(userhistory.FieldLastSeen, field.TypeTime, value)
	}
	if uhu.mutation.LastSeenCleared() {
		_spec.ClearField(userhistory.FieldLastSeen, field.TypeTime)
	}
	if value, ok := uhu.mutation.Password(); ok {
		_spec.SetField(userhistory.FieldPassword, field.TypeString, value)
	}
	if uhu.mutation.PasswordCleared() {
		_spec.ClearField(userhistory.FieldPassword, field.TypeString)
	}
	if value, ok := uhu.mutation.Sub(); ok {
		_spec.SetField(userhistory.FieldSub, field.TypeString, value)
	}
	if uhu.mutation.SubCleared() {
		_spec.ClearField(userhistory.FieldSub, field.TypeString)
	}
	if value, ok := uhu.mutation.AuthProvider(); ok {
		_spec.SetField(userhistory.FieldAuthProvider, field.TypeEnum, value)
	}
	if value, ok := uhu.mutation.Role(); ok {
		_spec.SetField(userhistory.FieldRole, field.TypeEnum, value)
	}
	if uhu.mutation.RoleCleared() {
		_spec.ClearField(userhistory.FieldRole, field.TypeEnum)
	}
	_spec.Node.Schema = uhu.schemaConfig.UserHistory
	ctx = internal.NewSchemaConfigContext(ctx, uhu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, uhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uhu.mutation.done = true
	return n, nil
}

// UserHistoryUpdateOne is the builder for updating a single UserHistory entity.
type UserHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uhuo *UserHistoryUpdateOne) SetUpdatedAt(t time.Time) *UserHistoryUpdateOne {
	uhuo.mutation.SetUpdatedAt(t)
	return uhuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserHistoryUpdateOne {
	if t != nil {
		uhuo.SetUpdatedAt(*t)
	}
	return uhuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (uhuo *UserHistoryUpdateOne) ClearUpdatedAt() *UserHistoryUpdateOne {
	uhuo.mutation.ClearUpdatedAt()
	return uhuo
}

// SetUpdatedBy sets the "updated_by" field.
func (uhuo *UserHistoryUpdateOne) SetUpdatedBy(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetUpdatedBy(s)
	return uhuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableUpdatedBy(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetUpdatedBy(*s)
	}
	return uhuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uhuo *UserHistoryUpdateOne) ClearUpdatedBy() *UserHistoryUpdateOne {
	uhuo.mutation.ClearUpdatedBy()
	return uhuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uhuo *UserHistoryUpdateOne) SetDeletedAt(t time.Time) *UserHistoryUpdateOne {
	uhuo.mutation.SetDeletedAt(t)
	return uhuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *UserHistoryUpdateOne {
	if t != nil {
		uhuo.SetDeletedAt(*t)
	}
	return uhuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uhuo *UserHistoryUpdateOne) ClearDeletedAt() *UserHistoryUpdateOne {
	uhuo.mutation.ClearDeletedAt()
	return uhuo
}

// SetDeletedBy sets the "deleted_by" field.
func (uhuo *UserHistoryUpdateOne) SetDeletedBy(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetDeletedBy(s)
	return uhuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableDeletedBy(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetDeletedBy(*s)
	}
	return uhuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (uhuo *UserHistoryUpdateOne) ClearDeletedBy() *UserHistoryUpdateOne {
	uhuo.mutation.ClearDeletedBy()
	return uhuo
}

// SetTags sets the "tags" field.
func (uhuo *UserHistoryUpdateOne) SetTags(s []string) *UserHistoryUpdateOne {
	uhuo.mutation.SetTags(s)
	return uhuo
}

// AppendTags appends s to the "tags" field.
func (uhuo *UserHistoryUpdateOne) AppendTags(s []string) *UserHistoryUpdateOne {
	uhuo.mutation.AppendTags(s)
	return uhuo
}

// ClearTags clears the value of the "tags" field.
func (uhuo *UserHistoryUpdateOne) ClearTags() *UserHistoryUpdateOne {
	uhuo.mutation.ClearTags()
	return uhuo
}

// SetEmail sets the "email" field.
func (uhuo *UserHistoryUpdateOne) SetEmail(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetEmail(s)
	return uhuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableEmail(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetEmail(*s)
	}
	return uhuo
}

// SetFirstName sets the "first_name" field.
func (uhuo *UserHistoryUpdateOne) SetFirstName(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetFirstName(s)
	return uhuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableFirstName(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetFirstName(*s)
	}
	return uhuo
}

// ClearFirstName clears the value of the "first_name" field.
func (uhuo *UserHistoryUpdateOne) ClearFirstName() *UserHistoryUpdateOne {
	uhuo.mutation.ClearFirstName()
	return uhuo
}

// SetLastName sets the "last_name" field.
func (uhuo *UserHistoryUpdateOne) SetLastName(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetLastName(s)
	return uhuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableLastName(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetLastName(*s)
	}
	return uhuo
}

// ClearLastName clears the value of the "last_name" field.
func (uhuo *UserHistoryUpdateOne) ClearLastName() *UserHistoryUpdateOne {
	uhuo.mutation.ClearLastName()
	return uhuo
}

// SetDisplayName sets the "display_name" field.
func (uhuo *UserHistoryUpdateOne) SetDisplayName(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetDisplayName(s)
	return uhuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableDisplayName(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetDisplayName(*s)
	}
	return uhuo
}

// SetAvatarRemoteURL sets the "avatar_remote_url" field.
func (uhuo *UserHistoryUpdateOne) SetAvatarRemoteURL(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetAvatarRemoteURL(s)
	return uhuo
}

// SetNillableAvatarRemoteURL sets the "avatar_remote_url" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableAvatarRemoteURL(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetAvatarRemoteURL(*s)
	}
	return uhuo
}

// ClearAvatarRemoteURL clears the value of the "avatar_remote_url" field.
func (uhuo *UserHistoryUpdateOne) ClearAvatarRemoteURL() *UserHistoryUpdateOne {
	uhuo.mutation.ClearAvatarRemoteURL()
	return uhuo
}

// SetAvatarLocalFile sets the "avatar_local_file" field.
func (uhuo *UserHistoryUpdateOne) SetAvatarLocalFile(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetAvatarLocalFile(s)
	return uhuo
}

// SetNillableAvatarLocalFile sets the "avatar_local_file" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableAvatarLocalFile(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetAvatarLocalFile(*s)
	}
	return uhuo
}

// ClearAvatarLocalFile clears the value of the "avatar_local_file" field.
func (uhuo *UserHistoryUpdateOne) ClearAvatarLocalFile() *UserHistoryUpdateOne {
	uhuo.mutation.ClearAvatarLocalFile()
	return uhuo
}

// SetAvatarUpdatedAt sets the "avatar_updated_at" field.
func (uhuo *UserHistoryUpdateOne) SetAvatarUpdatedAt(t time.Time) *UserHistoryUpdateOne {
	uhuo.mutation.SetAvatarUpdatedAt(t)
	return uhuo
}

// SetNillableAvatarUpdatedAt sets the "avatar_updated_at" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableAvatarUpdatedAt(t *time.Time) *UserHistoryUpdateOne {
	if t != nil {
		uhuo.SetAvatarUpdatedAt(*t)
	}
	return uhuo
}

// ClearAvatarUpdatedAt clears the value of the "avatar_updated_at" field.
func (uhuo *UserHistoryUpdateOne) ClearAvatarUpdatedAt() *UserHistoryUpdateOne {
	uhuo.mutation.ClearAvatarUpdatedAt()
	return uhuo
}

// SetLastSeen sets the "last_seen" field.
func (uhuo *UserHistoryUpdateOne) SetLastSeen(t time.Time) *UserHistoryUpdateOne {
	uhuo.mutation.SetLastSeen(t)
	return uhuo
}

// SetNillableLastSeen sets the "last_seen" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableLastSeen(t *time.Time) *UserHistoryUpdateOne {
	if t != nil {
		uhuo.SetLastSeen(*t)
	}
	return uhuo
}

// ClearLastSeen clears the value of the "last_seen" field.
func (uhuo *UserHistoryUpdateOne) ClearLastSeen() *UserHistoryUpdateOne {
	uhuo.mutation.ClearLastSeen()
	return uhuo
}

// SetPassword sets the "password" field.
func (uhuo *UserHistoryUpdateOne) SetPassword(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetPassword(s)
	return uhuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillablePassword(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetPassword(*s)
	}
	return uhuo
}

// ClearPassword clears the value of the "password" field.
func (uhuo *UserHistoryUpdateOne) ClearPassword() *UserHistoryUpdateOne {
	uhuo.mutation.ClearPassword()
	return uhuo
}

// SetSub sets the "sub" field.
func (uhuo *UserHistoryUpdateOne) SetSub(s string) *UserHistoryUpdateOne {
	uhuo.mutation.SetSub(s)
	return uhuo
}

// SetNillableSub sets the "sub" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableSub(s *string) *UserHistoryUpdateOne {
	if s != nil {
		uhuo.SetSub(*s)
	}
	return uhuo
}

// ClearSub clears the value of the "sub" field.
func (uhuo *UserHistoryUpdateOne) ClearSub() *UserHistoryUpdateOne {
	uhuo.mutation.ClearSub()
	return uhuo
}

// SetAuthProvider sets the "auth_provider" field.
func (uhuo *UserHistoryUpdateOne) SetAuthProvider(ep enums.AuthProvider) *UserHistoryUpdateOne {
	uhuo.mutation.SetAuthProvider(ep)
	return uhuo
}

// SetNillableAuthProvider sets the "auth_provider" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableAuthProvider(ep *enums.AuthProvider) *UserHistoryUpdateOne {
	if ep != nil {
		uhuo.SetAuthProvider(*ep)
	}
	return uhuo
}

// SetRole sets the "role" field.
func (uhuo *UserHistoryUpdateOne) SetRole(e enums.Role) *UserHistoryUpdateOne {
	uhuo.mutation.SetRole(e)
	return uhuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uhuo *UserHistoryUpdateOne) SetNillableRole(e *enums.Role) *UserHistoryUpdateOne {
	if e != nil {
		uhuo.SetRole(*e)
	}
	return uhuo
}

// ClearRole clears the value of the "role" field.
func (uhuo *UserHistoryUpdateOne) ClearRole() *UserHistoryUpdateOne {
	uhuo.mutation.ClearRole()
	return uhuo
}

// Mutation returns the UserHistoryMutation object of the builder.
func (uhuo *UserHistoryUpdateOne) Mutation() *UserHistoryMutation {
	return uhuo.mutation
}

// Where appends a list predicates to the UserHistoryUpdate builder.
func (uhuo *UserHistoryUpdateOne) Where(ps ...predicate.UserHistory) *UserHistoryUpdateOne {
	uhuo.mutation.Where(ps...)
	return uhuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uhuo *UserHistoryUpdateOne) Select(field string, fields ...string) *UserHistoryUpdateOne {
	uhuo.fields = append([]string{field}, fields...)
	return uhuo
}

// Save executes the query and returns the updated UserHistory entity.
func (uhuo *UserHistoryUpdateOne) Save(ctx context.Context) (*UserHistory, error) {
	return withHooks(ctx, uhuo.sqlSave, uhuo.mutation, uhuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uhuo *UserHistoryUpdateOne) SaveX(ctx context.Context) *UserHistory {
	node, err := uhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uhuo *UserHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := uhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uhuo *UserHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := uhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uhuo *UserHistoryUpdateOne) check() error {
	if v, ok := uhuo.mutation.AuthProvider(); ok {
		if err := userhistory.AuthProviderValidator(v); err != nil {
			return &ValidationError{Name: "auth_provider", err: fmt.Errorf(`generated: validator failed for field "UserHistory.auth_provider": %w`, err)}
		}
	}
	if v, ok := uhuo.mutation.Role(); ok {
		if err := userhistory.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "UserHistory.role": %w`, err)}
		}
	}
	return nil
}

func (uhuo *UserHistoryUpdateOne) sqlSave(ctx context.Context) (_node *UserHistory, err error) {
	if err := uhuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userhistory.Table, userhistory.Columns, sqlgraph.NewFieldSpec(userhistory.FieldID, field.TypeString))
	id, ok := uhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "UserHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uhuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userhistory.FieldID)
		for _, f := range fields {
			if !userhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != userhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uhuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uhuo.mutation.RefCleared() {
		_spec.ClearField(userhistory.FieldRef, field.TypeString)
	}
	if uhuo.mutation.CreatedAtCleared() {
		_spec.ClearField(userhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := uhuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if uhuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(userhistory.FieldUpdatedAt, field.TypeTime)
	}
	if uhuo.mutation.CreatedByCleared() {
		_spec.ClearField(userhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := uhuo.mutation.UpdatedBy(); ok {
		_spec.SetField(userhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if uhuo.mutation.UpdatedByCleared() {
		_spec.ClearField(userhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := uhuo.mutation.DeletedAt(); ok {
		_spec.SetField(userhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if uhuo.mutation.DeletedAtCleared() {
		_spec.ClearField(userhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uhuo.mutation.DeletedBy(); ok {
		_spec.SetField(userhistory.FieldDeletedBy, field.TypeString, value)
	}
	if uhuo.mutation.DeletedByCleared() {
		_spec.ClearField(userhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := uhuo.mutation.Tags(); ok {
		_spec.SetField(userhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := uhuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, userhistory.FieldTags, value)
		})
	}
	if uhuo.mutation.TagsCleared() {
		_spec.ClearField(userhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := uhuo.mutation.Email(); ok {
		_spec.SetField(userhistory.FieldEmail, field.TypeString, value)
	}
	if value, ok := uhuo.mutation.FirstName(); ok {
		_spec.SetField(userhistory.FieldFirstName, field.TypeString, value)
	}
	if uhuo.mutation.FirstNameCleared() {
		_spec.ClearField(userhistory.FieldFirstName, field.TypeString)
	}
	if value, ok := uhuo.mutation.LastName(); ok {
		_spec.SetField(userhistory.FieldLastName, field.TypeString, value)
	}
	if uhuo.mutation.LastNameCleared() {
		_spec.ClearField(userhistory.FieldLastName, field.TypeString)
	}
	if value, ok := uhuo.mutation.DisplayName(); ok {
		_spec.SetField(userhistory.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uhuo.mutation.AvatarRemoteURL(); ok {
		_spec.SetField(userhistory.FieldAvatarRemoteURL, field.TypeString, value)
	}
	if uhuo.mutation.AvatarRemoteURLCleared() {
		_spec.ClearField(userhistory.FieldAvatarRemoteURL, field.TypeString)
	}
	if value, ok := uhuo.mutation.AvatarLocalFile(); ok {
		_spec.SetField(userhistory.FieldAvatarLocalFile, field.TypeString, value)
	}
	if uhuo.mutation.AvatarLocalFileCleared() {
		_spec.ClearField(userhistory.FieldAvatarLocalFile, field.TypeString)
	}
	if value, ok := uhuo.mutation.AvatarUpdatedAt(); ok {
		_spec.SetField(userhistory.FieldAvatarUpdatedAt, field.TypeTime, value)
	}
	if uhuo.mutation.AvatarUpdatedAtCleared() {
		_spec.ClearField(userhistory.FieldAvatarUpdatedAt, field.TypeTime)
	}
	if value, ok := uhuo.mutation.LastSeen(); ok {
		_spec.SetField(userhistory.FieldLastSeen, field.TypeTime, value)
	}
	if uhuo.mutation.LastSeenCleared() {
		_spec.ClearField(userhistory.FieldLastSeen, field.TypeTime)
	}
	if value, ok := uhuo.mutation.Password(); ok {
		_spec.SetField(userhistory.FieldPassword, field.TypeString, value)
	}
	if uhuo.mutation.PasswordCleared() {
		_spec.ClearField(userhistory.FieldPassword, field.TypeString)
	}
	if value, ok := uhuo.mutation.Sub(); ok {
		_spec.SetField(userhistory.FieldSub, field.TypeString, value)
	}
	if uhuo.mutation.SubCleared() {
		_spec.ClearField(userhistory.FieldSub, field.TypeString)
	}
	if value, ok := uhuo.mutation.AuthProvider(); ok {
		_spec.SetField(userhistory.FieldAuthProvider, field.TypeEnum, value)
	}
	if value, ok := uhuo.mutation.Role(); ok {
		_spec.SetField(userhistory.FieldRole, field.TypeEnum, value)
	}
	if uhuo.mutation.RoleCleared() {
		_spec.ClearField(userhistory.FieldRole, field.TypeEnum)
	}
	_spec.Node.Schema = uhuo.schemaConfig.UserHistory
	ctx = internal.NewSchemaConfigContext(ctx, uhuo.schemaConfig)
	_node = &UserHistory{config: uhuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uhuo.mutation.done = true
	return _node, nil
}
