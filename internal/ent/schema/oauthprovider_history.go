// Code generated by enthistory, DO NOT EDIT.
package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/enthistory"

	"github.com/datumforge/entx"

	"time"
)

// OauthProviderHistory holds the schema definition for the OauthProviderHistory entity.
type OauthProviderHistory struct {
	ent.Schema
}

// Annotations of the OauthProviderHistory.
func (OauthProviderHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.SchemaGenSkip(true),
		entsql.Annotation{
			Table: "oauth_provider_history",
		},
		enthistory.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

// Fields of the OauthProviderHistory.
func (OauthProviderHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.String("ref").
			Immutable().
			Optional(),
		field.Enum("operation").
			GoType(enthistory.OpType("")).
			Immutable(),
	}

	// get the fields from the mixins
	// we only want to include mixin fields, not edges
	// so this prevents FKs back to the main tables
	mixins := OauthProvider{}.Mixin()
	for _, mixin := range mixins {
		for _, field := range mixin.Fields() {
			historyFields = append(historyFields, field)
		}
	}

	original := OauthProvider{}
	for _, field := range original.Fields() {
		historyFields = append(historyFields, field)
	}

	return historyFields
}
func (OauthProviderHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("history_time"),
	}
}
