// Code generated by ent, DO NOT EDIT.

package groupsetting

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
	// Label holds the string label denoting the groupsetting type in the database.
	Label = "group_setting"
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
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// FieldJoinPolicy holds the string denoting the join_policy field in the database.
	FieldJoinPolicy = "join_policy"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldSyncToSlack holds the string denoting the sync_to_slack field in the database.
	FieldSyncToSlack = "sync_to_slack"
	// FieldSyncToGithub holds the string denoting the sync_to_github field in the database.
	FieldSyncToGithub = "sync_to_github"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the groupsetting in the database.
	Table = "group_settings"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "group_settings"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_setting"
)

// Columns holds all SQL columns for groupsetting fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldVisibility,
	FieldJoinPolicy,
	FieldTags,
	FieldSyncToSlack,
	FieldSyncToGithub,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "group_settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_setting",
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
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultSyncToSlack holds the default value on creation for the "sync_to_slack" field.
	DefaultSyncToSlack bool
	// DefaultSyncToGithub holds the default value on creation for the "sync_to_github" field.
	DefaultSyncToGithub bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// VisibilityPublic is the default value of the Visibility enum.
const DefaultVisibility = VisibilityPublic

// Visibility values.
const (
	VisibilityPublic  Visibility = "PUBLIC"
	VisibilityPrivate Visibility = "PRIVATE"
)

func (v Visibility) String() string {
	return string(v)
}

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v Visibility) error {
	switch v {
	case VisibilityPublic, VisibilityPrivate:
		return nil
	default:
		return fmt.Errorf("groupsetting: invalid enum value for visibility field: %q", v)
	}
}

// JoinPolicy defines the type for the "join_policy" enum field.
type JoinPolicy string

// JoinPolicyInviteOrApplication is the default value of the JoinPolicy enum.
const DefaultJoinPolicy = JoinPolicyInviteOrApplication

// JoinPolicy values.
const (
	JoinPolicyOpen                JoinPolicy = "OPEN"
	JoinPolicyInviteOnly          JoinPolicy = "INVITE_ONLY"
	JoinPolicyApplicationOnly     JoinPolicy = "APPLICATION_ONLY"
	JoinPolicyInviteOrApplication JoinPolicy = "INVITE_OR_APPLICATION"
)

func (jp JoinPolicy) String() string {
	return string(jp)
}

// JoinPolicyValidator is a validator for the "join_policy" field enum values. It is called by the builders before save.
func JoinPolicyValidator(jp JoinPolicy) error {
	switch jp {
	case JoinPolicyOpen, JoinPolicyInviteOnly, JoinPolicyApplicationOnly, JoinPolicyInviteOrApplication:
		return nil
	default:
		return fmt.Errorf("groupsetting: invalid enum value for join_policy field: %q", jp)
	}
}

// OrderOption defines the ordering options for the GroupSetting queries.
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

// ByVisibility orders the results by the visibility field.
func ByVisibility(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisibility, opts...).ToFunc()
}

// ByJoinPolicy orders the results by the join_policy field.
func ByJoinPolicy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJoinPolicy, opts...).ToFunc()
}

// BySyncToSlack orders the results by the sync_to_slack field.
func BySyncToSlack(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSyncToSlack, opts...).ToFunc()
}

// BySyncToGithub orders the results by the sync_to_github field.
func BySyncToGithub(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSyncToGithub, opts...).ToFunc()
}

// ByGroupField orders the results by group field.
func ByGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupStep(), sql.OrderByField(field, opts...))
	}
}
func newGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, GroupTable, GroupColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Visibility) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Visibility) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Visibility(str)
	if err := VisibilityValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Visibility", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e JoinPolicy) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *JoinPolicy) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = JoinPolicy(str)
	if err := JoinPolicyValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid JoinPolicy", str)
	}
	return nil
}
