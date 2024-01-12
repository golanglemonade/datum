package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateOrgMembers is the resolver for the createOrgMembers field.
func (r *createOrganizationInputResolver) CreateOrgMembers(ctx context.Context, obj *generated.CreateOrganizationInput, data []*generated.CreateOrgMembershipInput) error {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.OrgMembershipCreate, len(data))
	for i := range data {
		builders[i] = c.OrgMembership.Create().SetInput(*data[i])
	}

	orgMembers, err := c.OrgMembership.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return err
	}

	ids := make([]string, len(orgMembers))
	for i, member := range orgMembers {
		ids[i] = member.UserID
	}

	obj.UserIDs = append(obj.UserIDs, ids...)

	return nil
}

// CreateOrgSettings is the resolver for the createOrgSettings field.
func (r *createOrganizationInputResolver) CreateOrgSettings(ctx context.Context, obj *generated.CreateOrganizationInput, data *generated.CreateOrganizationSettingInput) error {
	c := withTransactionalMutation(ctx)

	orgSettings, err := c.OrganizationSetting.Create().SetInput(*data).Save(ctx)
	if err != nil {
		return err
	}

	obj.SettingID = &orgSettings.ID

	return nil
}