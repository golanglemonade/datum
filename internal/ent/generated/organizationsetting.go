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
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
)

// OrganizationSetting is the model entity for the OrganizationSetting schema.
type OrganizationSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// domains associated with the organization
	Domains []string `json:"domains,omitempty"`
	// Name of the person to contact for billing
	BillingContact string `json:"billing_contact,omitempty"`
	// Email address of the person to contact for billing
	BillingEmail string `json:"billing_email,omitempty"`
	// Phone number to contact for billing
	BillingPhone string `json:"billing_phone,omitempty"`
	// Address to send billing information to
	BillingAddress string `json:"billing_address,omitempty"`
	// Usually government-issued tax ID or business ID such as ABN in Australia
	TaxIdentifier string `json:"tax_identifier,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// geographical location of the organization
	GeoLocation enums.Region `json:"geo_location,omitempty"`
	// the ID of the organization the settings belong to
	OrganizationID string `json:"organization_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationSettingQuery when eager-loading is set.
	Edges        OrganizationSettingEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrganizationSettingEdges holds the relations/edges for other nodes in the graph.
type OrganizationSettingEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationSettingEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrganizationSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organizationsetting.FieldDomains, organizationsetting.FieldTags:
			values[i] = new([]byte)
		case organizationsetting.FieldID, organizationsetting.FieldCreatedBy, organizationsetting.FieldUpdatedBy, organizationsetting.FieldMappingID, organizationsetting.FieldDeletedBy, organizationsetting.FieldBillingContact, organizationsetting.FieldBillingEmail, organizationsetting.FieldBillingPhone, organizationsetting.FieldBillingAddress, organizationsetting.FieldTaxIdentifier, organizationsetting.FieldGeoLocation, organizationsetting.FieldOrganizationID:
			values[i] = new(sql.NullString)
		case organizationsetting.FieldCreatedAt, organizationsetting.FieldUpdatedAt, organizationsetting.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrganizationSetting fields.
func (os *OrganizationSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organizationsetting.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				os.ID = value.String
			}
		case organizationsetting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				os.CreatedAt = value.Time
			}
		case organizationsetting.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				os.UpdatedAt = value.Time
			}
		case organizationsetting.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				os.CreatedBy = value.String
			}
		case organizationsetting.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				os.UpdatedBy = value.String
			}
		case organizationsetting.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				os.MappingID = value.String
			}
		case organizationsetting.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				os.DeletedAt = value.Time
			}
		case organizationsetting.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				os.DeletedBy = value.String
			}
		case organizationsetting.FieldDomains:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field domains", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &os.Domains); err != nil {
					return fmt.Errorf("unmarshal field domains: %w", err)
				}
			}
		case organizationsetting.FieldBillingContact:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field billing_contact", values[i])
			} else if value.Valid {
				os.BillingContact = value.String
			}
		case organizationsetting.FieldBillingEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field billing_email", values[i])
			} else if value.Valid {
				os.BillingEmail = value.String
			}
		case organizationsetting.FieldBillingPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field billing_phone", values[i])
			} else if value.Valid {
				os.BillingPhone = value.String
			}
		case organizationsetting.FieldBillingAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field billing_address", values[i])
			} else if value.Valid {
				os.BillingAddress = value.String
			}
		case organizationsetting.FieldTaxIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tax_identifier", values[i])
			} else if value.Valid {
				os.TaxIdentifier = value.String
			}
		case organizationsetting.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &os.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case organizationsetting.FieldGeoLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field geo_location", values[i])
			} else if value.Valid {
				os.GeoLocation = enums.Region(value.String)
			}
		case organizationsetting.FieldOrganizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				os.OrganizationID = value.String
			}
		default:
			os.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrganizationSetting.
// This includes values selected through modifiers, order, etc.
func (os *OrganizationSetting) Value(name string) (ent.Value, error) {
	return os.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the OrganizationSetting entity.
func (os *OrganizationSetting) QueryOrganization() *OrganizationQuery {
	return NewOrganizationSettingClient(os.config).QueryOrganization(os)
}

// Update returns a builder for updating this OrganizationSetting.
// Note that you need to call OrganizationSetting.Unwrap() before calling this method if this OrganizationSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *OrganizationSetting) Update() *OrganizationSettingUpdateOne {
	return NewOrganizationSettingClient(os.config).UpdateOne(os)
}

// Unwrap unwraps the OrganizationSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (os *OrganizationSetting) Unwrap() *OrganizationSetting {
	_tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("generated: OrganizationSetting is not a transactional entity")
	}
	os.config.driver = _tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *OrganizationSetting) String() string {
	var builder strings.Builder
	builder.WriteString("OrganizationSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", os.ID))
	builder.WriteString("created_at=")
	builder.WriteString(os.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(os.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(os.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(os.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(os.MappingID)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(os.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(os.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("domains=")
	builder.WriteString(fmt.Sprintf("%v", os.Domains))
	builder.WriteString(", ")
	builder.WriteString("billing_contact=")
	builder.WriteString(os.BillingContact)
	builder.WriteString(", ")
	builder.WriteString("billing_email=")
	builder.WriteString(os.BillingEmail)
	builder.WriteString(", ")
	builder.WriteString("billing_phone=")
	builder.WriteString(os.BillingPhone)
	builder.WriteString(", ")
	builder.WriteString("billing_address=")
	builder.WriteString(os.BillingAddress)
	builder.WriteString(", ")
	builder.WriteString("tax_identifier=")
	builder.WriteString(os.TaxIdentifier)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", os.Tags))
	builder.WriteString(", ")
	builder.WriteString("geo_location=")
	builder.WriteString(fmt.Sprintf("%v", os.GeoLocation))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(os.OrganizationID)
	builder.WriteByte(')')
	return builder.String()
}

// OrganizationSettings is a parsable slice of OrganizationSetting.
type OrganizationSettings []*OrganizationSetting
