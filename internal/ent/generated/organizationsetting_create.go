// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
)

// OrganizationSettingCreate is the builder for creating a OrganizationSetting entity.
type OrganizationSettingCreate struct {
	config
	mutation *OrganizationSettingMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (osc *OrganizationSettingCreate) SetCreatedAt(t time.Time) *OrganizationSettingCreate {
	osc.mutation.SetCreatedAt(t)
	return osc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableCreatedAt(t *time.Time) *OrganizationSettingCreate {
	if t != nil {
		osc.SetCreatedAt(*t)
	}
	return osc
}

// SetUpdatedAt sets the "updated_at" field.
func (osc *OrganizationSettingCreate) SetUpdatedAt(t time.Time) *OrganizationSettingCreate {
	osc.mutation.SetUpdatedAt(t)
	return osc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableUpdatedAt(t *time.Time) *OrganizationSettingCreate {
	if t != nil {
		osc.SetUpdatedAt(*t)
	}
	return osc
}

// SetCreatedBy sets the "created_by" field.
func (osc *OrganizationSettingCreate) SetCreatedBy(s string) *OrganizationSettingCreate {
	osc.mutation.SetCreatedBy(s)
	return osc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableCreatedBy(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetCreatedBy(*s)
	}
	return osc
}

// SetUpdatedBy sets the "updated_by" field.
func (osc *OrganizationSettingCreate) SetUpdatedBy(s string) *OrganizationSettingCreate {
	osc.mutation.SetUpdatedBy(s)
	return osc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableUpdatedBy(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetUpdatedBy(*s)
	}
	return osc
}

// SetDeletedAt sets the "deleted_at" field.
func (osc *OrganizationSettingCreate) SetDeletedAt(t time.Time) *OrganizationSettingCreate {
	osc.mutation.SetDeletedAt(t)
	return osc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableDeletedAt(t *time.Time) *OrganizationSettingCreate {
	if t != nil {
		osc.SetDeletedAt(*t)
	}
	return osc
}

// SetDeletedBy sets the "deleted_by" field.
func (osc *OrganizationSettingCreate) SetDeletedBy(s string) *OrganizationSettingCreate {
	osc.mutation.SetDeletedBy(s)
	return osc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableDeletedBy(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetDeletedBy(*s)
	}
	return osc
}

// SetDomains sets the "domains" field.
func (osc *OrganizationSettingCreate) SetDomains(s []string) *OrganizationSettingCreate {
	osc.mutation.SetDomains(s)
	return osc
}

// SetBillingContact sets the "billing_contact" field.
func (osc *OrganizationSettingCreate) SetBillingContact(s string) *OrganizationSettingCreate {
	osc.mutation.SetBillingContact(s)
	return osc
}

// SetNillableBillingContact sets the "billing_contact" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableBillingContact(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetBillingContact(*s)
	}
	return osc
}

// SetBillingEmail sets the "billing_email" field.
func (osc *OrganizationSettingCreate) SetBillingEmail(s string) *OrganizationSettingCreate {
	osc.mutation.SetBillingEmail(s)
	return osc
}

// SetNillableBillingEmail sets the "billing_email" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableBillingEmail(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetBillingEmail(*s)
	}
	return osc
}

// SetBillingPhone sets the "billing_phone" field.
func (osc *OrganizationSettingCreate) SetBillingPhone(s string) *OrganizationSettingCreate {
	osc.mutation.SetBillingPhone(s)
	return osc
}

// SetNillableBillingPhone sets the "billing_phone" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableBillingPhone(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetBillingPhone(*s)
	}
	return osc
}

// SetBillingAddress sets the "billing_address" field.
func (osc *OrganizationSettingCreate) SetBillingAddress(s string) *OrganizationSettingCreate {
	osc.mutation.SetBillingAddress(s)
	return osc
}

// SetNillableBillingAddress sets the "billing_address" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableBillingAddress(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetBillingAddress(*s)
	}
	return osc
}

// SetTaxIdentifier sets the "tax_identifier" field.
func (osc *OrganizationSettingCreate) SetTaxIdentifier(s string) *OrganizationSettingCreate {
	osc.mutation.SetTaxIdentifier(s)
	return osc
}

// SetNillableTaxIdentifier sets the "tax_identifier" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableTaxIdentifier(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetTaxIdentifier(*s)
	}
	return osc
}

// SetTags sets the "tags" field.
func (osc *OrganizationSettingCreate) SetTags(s []string) *OrganizationSettingCreate {
	osc.mutation.SetTags(s)
	return osc
}

// SetGeoLocation sets the "geo_location" field.
func (osc *OrganizationSettingCreate) SetGeoLocation(e enums.Region) *OrganizationSettingCreate {
	osc.mutation.SetGeoLocation(e)
	return osc
}

// SetNillableGeoLocation sets the "geo_location" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableGeoLocation(e *enums.Region) *OrganizationSettingCreate {
	if e != nil {
		osc.SetGeoLocation(*e)
	}
	return osc
}

// SetOrganizationID sets the "organization_id" field.
func (osc *OrganizationSettingCreate) SetOrganizationID(s string) *OrganizationSettingCreate {
	osc.mutation.SetOrganizationID(s)
	return osc
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableOrganizationID(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetOrganizationID(*s)
	}
	return osc
}

// SetID sets the "id" field.
func (osc *OrganizationSettingCreate) SetID(s string) *OrganizationSettingCreate {
	osc.mutation.SetID(s)
	return osc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (osc *OrganizationSettingCreate) SetNillableID(s *string) *OrganizationSettingCreate {
	if s != nil {
		osc.SetID(*s)
	}
	return osc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (osc *OrganizationSettingCreate) SetOrganization(o *Organization) *OrganizationSettingCreate {
	return osc.SetOrganizationID(o.ID)
}

// Mutation returns the OrganizationSettingMutation object of the builder.
func (osc *OrganizationSettingCreate) Mutation() *OrganizationSettingMutation {
	return osc.mutation
}

// Save creates the OrganizationSetting in the database.
func (osc *OrganizationSettingCreate) Save(ctx context.Context) (*OrganizationSetting, error) {
	if err := osc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, osc.sqlSave, osc.mutation, osc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (osc *OrganizationSettingCreate) SaveX(ctx context.Context) *OrganizationSetting {
	v, err := osc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (osc *OrganizationSettingCreate) Exec(ctx context.Context) error {
	_, err := osc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osc *OrganizationSettingCreate) ExecX(ctx context.Context) {
	if err := osc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osc *OrganizationSettingCreate) defaults() error {
	if _, ok := osc.mutation.CreatedAt(); !ok {
		if organizationsetting.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized organizationsetting.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := organizationsetting.DefaultCreatedAt()
		osc.mutation.SetCreatedAt(v)
	}
	if _, ok := osc.mutation.UpdatedAt(); !ok {
		if organizationsetting.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized organizationsetting.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := organizationsetting.DefaultUpdatedAt()
		osc.mutation.SetUpdatedAt(v)
	}
	if _, ok := osc.mutation.Tags(); !ok {
		v := organizationsetting.DefaultTags
		osc.mutation.SetTags(v)
	}
	if _, ok := osc.mutation.GeoLocation(); !ok {
		v := organizationsetting.DefaultGeoLocation
		osc.mutation.SetGeoLocation(v)
	}
	if _, ok := osc.mutation.ID(); !ok {
		if organizationsetting.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized organizationsetting.DefaultID (forgotten import generated/runtime?)")
		}
		v := organizationsetting.DefaultID()
		osc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (osc *OrganizationSettingCreate) check() error {
	if v, ok := osc.mutation.Domains(); ok {
		if err := organizationsetting.DomainsValidator(v); err != nil {
			return &ValidationError{Name: "domains", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.domains": %w`, err)}
		}
	}
	if v, ok := osc.mutation.BillingEmail(); ok {
		if err := organizationsetting.BillingEmailValidator(v); err != nil {
			return &ValidationError{Name: "billing_email", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_email": %w`, err)}
		}
	}
	if v, ok := osc.mutation.BillingPhone(); ok {
		if err := organizationsetting.BillingPhoneValidator(v); err != nil {
			return &ValidationError{Name: "billing_phone", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_phone": %w`, err)}
		}
	}
	if v, ok := osc.mutation.GeoLocation(); ok {
		if err := organizationsetting.GeoLocationValidator(v); err != nil {
			return &ValidationError{Name: "geo_location", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.geo_location": %w`, err)}
		}
	}
	return nil
}

func (osc *OrganizationSettingCreate) sqlSave(ctx context.Context) (*OrganizationSetting, error) {
	if err := osc.check(); err != nil {
		return nil, err
	}
	_node, _spec := osc.createSpec()
	if err := sqlgraph.CreateNode(ctx, osc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OrganizationSetting.ID type: %T", _spec.ID.Value)
		}
	}
	osc.mutation.id = &_node.ID
	osc.mutation.done = true
	return _node, nil
}

func (osc *OrganizationSettingCreate) createSpec() (*OrganizationSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationSetting{config: osc.config}
		_spec = sqlgraph.NewCreateSpec(organizationsetting.Table, sqlgraph.NewFieldSpec(organizationsetting.FieldID, field.TypeString))
	)
	_spec.Schema = osc.schemaConfig.OrganizationSetting
	if id, ok := osc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := osc.mutation.CreatedAt(); ok {
		_spec.SetField(organizationsetting.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := osc.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationsetting.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := osc.mutation.CreatedBy(); ok {
		_spec.SetField(organizationsetting.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := osc.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationsetting.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := osc.mutation.DeletedAt(); ok {
		_spec.SetField(organizationsetting.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := osc.mutation.DeletedBy(); ok {
		_spec.SetField(organizationsetting.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := osc.mutation.Domains(); ok {
		_spec.SetField(organizationsetting.FieldDomains, field.TypeJSON, value)
		_node.Domains = value
	}
	if value, ok := osc.mutation.BillingContact(); ok {
		_spec.SetField(organizationsetting.FieldBillingContact, field.TypeString, value)
		_node.BillingContact = value
	}
	if value, ok := osc.mutation.BillingEmail(); ok {
		_spec.SetField(organizationsetting.FieldBillingEmail, field.TypeString, value)
		_node.BillingEmail = value
	}
	if value, ok := osc.mutation.BillingPhone(); ok {
		_spec.SetField(organizationsetting.FieldBillingPhone, field.TypeString, value)
		_node.BillingPhone = value
	}
	if value, ok := osc.mutation.BillingAddress(); ok {
		_spec.SetField(organizationsetting.FieldBillingAddress, field.TypeString, value)
		_node.BillingAddress = value
	}
	if value, ok := osc.mutation.TaxIdentifier(); ok {
		_spec.SetField(organizationsetting.FieldTaxIdentifier, field.TypeString, value)
		_node.TaxIdentifier = value
	}
	if value, ok := osc.mutation.Tags(); ok {
		_spec.SetField(organizationsetting.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := osc.mutation.GeoLocation(); ok {
		_spec.SetField(organizationsetting.FieldGeoLocation, field.TypeEnum, value)
		_node.GeoLocation = value
	}
	if nodes := osc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   organizationsetting.OrganizationTable,
			Columns: []string{organizationsetting.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = osc.schemaConfig.OrganizationSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationSettingCreateBulk is the builder for creating many OrganizationSetting entities in bulk.
type OrganizationSettingCreateBulk struct {
	config
	err      error
	builders []*OrganizationSettingCreate
}

// Save creates the OrganizationSetting entities in the database.
func (oscb *OrganizationSettingCreateBulk) Save(ctx context.Context) ([]*OrganizationSetting, error) {
	if oscb.err != nil {
		return nil, oscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oscb.builders))
	nodes := make([]*OrganizationSetting, len(oscb.builders))
	mutators := make([]Mutator, len(oscb.builders))
	for i := range oscb.builders {
		func(i int, root context.Context) {
			builder := oscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationSettingMutation)
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
					_, err = mutators[i+1].Mutate(root, oscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oscb *OrganizationSettingCreateBulk) SaveX(ctx context.Context) []*OrganizationSetting {
	v, err := oscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oscb *OrganizationSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := oscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oscb *OrganizationSettingCreateBulk) ExecX(ctx context.Context) {
	if err := oscb.Exec(ctx); err != nil {
		panic(err)
	}
}
