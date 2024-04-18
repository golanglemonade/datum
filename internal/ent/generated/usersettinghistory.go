// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/usersettinghistory"
	"github.com/datumforge/enthistory"
)

// UserSettingHistory is the model entity for the UserSettingHistory schema.
type UserSettingHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref string `json:"ref,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// user account is locked if unconfirmed or explicitly locked
	Locked bool `json:"locked,omitempty"`
	// The time notifications regarding the user were silenced
	SilencedAt *time.Time `json:"silenced_at,omitempty"`
	// The time the user was suspended
	SuspendedAt *time.Time `json:"suspended_at,omitempty"`
	// Status holds the value of the "status" field.
	Status enums.UserStatus `json:"status,omitempty"`
	// EmailConfirmed holds the value of the "email_confirmed" field.
	EmailConfirmed bool `json:"email_confirmed,omitempty"`
	// tags associated with the user
	Tags []string `json:"tags,omitempty"`
	// specifies a user may complete authentication by verifying a WebAuthn capable device
	IsWebauthnAllowed bool `json:"is_webauthn_allowed,omitempty"`
	// whether the user has two factor authentication enabled
	IsTfaEnabled bool `json:"is_tfa_enabled,omitempty"`
	// phone number associated with the account, used 2factor SMS authentication
	PhoneNumber  *string `json:"phone_number,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserSettingHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usersettinghistory.FieldTags:
			values[i] = new([]byte)
		case usersettinghistory.FieldOperation:
			values[i] = new(enthistory.OpType)
		case usersettinghistory.FieldLocked, usersettinghistory.FieldEmailConfirmed, usersettinghistory.FieldIsWebauthnAllowed, usersettinghistory.FieldIsTfaEnabled:
			values[i] = new(sql.NullBool)
		case usersettinghistory.FieldID, usersettinghistory.FieldRef, usersettinghistory.FieldCreatedBy, usersettinghistory.FieldUpdatedBy, usersettinghistory.FieldDeletedBy, usersettinghistory.FieldUserID, usersettinghistory.FieldStatus, usersettinghistory.FieldPhoneNumber:
			values[i] = new(sql.NullString)
		case usersettinghistory.FieldHistoryTime, usersettinghistory.FieldCreatedAt, usersettinghistory.FieldUpdatedAt, usersettinghistory.FieldDeletedAt, usersettinghistory.FieldSilencedAt, usersettinghistory.FieldSuspendedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserSettingHistory fields.
func (ush *UserSettingHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usersettinghistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ush.ID = value.String
			}
		case usersettinghistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				ush.HistoryTime = value.Time
			}
		case usersettinghistory.FieldOperation:
			if value, ok := values[i].(*enthistory.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				ush.Operation = *value
			}
		case usersettinghistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				ush.Ref = value.String
			}
		case usersettinghistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ush.CreatedAt = value.Time
			}
		case usersettinghistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ush.UpdatedAt = value.Time
			}
		case usersettinghistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ush.CreatedBy = value.String
			}
		case usersettinghistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ush.UpdatedBy = value.String
			}
		case usersettinghistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ush.DeletedAt = value.Time
			}
		case usersettinghistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				ush.DeletedBy = value.String
			}
		case usersettinghistory.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ush.UserID = value.String
			}
		case usersettinghistory.FieldLocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				ush.Locked = value.Bool
			}
		case usersettinghistory.FieldSilencedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field silenced_at", values[i])
			} else if value.Valid {
				ush.SilencedAt = new(time.Time)
				*ush.SilencedAt = value.Time
			}
		case usersettinghistory.FieldSuspendedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field suspended_at", values[i])
			} else if value.Valid {
				ush.SuspendedAt = new(time.Time)
				*ush.SuspendedAt = value.Time
			}
		case usersettinghistory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ush.Status = enums.UserStatus(value.String)
			}
		case usersettinghistory.FieldEmailConfirmed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_confirmed", values[i])
			} else if value.Valid {
				ush.EmailConfirmed = value.Bool
			}
		case usersettinghistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ush.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case usersettinghistory.FieldIsWebauthnAllowed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_webauthn_allowed", values[i])
			} else if value.Valid {
				ush.IsWebauthnAllowed = value.Bool
			}
		case usersettinghistory.FieldIsTfaEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_tfa_enabled", values[i])
			} else if value.Valid {
				ush.IsTfaEnabled = value.Bool
			}
		case usersettinghistory.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				ush.PhoneNumber = new(string)
				*ush.PhoneNumber = value.String
			}
		default:
			ush.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserSettingHistory.
// This includes values selected through modifiers, order, etc.
func (ush *UserSettingHistory) Value(name string) (ent.Value, error) {
	return ush.selectValues.Get(name)
}

// Update returns a builder for updating this UserSettingHistory.
// Note that you need to call UserSettingHistory.Unwrap() before calling this method if this UserSettingHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ush *UserSettingHistory) Update() *UserSettingHistoryUpdateOne {
	return NewUserSettingHistoryClient(ush.config).UpdateOne(ush)
}

// Unwrap unwraps the UserSettingHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ush *UserSettingHistory) Unwrap() *UserSettingHistory {
	_tx, ok := ush.config.driver.(*txDriver)
	if !ok {
		panic("generated: UserSettingHistory is not a transactional entity")
	}
	ush.config.driver = _tx.drv
	return ush
}

// String implements the fmt.Stringer.
func (ush *UserSettingHistory) String() string {
	var builder strings.Builder
	builder.WriteString("UserSettingHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ush.ID))
	builder.WriteString("history_time=")
	builder.WriteString(ush.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", ush.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(ush.Ref)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ush.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ush.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(ush.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(ush.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ush.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(ush.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(ush.UserID)
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", ush.Locked))
	builder.WriteString(", ")
	if v := ush.SilencedAt; v != nil {
		builder.WriteString("silenced_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := ush.SuspendedAt; v != nil {
		builder.WriteString("suspended_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ush.Status))
	builder.WriteString(", ")
	builder.WriteString("email_confirmed=")
	builder.WriteString(fmt.Sprintf("%v", ush.EmailConfirmed))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", ush.Tags))
	builder.WriteString(", ")
	builder.WriteString("is_webauthn_allowed=")
	builder.WriteString(fmt.Sprintf("%v", ush.IsWebauthnAllowed))
	builder.WriteString(", ")
	builder.WriteString("is_tfa_enabled=")
	builder.WriteString(fmt.Sprintf("%v", ush.IsTfaEnabled))
	builder.WriteString(", ")
	if v := ush.PhoneNumber; v != nil {
		builder.WriteString("phone_number=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// UserSettingHistories is a parsable slice of UserSettingHistory.
type UserSettingHistories []*UserSettingHistory