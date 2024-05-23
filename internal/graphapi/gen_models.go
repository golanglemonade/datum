// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphapi

import (
	"entgo.io/contrib/entgql"
	"github.com/datumforge/datum/internal/ent/generated"
)

type GlobalSearchResult interface {
	IsGlobalSearchResult()
}

// Return response for createAPIToken mutation
type APITokenCreatePayload struct {
	// Created apiToken
	APIToken *generated.APIToken `json:"apiToken"`
}

// Return response for deleteAPIToken mutation
type APITokenDeletePayload struct {
	// Deleted apiToken ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateAPIToken mutation
type APITokenUpdatePayload struct {
	// Updated apiToken
	APIToken *generated.APIToken `json:"apiToken"`
}

// Return response for createDocumentData mutation
type DocumentDataCreatePayload struct {
	// Created documentData
	DocumentData *generated.DocumentData `json:"documentData"`
}

// Return response for deleteDocumentData mutation
type DocumentDataDeletePayload struct {
	// Deleted documentData ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateDocumentData mutation
type DocumentDataUpdatePayload struct {
	// Updated documentData
	DocumentData *generated.DocumentData `json:"documentData"`
}

type Dummy struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// Return response for createEntitlement mutation
type EntitlementCreatePayload struct {
	// Created entitlement
	Entitlement *generated.Entitlement `json:"entitlement"`
}

// Return response for deleteEntitlement mutation
type EntitlementDeletePayload struct {
	// Deleted entitlement ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateEntitlement mutation
type EntitlementUpdatePayload struct {
	// Updated entitlement
	Entitlement *generated.Entitlement `json:"entitlement"`
}

// Return response for createEvent mutation
type EventCreatePayload struct {
	// Created event
	Event *generated.Event `json:"event"`
}

// Return response for deleteEvent mutation
type EventDeletePayload struct {
	// Deleted event ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateEvent mutation
type EventUpdatePayload struct {
	// Updated event
	Event *generated.Event `json:"event"`
}

// Return response for createFeature mutation
type FeatureCreatePayload struct {
	// Created feature
	Feature *generated.Feature `json:"feature"`
}

// Return response for deleteFeature mutation
type FeatureDeletePayload struct {
	// Deleted feature ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateFeature mutation
type FeatureUpdatePayload struct {
	// Updated feature
	Feature *generated.Feature `json:"feature"`
}

// Return response for createFile mutation
type FileCreatePayload struct {
	// Created file
	File *generated.File `json:"file"`
}

// Return response for deleteFile mutation
type FileDeletePayload struct {
	// Deleted file ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateFile mutation
type FileUpdatePayload struct {
	// Updated file
	File *generated.File `json:"file"`
}

type GlobalSearchResultConnection struct {
	Page  *entgql.PageInfo[string] `json:"page"`
	Nodes []GlobalSearchResult     `json:"nodes"`
}

// Return response for createGroup mutation
type GroupCreatePayload struct {
	// Created group
	Group *generated.Group `json:"group"`
}

// Return response for deleteGroup mutation
type GroupDeletePayload struct {
	// Deleted group ID
	DeletedID string `json:"deletedID"`
}

// Return response for createGroupMembership mutation
type GroupMembershipCreatePayload struct {
	// Created groupMembership
	GroupMembership *generated.GroupMembership `json:"groupMembership"`
}

// Return response for deleteGroupMembership mutation
type GroupMembershipDeletePayload struct {
	// Deleted groupMembership ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateGroupMembership mutation
type GroupMembershipUpdatePayload struct {
	// Updated groupMembership
	GroupMembership *generated.GroupMembership `json:"groupMembership"`
}

type GroupSearchResult struct {
	Groups []*generated.Group `json:"groups,omitempty"`
}

func (GroupSearchResult) IsGlobalSearchResult() {}

// Return response for createGroupSetting mutation
type GroupSettingCreatePayload struct {
	// Created groupSetting
	GroupSetting *generated.GroupSetting `json:"groupSetting"`
}

// Return response for deleteGroupSetting mutation
type GroupSettingDeletePayload struct {
	// Deleted groupSetting ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateGroupSetting mutation
type GroupSettingUpdatePayload struct {
	// Updated groupSetting
	GroupSetting *generated.GroupSetting `json:"groupSetting"`
}

// Return response for updateGroup mutation
type GroupUpdatePayload struct {
	// Updated group
	Group *generated.Group `json:"group"`
}

// Return response for createHush mutation
type HushCreatePayload struct {
	// Created hush
	Hush *generated.Hush `json:"hush"`
}

// Return response for deleteHush mutation
type HushDeletePayload struct {
	// Deleted hush ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateHush mutation
type HushUpdatePayload struct {
	// Updated hush
	Hush *generated.Hush `json:"hush"`
}

// Return response for createIntegration mutation
type IntegrationCreatePayload struct {
	// Created integration
	Integration *generated.Integration `json:"integration"`
}

// Return response for deleteIntegration mutation
type IntegrationDeletePayload struct {
	// Deleted integration ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateIntegration mutation
type IntegrationUpdatePayload struct {
	// Updated integration
	Integration *generated.Integration `json:"integration"`
}

// Return response for createInvite mutation
type InviteCreatePayload struct {
	// Created invite
	Invite *generated.Invite `json:"invite"`
}

// Return response for deleteInvite mutation
type InviteDeletePayload struct {
	// Deleted invite ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateInvite mutation
type InviteUpdatePayload struct {
	// Updated invite
	Invite *generated.Invite `json:"invite"`
}

// Return response for createOauthProvider mutation
type OauthProviderCreatePayload struct {
	// Created oauthProvider
	OauthProvider *generated.OauthProvider `json:"oauthProvider"`
}

// Return response for deleteOauthProvider mutation
type OauthProviderDeletePayload struct {
	// Deleted oauthProvider ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateOauthProvider mutation
type OauthProviderUpdatePayload struct {
	// Updated oauthProvider
	OauthProvider *generated.OauthProvider `json:"oauthProvider"`
}

// Return response for createOhAuthTooToken mutation
type OhAuthTooTokenCreatePayload struct {
	// Created ohAuthTooToken
	OhAuthTooToken *generated.OhAuthTooToken `json:"ohAuthTooToken"`
}

// Return response for deleteOhAuthTooToken mutation
type OhAuthTooTokenDeletePayload struct {
	// Deleted ohAuthTooToken ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateOhAuthTooToken mutation
type OhAuthTooTokenUpdatePayload struct {
	// Updated ohAuthTooToken
	OhAuthTooToken *generated.OhAuthTooToken `json:"ohAuthTooToken"`
}

// Return response for createOrgMembership mutation
type OrgMembershipCreatePayload struct {
	// Created orgMembership
	OrgMembership *generated.OrgMembership `json:"orgMembership"`
}

// Return response for deleteOrgMembership mutation
type OrgMembershipDeletePayload struct {
	// Deleted orgMembership ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateOrgMembership mutation
type OrgMembershipUpdatePayload struct {
	// Updated orgMembership
	OrgMembership *generated.OrgMembership `json:"orgMembership"`
}

// Return response for createOrganization mutation
type OrganizationCreatePayload struct {
	// Created organization
	Organization *generated.Organization `json:"organization"`
}

// Return response for deleteOrganization mutation
type OrganizationDeletePayload struct {
	// Deleted organization ID
	DeletedID string `json:"deletedID"`
}

type OrganizationSearchResult struct {
	Organizations []*generated.Organization `json:"organizations,omitempty"`
}

func (OrganizationSearchResult) IsGlobalSearchResult() {}

// Return response for createOrganizationSetting mutation
type OrganizationSettingCreatePayload struct {
	// Created organizationSetting
	OrganizationSetting *generated.OrganizationSetting `json:"organizationSetting"`
}

// Return response for deleteOrganizationSetting mutation
type OrganizationSettingDeletePayload struct {
	// Deleted organizationSetting ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateOrganizationSetting mutation
type OrganizationSettingUpdatePayload struct {
	// Updated organizationSetting
	OrganizationSetting *generated.OrganizationSetting `json:"organizationSetting"`
}

// Return response for updateOrganization mutation
type OrganizationUpdatePayload struct {
	// Updated organization
	Organization *generated.Organization `json:"organization"`
}

// Return response for createPersonalAccessToken mutation
type PersonalAccessTokenCreatePayload struct {
	// Created personalAccessToken
	PersonalAccessToken *generated.PersonalAccessToken `json:"personalAccessToken"`
}

// Return response for deletePersonalAccessToken mutation
type PersonalAccessTokenDeletePayload struct {
	// Deleted personalAccessToken ID
	DeletedID string `json:"deletedID"`
}

// Return response for updatePersonalAccessToken mutation
type PersonalAccessTokenUpdatePayload struct {
	// Updated personalAccessToken
	PersonalAccessToken *generated.PersonalAccessToken `json:"personalAccessToken"`
}

// Return response for createSubscriber mutation
type SubscriberCreatePayload struct {
	// Created subscriber
	Subscriber *generated.Subscriber `json:"subscriber"`
}

// Return response for deleteSubscriber mutation
type SubscriberDeletePayload struct {
	// Deleted subscriber email
	Email string `json:"email"`
}

type SubscriberSearchResult struct {
	Subscribers []*generated.Subscriber `json:"subscribers,omitempty"`
}

func (SubscriberSearchResult) IsGlobalSearchResult() {}

// Return response for updateSubscriber mutation
type SubscriberUpdatePayload struct {
	// Updated subscriber
	Subscriber *generated.Subscriber `json:"subscriber"`
}

type Subscription struct {
}

// Return response for createTFASetting mutation
type TFASettingCreatePayload struct {
	// Created tfaSetting
	TfaSetting *generated.TFASetting `json:"tfaSetting"`
}

// Return response for updateTFASetting mutation
type TFASettingUpdatePayload struct {
	// Updated tfaSetting
	TfaSetting *generated.TFASetting `json:"tfaSetting"`
}

// Return response for createTemplate mutation
type TemplateCreatePayload struct {
	// Created template
	Template *generated.Template `json:"template"`
}

// Return response for deleteTemplate mutation
type TemplateDeletePayload struct {
	// Deleted template ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateTemplate mutation
type TemplateUpdatePayload struct {
	// Updated template
	Template *generated.Template `json:"template"`
}

// Return response for createUser mutation
type UserCreatePayload struct {
	// Created user
	User *generated.User `json:"user"`
}

// Return response for deleteUser mutation
type UserDeletePayload struct {
	// Deleted user ID
	DeletedID string `json:"deletedID"`
}

type UserSearchResult struct {
	Users []*generated.User `json:"users,omitempty"`
}

func (UserSearchResult) IsGlobalSearchResult() {}

// Return response for updateUserSetting mutation
type UserSettingUpdatePayload struct {
	// Updated userSetting
	UserSetting *generated.UserSetting `json:"userSetting"`
}

// Return response for updateUser mutation
type UserUpdatePayload struct {
	// Updated user
	User *generated.User `json:"user"`
}

// Return response for createWebhook mutation
type WebhookCreatePayload struct {
	// Created webhook
	Webhook *generated.Webhook `json:"webhook"`
}

// Return response for deleteWebhook mutation
type WebhookDeletePayload struct {
	// Deleted webhook ID
	DeletedID string `json:"deletedID"`
}

// Return response for updateWebhook mutation
type WebhookUpdatePayload struct {
	// Updated webhook
	Webhook *generated.Webhook `json:"webhook"`
}
