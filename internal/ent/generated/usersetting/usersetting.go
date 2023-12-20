// Code generated by ent, DO NOT EDIT.

package usersetting

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the usersetting type in the database.
	Label = "user_setting"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldLocked holds the string denoting the locked field in the database.
	FieldLocked = "locked"
	// FieldSilencedAt holds the string denoting the silenced_at field in the database.
	FieldSilencedAt = "silenced_at"
	// FieldSuspendedAt holds the string denoting the suspended_at field in the database.
	FieldSuspendedAt = "suspended_at"
	// FieldRecoveryCode holds the string denoting the recovery_code field in the database.
	FieldRecoveryCode = "recovery_code"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldPermissions holds the string denoting the permissions field in the database.
	FieldPermissions = "permissions"
	// FieldEmailConfirmed holds the string denoting the email_confirmed field in the database.
	FieldEmailConfirmed = "email_confirmed"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the usersetting in the database.
	Table = "user_settings"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_settings"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_setting"
)

// Columns holds all SQL columns for usersetting fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldLocked,
	FieldSilencedAt,
	FieldSuspendedAt,
	FieldRecoveryCode,
	FieldStatus,
	FieldRole,
	FieldPermissions,
	FieldEmailConfirmed,
	FieldTags,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_setting",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/datumforge/datum/internal/ent/generated/runtime"
var (
	Hooks        [2]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultLocked holds the default value on creation for the "locked" field.
	DefaultLocked bool
	// DefaultPermissions holds the default value on creation for the "permissions" field.
	DefaultPermissions []string
	// DefaultEmailConfirmed holds the default value on creation for the "email_confirmed" field.
	DefaultEmailConfirmed bool
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// Status defines the type for the "status" enum field.
type Status string

// StatusActive is the default value of the Status enum.
const DefaultStatus = StatusActive

// Status values.
const (
	StatusActive      Status = "ACTIVE"
	StatusInactive    Status = "INACTIVE"
	StatusDeactivated Status = "DEACTIVATED"
	StatusSuspended   Status = "SUSPENDED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusInactive, StatusDeactivated, StatusSuspended:
		return nil
	default:
		return fmt.Errorf("usersetting: invalid enum value for status field: %q", s)
	}
}

// Role defines the type for the "role" enum field.
type Role string

// RoleUser is the default value of the Role enum.
const DefaultRole = RoleUser

// Role values.
const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
	RoleOwner Role = "OWNER"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleUser, RoleAdmin, RoleOwner:
		return nil
	default:
		return fmt.Errorf("usersetting: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the UserSetting queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByLocked orders the results by the locked field.
func ByLocked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocked, opts...).ToFunc()
}

// BySilencedAt orders the results by the silenced_at field.
func BySilencedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSilencedAt, opts...).ToFunc()
}

// BySuspendedAt orders the results by the suspended_at field.
func BySuspendedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSuspendedAt, opts...).ToFunc()
}

// ByRecoveryCode orders the results by the recovery_code field.
func ByRecoveryCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecoveryCode, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByEmailConfirmed orders the results by the email_confirmed field.
func ByEmailConfirmed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailConfirmed, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Role) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Role) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Role(str)
	if err := RoleValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}
